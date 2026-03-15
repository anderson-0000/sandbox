import random
import math
import statistics
from datetime import datetime

def get_gaussian_random():
    """Box-Muller変換を使用して、標準正規分布に従う乱数を生成"""
    u1, u2 = random.random(), random.random()
    return math.sqrt(-2.0 * math.log(u1)) * math.cos(2.0 * math.pi * u2)

def get_normal_random(mean, std_dev):
    """正規分布に従う乱数を生成"""
    return mean + get_gaussian_random() * std_dev

def run_monte_carlo_simulation(investment_data1, investment_data2, existing_savings, life_events, 
                               market_event=None, num_simulations=5000, num_sample_paths=100):
    """
    market_event: {
        "enabled": true,
        "start_year": 5,
        "interval_years": 10,
        "rate": 30
    }
    """
    all_simulation_paths = []
    total_investment_period_years = max(investment_data1['period'], investment_data2['period'])
    total_months = total_investment_period_years * 12

    # 市場イベント（周期的な暴落）の月リストを作成
    crash_months = []
    crash_factor = 1.0
    if market_event and market_event.get('enabled'):
        start_year = int(market_event['start_year'])
        interval_years = int(market_event['interval_years'])
        crash_rate = float(market_event['rate'])
        crash_factor = 1.0 - (crash_rate / 100.0)

        current_crash_year = start_year
        while current_crash_year <= total_investment_period_years:
            m = (current_crash_year - 1) * 12 + 1
            crash_months.append(m)
            current_crash_year += interval_years

    # ライフイベントを月単位に変換
    life_events_by_month = {}
    for event in life_events:
        month = int(event['year'] * 12)
        amount = event['amount']
        if month not in life_events_by_month:
            life_events_by_month[month] = 0
        life_events_by_month[month] += amount

    def get_monthly_changes_by_month(change_settings):
        changes = {}
        sorted_settings = sorted(change_settings, key=lambda x: x['year'])
        for setting in sorted_settings:
            month = int(setting['year'] * 12)
            changes[month] = setting['monthly']
        return changes

    inv1_monthly_changes = get_monthly_changes_by_month(investment_data1.get('change_settings', []))
    inv2_monthly_changes = get_monthly_changes_by_month(investment_data2.get('change_settings', []))

    # 月次リターン・リスクへの変換
    r1_monthly_mean = (investment_data1['return'] / 100) / 12
    s1_monthly_std = (investment_data1['risk'] / 100) / math.sqrt(12)
    r2_monthly_mean = (investment_data2['return'] / 100) / 12
    s2_monthly_std = (investment_data2['risk'] / 100) / math.sqrt(12)

    inv1_is_lower_return = investment_data1['return'] <= investment_data2['return']

    for _ in range(num_simulations):
        val1 = investment_data1['initial']
        val2 = investment_data2['initial']
        
        combined_path = []
        initial_total = val1 + val2 + existing_savings
        combined_path.append(initial_total)

        current_monthly1 = investment_data1['monthly']
        current_monthly2 = investment_data2['monthly']

        for m in range(1, total_months + 1):
            if m in inv1_monthly_changes:
                current_monthly1 = inv1_monthly_changes[m]
            if m in inv2_monthly_changes:
                current_monthly2 = inv2_monthly_changes[m]

            # 1. リターンの適用 (正規分布)
            m_ret1 = get_normal_random(r1_monthly_mean, s1_monthly_std)
            m_ret2 = get_normal_random(r2_monthly_mean, s2_monthly_std)
            
            if m in crash_months:
                val1 = val1 * crash_factor
                val2 = val2 * crash_factor
            else:
                val1 = val1 * (1 + m_ret1)
                val2 = val2 * (1 + m_ret2)

            # 2. 積立の実施
            if m <= investment_data1['period'] * 12:
                val1 += current_monthly1
            if m <= investment_data2['period'] * 12:
                val2 += current_monthly2

            # 3. ライフイベント費用の差し引き
            if m in life_events_by_month:
                expense = life_events_by_month[m]
                if inv1_is_lower_return:
                    if val1 >= expense: val1 -= expense
                    else:
                        remaining = expense - val1
                        val1 = 0
                        val2 -= remaining
                else:
                    if val2 >= expense: val2 -= expense
                    else:
                        remaining = expense - val2
                        val2 = 0
                        val1 -= remaining
            
            combined_path.append(val1 + val2 + existing_savings)
        
        all_simulation_paths.append(combined_path)

    def get_percentile(data, percentile):
        if not data: return 0
        k = (len(data) - 1) * percentile / 100.0
        f = math.floor(k); c = math.ceil(k)
        if f == c: return data[int(k)]
        return data[int(f)] * (c - k) + data[int(c)] * (k - f)

    monthly_results = []
    for m in range(total_months + 1):
        m_values = sorted([path[m] for path in all_simulation_paths])
        monthly_results.append({
            "month": m,
            "min": m_values[0],
            "p10": get_percentile(m_values, 10),
            "p25": get_percentile(m_values, 25),
            "median": statistics.median(m_values),
            "p75": get_percentile(m_values, 75),
            "p90": get_percentile(m_values, 90),
            "max": m_values[-1],
            "average": statistics.mean(m_values)
        })
    
    sample_paths = random.sample(all_simulation_paths, min(num_sample_paths, len(all_simulation_paths)))

    return {
        "monthly_results": monthly_results,
        "sample_paths": sample_paths
    }
