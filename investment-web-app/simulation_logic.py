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

def run_monte_carlo_simulation(investment_data1, investment_data2, existing_savings, life_events, num_simulations=5000, num_sample_paths=100):
    all_simulation_paths = []
    total_investment_period_years = max(investment_data1['period'], investment_data2['period'])
    total_months = total_investment_period_years * 12

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

    # 年利・年リスクを月次リターン・月次リスクに変換
    # 簡略化のため r_monthly = r_annual / 12, sigma_monthly = sigma_annual / sqrt(12) を使用
    r1_monthly_mean = (investment_data1['return'] / 100) / 12
    s1_monthly_std = (investment_data1['risk'] / 100) / math.sqrt(12)
    
    r2_monthly_mean = (investment_data2['return'] / 100) / 12
    s2_monthly_std = (investment_data2['risk'] / 100) / math.sqrt(12)

    for _ in range(num_simulations):
        val1 = investment_data1['initial']
        val2 = investment_data2['initial']
        
        path1_values = [val1]
        path2_values = [val2]

        # 投資1のシミュレーション
        current_monthly_contribution1 = investment_data1['monthly']
        for m in range(1, total_months + 1):
            if m in inv1_monthly_changes:
                current_monthly_contribution1 = inv1_monthly_changes[m]
            
            if m <= investment_data1['period'] * 12:
                # 複利計算 (毎月の積立は月末に行うと仮定し、その月のリターンを乗算)
                monthly_return = get_normal_random(r1_monthly_mean, s1_monthly_std)
                val1 = val1 * (1 + monthly_return) + current_monthly_contribution1
            else:
                # 期間終了後は運用のみ継続（もし必要なら。現状は期間終了後のデータも必要）
                monthly_return = get_normal_random(r1_monthly_mean, s1_monthly_std)
                val1 = val1 * (1 + monthly_return)
            path1_values.append(val1)

        # 投資2のシミュレーション
        current_monthly_contribution2 = investment_data2['monthly']
        for m in range(1, total_months + 1):
            if m in inv2_monthly_changes:
                current_monthly_contribution2 = inv2_monthly_changes[m]
            
            if m <= investment_data2['period'] * 12:
                monthly_return = get_normal_random(r2_monthly_mean, s2_monthly_std)
                val2 = val2 * (1 + monthly_return) + current_monthly_contribution2
            else:
                monthly_return = get_normal_random(r2_monthly_mean, s2_monthly_std)
                val2 = val2 * (1 + monthly_return)
            path2_values.append(val2)
        
        combined_path = []
        initial_total = investment_data1['initial'] + investment_data2['initial'] + existing_savings
        combined_path.append(initial_total) # 0ヶ月目

        for m in range(1, total_months + 1):
            current_month_total = path1_values[m] + path2_values[m] + existing_savings
            if m in life_events_by_month:
                current_month_total -= life_events_by_month[m]
            combined_path.append(current_month_total)
        
        all_simulation_paths.append(combined_path)

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

    # 月次結果の集計
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
