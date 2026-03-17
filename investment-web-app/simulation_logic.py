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
    total_months = total_simulation_years * 12

    # ライフイベント
    life_events_by_month = {}
    for event in life_events:
        m = int(event['year'] * 12)
        if m not in life_events_by_month: life_events_by_month[m] = 0
        life_events_by_month[m] += event['amount']

    # 取り崩し設定の整理
    sorted_withdrawals = sorted(withdrawal_settings or [], key=lambda x: x['year'])
    
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

    # 全シミュレーションの状態を保持
    val1s = [float(investment_data1['initial'])] * num_simulations
    val2s = [float(investment_data2['initial'])] * num_simulations
    
    # 全パスの履歴
    all_paths = [[v1 + v2 + existing_savings] for v1, v2 in zip(val1s, val2s)]
    withdrawal_history = [0.0]
    
    cur_m1 = [float(investment_data1['monthly'])] * num_simulations
    cur_m2 = [float(investment_data2['monthly'])] * num_simulations

    for m in range(1, total_months + 1):
        # 1. 共通の取り崩し額の計算 (全パスの中央値資産額を元にする)
        # 前月末時点の資産額リスト
        current_totals = sorted([v1 + v2 + existing_savings for v1, v2 in zip(val1s, val2s)])
        median_assets = current_totals[num_simulations // 2]
        
        withdrawal_monthly = 0.0
        active_setting = None
        for s in sorted_withdrawals:
            if m >= int(s['year'] * 12): active_setting = s
            else: break
        
        if active_setting:
            if active_setting['type'] == 'amount':
                withdrawal_monthly = float(active_setting['value']) * 10000
            elif active_setting['type'] == 'percent':
                # 指定された月利(%)を適用
                withdrawal_monthly = median_assets * (float(active_setting['value']) / 100.0)
        
        withdrawal_history.append(withdrawal_monthly)
        expense = life_events_by_month.get(m, 0) + withdrawal_monthly

        # 2. 各パスを更新
        for i in range(num_simulations):
            # 積立額の変更チェック
            if m in inv1_changes: cur_m1[i] = float(inv1_changes[m])
            if m in inv2_changes: cur_m2[i] = float(inv2_changes[m])

            # リターン適用
            m_ret1, m_ret2 = get_normal_random(r1_m_mean, s1_m_std), get_normal_random(r2_m_mean, s2_m_std)
            if market_event_enabled and random.random() < MONTHLY_CRASH_PROB:
                cf = 1.0 - (random.uniform(20.0, 55.0) / 100.0)
                val1s[i] *= cf; val2s[i] *= cf
            else:
                val1s[i] *= (1 + m_ret1); val2s[i] *= (1 + m_ret2)

            # 積立
            val1s[i] += cur_m1[i]
            val2s[i] += cur_m2[i]

            # 取り崩し & ライフイベント
            if expense > 0:
                if inv1_is_lower:
                    if val1s[i] >= expense: val1s[i] -= expense
                    else: rem = expense - val1s[i]; val1s[i] = 0; val2s[i] = max(0.0, val2s[i] - rem)
                else:
                    if val2s[i] >= expense: val2s[i] -= expense
                    else: rem = expense - val2s[i]; val2s[i] = 0; val1s[i] = max(0.0, val1s[i] - rem)
            
            all_paths[i].append(val1s[i] + val2s[i] + existing_savings)

    def get_p(data, p):
        k = (len(data)-1) * p / 100.0; f = math.floor(k); c = math.ceil(k)
        return data[int(f)] if f == c else data[int(f)]*(c-k) + data[int(c)]*(k-f)

    results = []
    for m in range(total_months + 1):
        vals = sorted([p[m] for p in all_paths])
        results.append({
            "month": m, "min": vals[0], "p10": get_p(vals, 10), "p25": get_p(vals, 25),
            "median": statistics.median(vals), "p75": get_p(vals, 75), "p90": get_p(vals, 90),
            "max": vals[-1], "average": statistics.mean(vals),
            "withdrawal": withdrawal_history[m] # 計算に使用した共通の取り崩し額
        })
    
    sample_indices = random.sample(range(num_simulations), min(num_sample_paths, num_simulations))
    return {
        "monthly_results": results, 
        "sample_paths": [all_paths[i] for i in sample_indices]
    }
