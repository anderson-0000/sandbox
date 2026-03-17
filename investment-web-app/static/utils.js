// static/utils.js
export const formatCurrency = (value) => {
    return new Intl.NumberFormat('ja-JP', { style: 'currency', currency: 'JPY' }).format(Math.round(value));
};

export const formatNumber = (value) => {
    return new Intl.NumberFormat('ja-JP').format(Math.round(value));
};

export const createLifeEventItemHtml = (year = '', amount = '') => {
    return `
        <div class="life-event-item">
            <label>何年後:</label>
            <input type="number" class="life-event-year" min="1" value="${year}">
            <label>金額 (万円):</label>
            <input type="number" class="life-event-amount" min="0" value="${amount}">
            <button type="button" class="remove-life-event">削除</button>
        </div>
    `;
};

export const createChangeSettingItemHtml = (year = '', monthly = '') => {
    return `
        <div class="change-setting-item">
            <label>何年後:</label>
            <input type="number" class="change-setting-year" min="1" value="${year}">
            <label>変更後の積立額 (万円):</label>
            <input type="number" class="change-setting-monthly" min="0" value="${monthly}">
            <button type="button" class="remove-change-setting">削除</button>
        </div>
    `;
};

export const createWithdrawalSettingItemHtml = (year = '', value = '', type = 'amount') => {
    return `
        <div class="withdrawal-setting-item" style="display: flex; gap: 10px; margin-bottom: 10px; align-items: center; border-bottom: 1px solid #eee; padding-bottom: 10px;">
            <div style="flex: 1;">
                <label style="display: block; margin-bottom: 5px; font-weight: bold; color: #555;">何年後から:</label>
                <input type="number" class="withdrawal-year" min="0" value="${year}" style="width: 100%; height: 40px; padding: 0 10px; border: 1px solid #ccc; border-radius: 4px; font-size: 1em; box-sizing: border-box; margin: 0;">
            </div>
            <div style="flex: 1;">
                <label style="display: block; margin-bottom: 5px; font-weight: bold; color: #555;">毎月の取り崩し額/率:</label>
                <div style="display: flex; gap: 0; align-items: stretch;">
                    <input type="number" class="withdrawal-value" min="0" step="0.1" value="${value}" style="flex: 2; height: 40px; padding: 0 10px; border: 1px solid #ccc; border-right: none; border-radius: 4px 0 0 4px; font-size: 1em; box-sizing: border-box; margin: 0;">
                    <select class="withdrawal-type" style="flex: 1; height: 40px; padding: 0 10px; border: 1px solid #ccc; border-radius: 0 4px 4px 0; font-size: 1em; background-color: #fff; box-sizing: border-box; cursor: pointer; -webkit-appearance: none; -moz-appearance: none; appearance: none; background-image: url('data:image/svg+xml;charset=US-ASCII,%3Csvg%20xmlns%3D%22http%3A//www.w3.org/2000/svg%22%20width%3D%22292.4%22%20height%3D%22292.4%22%3E%3Cpath%20fill%3D%22%23007CB2%22%20d%3D%22M287%2069.4a17.6%2017.6%200%200%200-13-5.4H18.4c-5%200-9.3%201.8-12.9%205.4A17.6%2017.6%200%200%200%200%2082.2c0%205%201.8%209.3%205.4%2012.9l128%20127.9c3.6%203.6%207.8%205.4%2012.8%205.4s9.2-1.8%2012.8-5.4L287%2095c3.5-3.5%205.4-7.8%205.4-12.8%200-5-1.9-9.2-5.5-12.8z%22/%3E%3C/svg%3E'); background-repeat: no-repeat; background-position: right 10px top 50%; background-size: 12px auto; margin: 0;">
                        <option value="amount" ${type === 'amount' ? 'selected' : ''}>万円</option>
                        <option value="percent" ${type === 'percent' ? 'selected' : ''}>%</option>
                    </select>
                </div>
            </div>
            <button type="button" class="remove-withdrawal-setting" style="margin-top: 25px;">削除</button>
        </div>
    `;
};