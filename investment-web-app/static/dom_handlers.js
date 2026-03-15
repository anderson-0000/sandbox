// static/dom_handlers.js
import { createLifeEventItemHtml, createChangeSettingItemHtml } from './utils.js'; // 変更

export const populateForm = (prefix, data) => {
    document.getElementById(`${prefix}_initial`).value = data.initial;
    document.getElementById(`${prefix}_monthly`).value = data.monthly;
    document.getElementById(`${prefix}_return`).value = data.return;
    document.getElementById(`${prefix}_risk`).value = data.risk;

    // 積立額変更設定を初期表示 (変更)
    const changeSettingsContainer = document.getElementById(`${prefix}_change_settings_container`);
    if (changeSettingsContainer) {
        changeSettingsContainer.innerHTML = ''; // Clear existing
        if (data.change_settings && data.change_settings.length > 0) {
            data.change_settings.forEach(setting => {
                addChangeSettingItem(changeSettingsContainer, setting.year, setting.monthly);
            });
        } else {
            addChangeSettingItem(changeSettingsContainer); // Add one empty by default
        }
    }
};

export const addLifeEventItem = (lifeEventsContainer, year = '', amount = '') => {
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = createLifeEventItemHtml(year, amount);
    const newItem = tempDiv.firstElementChild;
    lifeEventsContainer.appendChild(newItem);
};

export const addChangeSettingItem = (container, year = '', monthly = '') => { // 追加
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = createChangeSettingItemHtml(year, monthly);
    const newItem = tempDiv.firstElementChild;
    container.appendChild(newItem);
};

export const getInvestmentData = (prefix) => { // 変更
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
        change_settings: changeSettings, // 複数の変更設定を配列として追加
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