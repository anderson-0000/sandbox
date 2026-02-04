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