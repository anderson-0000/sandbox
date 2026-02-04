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

def run_monte_carlo_simulation(investment_data1, investment_data2, existing_savings, life_events, num_simulations=5000, num_sample_paths=5):
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
        path1_values = [investment_data1['initial']]
        path2_values = [investment_data2['initial']]

        val1 = investment_data1['initial']
        monthly1 = investment_data1['monthly']
        for year_idx in range(investment_data1['period']):
            current_year = year_idx + 1
            if current_year in inv1_monthly_changes:
                monthly1 = inv1_monthly_changes[current_year]

            annual_return = get_normal_random(investment_data1['return'] / 100, investment_data1['risk'] / 100)
            val1 = val1 * (1 + annual_return) + (monthly1 * 12 * (1 + annual_return / 2))
            path1_values.append(val1)

        val2 = investment_data2['initial']
        monthly2 = investment_data2['monthly']
        for year_idx in range(investment_data2['period']):
            current_year = year_idx + 1
            if current_year in inv2_monthly_changes:
                monthly2 = inv2_monthly_changes[current_year]
            
            annual_return = get_normal_random(investment_data2['return'] / 100, investment_data2['risk'] / 100)
            val2 = val2 * (1 + annual_return) + (monthly2 * 12 * (1 + annual_return / 2))
            path2_values.append(val2)
        
        combined_path = []
        for year_idx in range(total_investment_period):
            year_val1 = path1_values[year_idx+1] if year_idx < investment_data1['period'] else path1_values[investment_data1['period']]
            year_val2 = path2_values[year_idx+1] if year_idx < investment_data2['period'] else path2_values[investment_data2['period']]
            
            current_year_total = year_val1 + year_val2 + existing_savings

            if (year_idx + 1) in life_events_by_year:
                current_year_total -= life_events_by_year[year_idx + 1]

            combined_path.append(current_year_total)
        all_simulation_paths.append(combined_path)

    yearly_results = []
    for year_idx in range(total_investment_period):
        year_values = sorted([path[year_idx] for path in all_simulation_paths])
        
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
    
    # ランダムなサンプルパスを選択
    sample_paths = random.sample(all_simulation_paths, min(num_sample_paths, len(all_simulation_paths)))

    return {
        "yearly_results": yearly_results,
        "sample_paths": sample_paths
    }
