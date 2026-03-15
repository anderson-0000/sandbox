from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS
import random
import math
import statistics
import os
import yaml # yamlモジュールをインポート
from simulation_logic import run_monte_carlo_simulation # 追加


app = Flask(__name__, static_folder='static', static_url_path='')
CORS(app) # 開発用にCORSを許可

# --- デフォルト値の読み込み ---
def load_defaults():
    default_values = {
        "inv1": {
            "initial": 10, # 10万円
            "monthly": 1,  # 1万円
            "return": 5.0,
            "risk": 10.0,
            "period": 20,
            "change_settings": []
        },
        "inv2": {
            "initial": 10, # 10万円
            "monthly": 1,  # 1万円
            "return": 5.0,
            "risk": 10.0,
            "period": 20,
            "change_settings": []
        },
        "existing_savings": 0,
        "current_age": 30,
        "crash_year": 5,
        "crash_rate": 30,
        "life_events": []
    }

    user_defaults_path = os.path.expanduser('~/parameter/investment-web-app-defaults.yaml')
    local_defaults_path = os.path.join(os.path.dirname(__file__), 'defaults.yaml')

    if os.path.exists(user_defaults_path):
        defaults_file_path = user_defaults_path
    elif os.path.exists(local_defaults_path):
        defaults_file_path = local_defaults_path
    else:
        defaults_file_path = None

    if defaults_file_path:
        try:
            with open(defaults_file_path, 'r', encoding='utf-8') as f:
                loaded_defaults = yaml.safe_load(f)
                for inv_key in ["inv1", "inv2"]:
                    if inv_key in loaded_defaults:
                        # 変更: change_settings を直接読み込む
                        if "change_settings" in loaded_defaults[inv_key] and isinstance(loaded_defaults[inv_key]["change_settings"], list):
                            # change_settings の monthly も万円から円に変換
                            default_values[inv_key]["change_settings"] = [
                                {**setting, "monthly": setting["monthly"] * 10000}
                                for setting in loaded_defaults[inv_key]["change_settings"]
                            ]
                        # その他のパラメーターは従来通り、金額は万円から円に変換
                        for param_key in ["initial", "monthly"]:
                            if param_key in loaded_defaults[inv_key]:
                                default_values[inv_key][param_key] = loaded_defaults[inv_key][param_key]
                        for param_key in ["return", "risk", "period"]: # 金額ではないパラメータ
                            if param_key in loaded_defaults[inv_key]:
                                default_values[inv_key][param_key] = loaded_defaults[inv_key][param_key]
                
                if "existing_savings" in loaded_defaults:
                    default_values["existing_savings"] = loaded_defaults["existing_savings"]
                if "current_age" in loaded_defaults:
                    default_values["current_age"] = loaded_defaults["current_age"]
                if "crash_year" in loaded_defaults:
                    default_values["crash_year"] = loaded_defaults["crash_year"]
                if "crash_rate" in loaded_defaults:
                    default_values["crash_rate"] = loaded_defaults["crash_rate"]
                if "life_events" in loaded_defaults and isinstance(loaded_defaults["life_events"], list):
                    # life_events の amount も万円から円に変換
                    default_values["life_events"] = [
                        {**event, "amount": event["amount"]}
                        for event in loaded_defaults["life_events"]
                    ]
        except yaml.YAMLError:
            app.logger.warning(f"defaults.yaml is malformed, using hardcoded defaults. Path: {defaults_file_path}")
        except Exception as e:
            app.logger.error(f"Error loading defaults.yaml: {e}, using hardcoded defaults.")
    else:
        app.logger.info(f"defaults.yaml not found at {defaults_file_path}, using hardcoded defaults.")
            
    return default_values

# --- Flask ルート ---

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
    life_events = data.get('life_events', [])
    market_event = data.get('market_event')

    if not all([investment1, investment2]):
        return jsonify({"error": "Missing investment data"}), 400

    # 金額関連の値を万円から円に変換
    investment1['initial'] *= 10000
    investment1['monthly'] *= 10000
    for setting in investment1.get('change_settings', []):
        setting['monthly'] *= 10000

    investment2['initial'] *= 10000
    investment2['monthly'] *= 10000
    for setting in investment2.get('change_settings', []):
        setting['monthly'] *= 10000
    
    existing_savings *= 10000
    for event in life_events:
        event['amount'] *= 10000

    try:
        simulation_output = run_monte_carlo_simulation(investment1, investment2, existing_savings, life_events, market_event)
        return jsonify(simulation_output) # 辞書全体を返す
    except Exception as e:
        app.logger.error(f"Simulation error: {e}")
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, port=5001)
