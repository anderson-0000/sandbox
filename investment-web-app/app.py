from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS
import os
import yaml
from simulation_logic import run_monte_carlo_simulation

app = Flask(__name__, static_folder='static', static_url_path='')
CORS(app)

def load_defaults():
    default_values = {
        "inv1": {"initial": 10, "monthly": 1, "return": 5.0, "risk": 10.0, "change_settings": []},
        "inv2": {"initial": 10, "monthly": 1, "return": 5.0, "risk": 10.0, "change_settings": []},
        "existing_savings": 0, "current_age": 33, "total_period": 50,
        "crash_enabled": False,
        "withdrawal_settings": [],
        "life_events": []
    }
    user_defaults_path = os.path.expanduser('~/parameter/investment-web-app-defaults.yaml')
    local_defaults_path = os.path.join(os.path.dirname(__file__), 'defaults.yaml')
    defaults_file_path = user_defaults_path if os.path.exists(user_defaults_path) else (local_defaults_path if os.path.exists(local_defaults_path) else None)

    if defaults_file_path:
        try:
            with open(defaults_file_path, 'r', encoding='utf-8') as f:
                loaded = yaml.safe_load(f)
                for inv in ["inv1", "inv2"]:
                    if inv in loaded:
                        if "change_settings" in loaded[inv]:
                            default_values[inv]["change_settings"] = loaded[inv]["change_settings"]
                        for k in ["initial", "monthly", "return", "risk"]:
                            if k in loaded[inv]: default_values[inv][k] = loaded[inv][k]
                for k in ["existing_savings", "current_age", "total_period", "crash_enabled"]:
                    if k in loaded: default_values[k] = loaded[k]
                if "withdrawal_settings" in loaded:
                    default_values["withdrawal_settings"] = loaded["withdrawal_settings"]
                elif "withdrawal_monthly" in loaded and "withdrawal_start" in loaded:
                    # 後方互換性のため
                    default_values["withdrawal_settings"] = [{"year": loaded["withdrawal_start"], "value": loaded["withdrawal_monthly"], "type": "amount"}]
                
                if "life_events" in loaded: default_values["life_events"] = loaded["life_events"]
        except Exception: pass
    return default_values

@app.route('/')
def serve_index(): return send_from_directory(app.static_folder, 'index.html')
@app.route('/defaults')
def get_defaults(): return jsonify(load_defaults())

@app.route('/simulate', methods=['POST'])
def simulate():
    data = request.get_json()
    inv1, inv2 = data.get('investment1'), data.get('investment2')
    
    # 単位変換
    inv1['initial'] *= 10000; inv1['monthly'] *= 10000
    for s in inv1.get('change_settings', []): s['monthly'] *= 10000
    inv2['initial'] *= 10000; inv2['monthly'] *= 10000
    for s in inv2.get('change_settings', []): s['monthly'] *= 10000
    
    savings = data.get('existing_savings', 0) * 10000
    life_events = [{**e, "amount": e["amount"] * 10000} for e in data.get('life_events', [])]
    crash_enabled = data.get('market_event_enabled', False)
    total_period = data.get('total_period', 50)
    withdrawal_settings = data.get('withdrawal_settings', [])

    try:
        res = run_monte_carlo_simulation(inv1, inv2, savings, life_events, crash_enabled, withdrawal_settings, total_period)
        return jsonify(res)
    except Exception as e: return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, port=5001)
