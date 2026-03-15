import random
import math
import statistics
from datetime import datetime

def get_gaussian_random():
    u1, u2 = random.random(), random.random()
    return math.sqrt(-2.0 * math.log(u1)) * math.cos(2.0 * math.pi * u2)

def get_normal_random(mean, std_dev):
    return mean + get_gaussian_random() * std_dev

def run_monte_carlo_simulation(investment_data1, investment_data2, existing_savings, life_events, 
                               market_event_enabled=False, withdrawal_settings=None, total_simulation_years=50,
                               num_simulations=5000, num_sample_paths=100):
    all_simulation_paths = []
    total_months = total_simulation_years * 12

    # ライフイベント
    life_events_by_month = {}
    for event in life_events:
        m = int(event['year'] * 12)
        if m not in life_events_by_month: life_events_by_month[m] = 0
        life_events_by_month[m] += event['amount']

    # 取り崩し設定
    w_start_m = int(withdrawal_settings.get('start_year', 0) * 12) if withdrawal_settings else -1
    w_monthly = float(withdrawal_settings.get('monthly', 0)) * 10000 if withdrawal_settings else 0

    def get_monthly_changes(change_settings):
        changes = {}
        for s in change_settings: changes[int(s['year'] * 12)] = s['monthly']
        return changes

    inv1_changes = get_monthly_changes(investment_data1.get('change_settings', []))
    inv2_changes = get_monthly_changes(investment_data2.get('change_settings', []))

    r1_m_mean = (investment_data1['return'] / 100) / 12
    s1_m_std = (investment_data1['risk'] / 100) / math.sqrt(12)
    r2_m_mean = (investment_data2['return'] / 100) / 12
    s2_m_std = (investment_data2['risk'] / 100) / math.sqrt(12)

    inv1_is_lower = investment_data1['return'] <= investment_data2['return']
    MONTHLY_CRASH_PROB = 1.0 / 67.0

    for _ in range(num_simulations):
        val1, val2 = investment_data1['initial'], investment_data2['initial']
        path = [val1 + val2 + existing_savings]
        cur_m1, cur_m2 = investment_data1['monthly'], investment_data2['monthly']

        for m in range(1, total_months + 1):
            if m in inv1_changes: cur_m1 = inv1_changes[m]
            if m in inv2_changes: cur_m2 = inv2_changes[m]

            # 1. リターン適用
            m_ret1, m_ret2 = get_normal_random(r1_m_mean, s1_m_std), get_normal_random(r2_m_mean, s2_m_std)
            if market_event_enabled and random.random() < MONTHLY_CRASH_PROB:
                cf = 1.0 - (random.uniform(20.0, 55.0) / 100.0)
                val1 *= cf; val2 *= cf
            else:
                val1 *= (1 + m_ret1); val2 *= (1 + m_ret2)

            # 2. 積立 (無期限に継続。停止したい場合は設定で0円にする)
            val1 += cur_m1
            val2 += cur_m2

            # 3. 取り崩し & ライフイベント
            expense = life_events_by_month.get(m, 0)
            if m >= w_start_m: expense += w_monthly

            if expense > 0:
                if inv1_is_lower:
                    if val1 >= expense: val1 -= expense
                    else: rem = expense - val1; val1 = 0; val2 = max(0, val2 - rem)
                else:
                    if val2 >= expense: val2 -= expense
                    else: rem = expense - val2; val2 = 0; val1 = max(0, val1 - rem)
            
            path.append(val1 + val2 + existing_savings)
        all_simulation_paths.append(path)

    def get_p(data, p):
        k = (len(data)-1) * p / 100.0; f = math.floor(k); c = math.ceil(k)
        return data[int(f)] if f == c else data[int(f)]*(c-k) + data[int(c)]*(k-f)

    results = []
    for m in range(total_months + 1):
        vals = sorted([p[m] for p in all_simulation_paths])
        results.append({
            "month": m, "min": vals[0], "p10": get_p(vals, 10), "p25": get_p(vals, 25),
            "median": statistics.median(vals), "p75": get_p(vals, 75), "p90": get_p(vals, 90),
            "max": vals[-1], "average": statistics.mean(vals)
        })
    return {"monthly_results": results, "sample_paths": random.sample(all_simulation_paths, min(num_sample_paths, len(all_simulation_paths)))}
