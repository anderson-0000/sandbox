// static/dom_handlers.js
import { createLifeEventItemHtml } from './utils.js';

export const populateForm = (prefix, data) => {
    document.getElementById(`${prefix}_initial`).value = data.initial;
    document.getElementById(`${prefix}_monthly`).value = data.monthly;
    document.getElementById(`${prefix}_return`).value = data.return;
    document.getElementById(`${prefix}_risk`).value = data.risk;
    document.getElementById(`${prefix}_period`).value = data.period;
    document.getElementById(`${prefix}_change_year`).value = data.change_year || '';
    document.getElementById(`${prefix}_changed_monthly`).value = data.changed_monthly || '';
};

export const addLifeEventItem = (lifeEventsContainer, year = '', amount = '') => {
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = createLifeEventItemHtml(year, amount);
    const newItem = tempDiv.firstElementChild;
    lifeEventsContainer.appendChild(newItem);
};

export const getInvestmentData = (prefix) => {
    const changeYearElement = document.getElementById(`${prefix}_change_year`);
    const changedMonthlyElement = document.getElementById(`${prefix}_changed_monthly`);

    return {
        initial: parseFloat(document.getElementById(`${prefix}_initial`).value),
        monthly: parseFloat(document.getElementById(`${prefix}_monthly`).value),
        return: parseFloat(document.getElementById(`${prefix}_return`).value),
        risk: parseFloat(document.getElementById(`${prefix}_risk`).value),
        period: parseInt(document.getElementById(`${prefix}_period`).value, 10),
        change_year: changeYearElement && changeYearElement.value ? parseInt(changeYearElement.value, 10) : 0,
        changed_monthly: changedMonthlyElement && changedMonthlyElement.value ? parseFloat(changedMonthlyElement.value) : 0,
    };
};

export const getLifeEventsData = () => {
    const events = [];
    document.querySelectorAll('.life-event-item').forEach(item => {
        const year = parseInt(item.querySelector('.life-event-year').value, 10);
        const amount = parseFloat(item.querySelector('.life-event-amount').value);
        if (!isNaN(year) && year > 0 && !isNaN(amount)) {
            events.push({ year, amount });
        }
    });
    return events;
};