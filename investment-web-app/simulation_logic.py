import random
import math
import statistics

def get_gaussian_random():
    """Box-Muller変換を使用して、標準正規分布に従う乱数を生成"""
    u1, u2 = random.random(), random.random()
    return math.sqrt(-2.0 * math.log(u1)) * math.cos(2.0 * math.pi * u2)

def get_normal_random(mean, std_dev):
    """正規分布に従う乱数を生成"""
    return mean + get_gaussian_random() * std_dev

def run_monte_carlo_simulation(investment_data1, investment_data2, existing_savings, life_events, num_simulations=5000, num_sample_paths=100):
    all_simulation_paths = []
    total_investment_period = max(investment_data1['period'], investment_data2['period'])

    life_events_by_year = {}
    for event in life_events:
        year = event['year']
        amount = event['amount']
        if year not in life_events_by_year:
            life_events_by_year[year] = 0
        life_events_by_year[year] += amount

    def get_monthly_changes_by_year(change_settings):
        changes = {}
        sorted_settings = sorted(change_settings, key=lambda x: x['year'])
        for setting in sorted_settings:
            changes[setting['year']] = setting['monthly']
        return changes

    inv1_monthly_changes = get_monthly_changes_by_year(investment_data1.get('change_settings', []))
    inv2_monthly_changes = get_monthly_changes_by_year(investment_data2.get('change_settings', []))

    for _ in range(num_simulations):
        # 0年目の値を初期値とする
        val1 = investment_data1['initial']
        val2 = investment_data2['initial']
        
        path1_values = [val1]
        path2_values = [val2]

        monthly1 = investment_data1['monthly']
        for year_idx in range(investment_data1['period']):
            current_year = year_idx + 1 # 1年目からperiod年目まで
            if current_year in inv1_monthly_changes:
                monthly1 = inv1_monthly_changes[current_year]

            annual_return = get_normal_random(investment_data1['return'] / 100, investment_data1['risk'] / 100)
            val1 = val1 * (1 + annual_return) + (monthly1 * 12 * (1 + annual_return / 2))
            path1_values.append(val1)

        monthly2 = investment_data2['monthly']
        for year_idx in range(investment_data2['period']):
            current_year = year_idx + 1 # 1年目からperiod年目まで
            if current_year in inv2_monthly_changes:
                monthly2 = inv2_monthly_changes[current_year]
            
            annual_return = get_normal_random(investment_data2['return'] / 100, investment_data2['risk'] / 100)
            val2 = val2 * (1 + annual_return) + (monthly2 * 12 * (1 + annual_return / 2))
            path2_values.append(val2)
        
        combined_path = []
        # 0年目の合計値を設定
        initial_total = investment_data1['initial'] + investment_data2['initial'] + existing_savings
        
        # ライフイベントは0年目には適用しない（通常1年後以降に発生すると仮定）
        # ただし、もし0年目のライフイベントも考慮する場合はここで調整が必要
        combined_path.append(initial_total) # 0年目のデータ

        for year_idx in range(total_investment_period): # 1年目からperiod年目まで
            # path_valuesは0年目からデータを持つため、year_idx+1でアクセス
            year_val1 = path1_values[year_idx+1] if year_idx < investment_data1['period'] else path1_values[investment_data1['period']]
            year_val2 = path2_values[year_idx+1] if year_idx < investment_data2['period'] else path2_values[investment_data2['period']]
            
            current_year_total = year_val1 + year_val2 + existing_savings

            if (year_idx + 1) in life_events_by_year: # ライフイベントは1年後から
                current_year_total -= life_events_by_year[year_idx + 1]

            combined_path.append(current_year_total)
        all_simulation_paths.append(combined_path)

    yearly_results = []
    # 0年目の結果を計算して追加
    initial_year_values = [path[0] for path in all_simulation_paths]
    initial_year_values.sort()

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
        "year": 0,
        "min": initial_year_values[0],
        "p10": get_percentile(initial_year_values, 10),
        "p25": get_percentile(initial_year_values, 25),
        "median": statistics.median(initial_year_values) if initial_year_values else 0,
        "p75": get_percentile(initial_year_values, 75),
        "p90": get_percentile(initial_year_values, 90),
        "max": initial_year_values[-1],
        "average": statistics.mean(initial_year_values) if initial_year_values else 0
    })

    for year_idx in range(total_investment_period): # 1年目からperiod年目まで
        year_values = sorted([path[year_idx+1] for path in all_simulation_paths]) # path[year_idx+1] に変更
        
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
    
    # ランダムなサンプルパスを選択
    sample_paths = random.sample(all_simulation_paths, min(num_sample_paths, len(all_simulation_paths)))

    return {
        "yearly_results": yearly_results,
        "sample_paths": sample_paths
    }
