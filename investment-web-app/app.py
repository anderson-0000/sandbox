from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS
import random
import math
import statistics
import os
import yaml # yamlモジュールをインポート


app = Flask(__name__, static_folder='static', static_url_path='')
CORS(app) # 開発用にCORSを許可

# --- デフォルト値の読み込み ---
def load_defaults():
    default_values = {
        "inv1": {
            "initial": 100000,
            "monthly": 10000,
            "return": 5.0,
            "risk": 10.0,
            "period": 20,
            "change_year": 0,
            "changed_monthly": 0
        },
        "inv2": {
            "initial": 100000,
            "monthly": 10000,
            "return": 5.0,
            "risk": 10.0,
            "period": 20,
            "change_year": 0,
            "changed_monthly": 0
        },
        "existing_savings": 0
    }

    defaults_file_path = os.path.join(os.path.dirname(__file__), 'defaults.yaml')
    if os.path.exists(defaults_file_path):
        try:
            with open(defaults_file_path, 'r', encoding='utf-8') as f:
                loaded_defaults = yaml.safe_load(f)
                for inv_key in ["inv1", "inv2"]:
                    if inv_key in loaded_defaults:
                        for param_key in default_values[inv_key]:
                            if param_key in loaded_defaults[inv_key]:
                                default_values[inv_key][param_key] = loaded_defaults[inv_key][param_key]
                if "existing_savings" in loaded_defaults:
                    default_values["existing_savings"] = loaded_defaults["existing_savings"]
        except yaml.YAMLError:
            app.logger.warning(f"defaults.yaml is malformed, using hardcoded defaults. Path: {defaults_file_path}")
        except Exception as e:
            app.logger.error(f"Error loading defaults.yaml: {e}, using hardcoded defaults.")
    else:
        app.logger.info(f"defaults.yaml not found at {defaults_file_path}, using hardcoded defaults.")
            
    return default_values

# --- モンテカルロシミュレーションロジック ---

def get_gaussian_random():
    """Box-Muller変換を使用して、標準正規分布に従う乱数を生成"""
    u1, u2 = random.random(), random.random()
    return math.sqrt(-2.0 * math.log(u1)) * math.cos(2.0 * math.pi * u2)

def get_normal_random(mean, std_dev):
    """正規分布に従う乱数を生成"""
    return mean + get_gaussian_random() * std_dev

def run_monte_carlo_simulation(investment_data1, investment_data2, existing_savings, life_events, num_simulations=5000):
    all_simulation_paths = []
    total_investment_period = max(investment_data1['period'], investment_data2['period'])

    # ライフイベントを年ごとに辞書にまとめる
    life_events_by_year = {}
    for event in life_events:
        year = event['year']
        amount = event['amount']
        if year not in life_events_by_year:
            life_events_by_year[year] = 0
        life_events_by_year[year] += amount

    for _ in range(num_simulations):
        path1_values = [investment_data1['initial']]
        path2_values = [investment_data2['initial']]

        val1 = investment_data1['initial']
        monthly1 = investment_data1['monthly']
        for year_idx in range(investment_data1['period']):
            if investment_data1['change_year'] > 0 and year_idx + 1 >= investment_data1['change_year']:
                monthly1 = investment_data1['changed_monthly']

            annual_return = get_normal_random(investment_data1['return'] / 100, investment_data1['risk'] / 100)
            val1 = val1 * (1 + annual_return) + (monthly1 * 12 * (1 + annual_return / 2))
            path1_values.append(val1)

        val2 = investment_data2['initial']
        monthly2 = investment_data2['monthly']
        for year_idx in range(investment_data2['period']):
            if investment_data2['change_year'] > 0 and year_idx + 1 >= investment_data2['change_year']:
                monthly2 = investment_data2['changed_monthly']
            
            annual_return = get_normal_random(investment_data2['return'] / 100, investment_data2['risk'] / 100)
            val2 = val2 * (1 + annual_return) + (monthly2 * 12 * (1 + annual_return / 2))
            path2_values.append(val2)
        
        # 2つのポートフォリオを合算し、既存貯金を加算、ライフイベント費用を減算
        combined_path = []
        for year_idx in range(total_investment_period):
            year_val1 = path1_values[year_idx+1] if year_idx < investment_data1['period'] else path1_values[investment_data1['period']]
            year_val2 = path2_values[year_idx+1] if year_idx < investment_data2['period'] else path2_values[investment_data2['period']]
            
            current_year_total = year_val1 + year_val2 + existing_savings

            # ライフイベント費用を差し引く
            if (year_idx + 1) in life_events_by_year:
                current_year_total -= life_events_by_year[year_idx + 1]

            combined_path.append(current_year_total)
        all_simulation_paths.append(combined_path)

    yearly_results = []
    for year_idx in range(total_investment_period):
        year_values = sorted([path[year_idx] for path in all_simulation_paths])
        
        # quantilesを計算するヘルパー関数
        def get_percentile(data, percentile):
            if not data:
                return 0
            k = (len(data) - 1) * percentile / 100.0
            f = math.floor(k)
            c = math.ceil(k)
            if f == c:
                return data[int(k)]
            d0 = data[int(f)] * (c - k)
            d1 = data[int(c)] * (k - f)
            return d0 + d1

        yearly_results.append({
            "year": year_idx + 1,
            "min": year_values[0],
            "p10": get_percentile(year_values, 10),
            "p25": get_percentile(year_values, 25),
            "median": statistics.median(year_values) if year_values else 0,
            "p75": get_percentile(year_values, 75),
            "p90": get_percentile(year_values, 90),
            "max": year_values[-1],
            "average": statistics.mean(year_values) if year_values else 0
        })
    return yearly_results

# --- Flask ルート ---

@app.route('/')
def serve_index():
    return send_from_directory(app.static_folder, 'index.html')

@app.route('/<path:filename>')
def serve_static(filename):
    return send_from_directory(app.static_folder, filename)

@app.route('/defaults')
def get_defaults():
    defaults = load_defaults()
    return jsonify(defaults)

@app.route('/simulate', methods=['POST'])
def simulate():
    data = request.get_json()
    if not data:
        return jsonify({"error": "No data provided"}), 400

    investment1 = data.get('investment1')
    investment2 = data.get('investment2')
    existing_savings = data.get('existing_savings', 0)
    life_events = data.get('life_events', []) # ライフイベントを取得、デフォルトは空リスト

    if not all([investment1, investment2]):
        return jsonify({"error": "Missing investment data"}), 400

    try:
        results = run_monte_carlo_simulation(investment1, investment2, existing_savings, life_events)
        return jsonify(results)
    except Exception as e:
        app.logger.error(f"Simulation error: {e}")
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, port=5001)