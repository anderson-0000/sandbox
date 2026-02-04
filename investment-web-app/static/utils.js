// static/utils.js
export const formatCurrency = (value) => {
    return new Intl.NumberFormat('ja-JP', { style: 'currency', currency: 'JPY' }).format(Math.round(value));
};

export const createLifeEventItemHtml = (year = '', amount = '') => {
    return `
        <div class="life-event-item">
            <label>何年後:</label>
            <input type="number" class="life-event-year" min="1" value="${year}">
            <label>金額 (円):</label>
            <input type="number" class="life-event-amount" min="0" value="${amount}">
            <button type="button" class="remove-life-event">削除</button>
        </div>
    `;
};