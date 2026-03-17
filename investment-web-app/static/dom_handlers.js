// static/dom_handlers.js
import { createLifeEventItemHtml, createChangeSettingItemHtml, createWithdrawalSettingItemHtml } from './utils.js';

export const populateForm = (prefix, data) => {
    document.getElementById(`${prefix}_initial`).value = data.initial;
    document.getElementById(`${prefix}_monthly`).value = data.monthly;
    document.getElementById(`${prefix}_return`).value = data.return;
    document.getElementById(`${prefix}_risk`).value = data.risk;

    // 積立額変更設定を初期表示
    const changeSettingsContainer = document.getElementById(`${prefix}_change_settings_container`);
    if (changeSettingsContainer) {
        changeSettingsContainer.innerHTML = '';
        if (data.change_settings && data.change_settings.length > 0) {
            data.change_settings.forEach(setting => {
                addChangeSettingItem(changeSettingsContainer, setting.year, setting.monthly);
            });
        } else {
            addChangeSettingItem(changeSettingsContainer);
        }
    }
};

export const populateWithdrawalSettings = (container, settings) => {
    container.innerHTML = '';
    if (settings && settings.length > 0) {
        settings.forEach(s => addWithdrawalSettingItem(container, s.year, s.value, s.type));
    } else {
        addWithdrawalSettingItem(container);
    }
};

export const addLifeEventItem = (lifeEventsContainer, year = '', amount = '') => {
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = createLifeEventItemHtml(year, amount);
    const newItem = tempDiv.firstElementChild;
    lifeEventsContainer.appendChild(newItem);
};

export const addChangeSettingItem = (container, year = '', monthly = '') => {
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = createChangeSettingItemHtml(year, monthly);
    const newItem = tempDiv.firstElementChild;
    container.appendChild(newItem);
};

export const addWithdrawalSettingItem = (container, year = '', value = '', type = 'amount') => {
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = createWithdrawalSettingItemHtml(year, value, type);
    const newItem = tempDiv.firstElementChild;
    container.appendChild(newItem);
};

export const getInvestmentData = (prefix) => {
    const changeSettings = [];
    document.querySelectorAll(`#${prefix}_change_settings_container .change-setting-item`).forEach(item => {
        const year = parseInt(item.querySelector('.change-setting-year').value, 10);
        const monthly = parseFloat(item.querySelector('.change-setting-monthly').value);
        if (!isNaN(year) && year > 0 && !isNaN(monthly)) {
            changeSettings.push({ year, monthly });
        }
    });

    return {
        initial: parseFloat(document.getElementById(`${prefix}_initial`).value),
        monthly: parseFloat(document.getElementById(`${prefix}_monthly`).value),
        return: parseFloat(document.getElementById(`${prefix}_return`).value),
        risk: parseFloat(document.getElementById(`${prefix}_risk`).value),
        change_settings: changeSettings,
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

export const getWithdrawalSettingsData = () => {
    const settings = [];
    document.querySelectorAll('.withdrawal-setting-item').forEach(item => {
        const year = parseInt(item.querySelector('.withdrawal-year').value, 10);
        const value = parseFloat(item.querySelector('.withdrawal-value').value);
        const type = item.querySelector('.withdrawal-type').value;
        if (!isNaN(year) && !isNaN(value)) {
            settings.push({ year, value, type });
        }
    });
    return settings;
};