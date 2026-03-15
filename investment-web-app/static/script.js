// static/script.js
import { formatCurrency, formatNumber } from './utils.js';
import { populateForm, addLifeEventItem, addChangeSettingItem, getInvestmentData, getLifeEventsData } from './dom_handlers.js';

document.addEventListener('DOMContentLoaded', () => {
    const simulateButton = document.getElementById('simulateButton');
    const downloadChartButton = document.getElementById('downloadChartButton');
    const resultsTableBody = document.querySelector('#resultsTable tbody');
    const summaryTableBody = document.querySelector('#summaryTable tbody');
    const chartCanvas = document.getElementById('simulationChart');

    const toggleSamplePaths = document.getElementById('toggleSamplePaths');
    const samplePathsSlider = document.getElementById('samplePathsSlider');
    const samplePathsCountSpan = document.getElementById('samplePathsCount');

    // Life Event elements
    const addLifeEventButton = document.getElementById('addLifeEventButton');
    const lifeEventsContainer = document.getElementById('life_events_container');

    // Investment elements
    const addInv1ChangeSettingButton = document.getElementById('addInv1ChangeSettingButton');
    const inv1ChangeSettingsContainer = document.getElementById('inv1_change_settings_container');
    const addInv2ChangeSettingButton = document.getElementById('addInv2ChangeSettingButton');
    const inv2ChangeSettingsContainer = document.getElementById('inv2_change_settings_container');

    // Define and register custom plugin for chart background
    const customCanvasBackgroundColor = {
        id: 'customCanvasBackgroundColor',
        beforeDraw: (chart, args, options) => {
            const {ctx} = chart; ctx.save();
            ctx.globalCompositeOperation = 'destination-over';
            ctx.fillStyle = options.color || '#fff';
            ctx.fillRect(0, 0, chart.width, chart.height);
            ctx.restore();
        }
    };
    Chart.register(customCanvasBackgroundColor);

    let simulationChart;
    let globalMonthlyResults = [];
    let globalSamplePaths = [];

    const loadDefaultValues = async () => {
        try {
            const response = await fetch('/defaults');
            if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
            const defaults = await response.json();
            populateForm('inv1', defaults.inv1);
            populateForm('inv2', defaults.inv2);
            document.getElementById('existing_savings').value = defaults.existing_savings || 0;
            document.getElementById('current_age').value = defaults.current_age || 33;
            document.getElementById('total_period').value = defaults.total_period || 50;
            document.getElementById('crash_enabled').checked = defaults.crash_enabled || false;
            document.getElementById('withdrawal_monthly').value = defaults.withdrawal_monthly || 0;
            document.getElementById('withdrawal_start').value = defaults.withdrawal_start || 10;

            lifeEventsContainer.innerHTML = '';
            if (defaults.life_events && defaults.life_events.length > 0) {
                defaults.life_events.forEach(event => addLifeEventItem(lifeEventsContainer, event.year, event.amount));
            } else {
                addLifeEventItem(lifeEventsContainer);
            }
        } catch (error) { console.error('Error loading defaults:', error); }
    };

    const initChart = () => {
        if (simulationChart) simulationChart.destroy();
        simulationChart = new Chart(chartCanvas, {
            type: 'line',
            data: { labels: [], datasets: [] },
            options: {
                responsive: true, maintainAspectRatio: false,
                plugins: {
                    customCanvasBackgroundColor: { color: 'white' },
                    legend: { position: 'top' },
                    title: { display: true, text: 'ポートフォリオ価値の推移' },
                    tooltip: {
                        callbacks: {
                            label: (context) => {
                                let label = context.dataset.label || '';
                                if (label) label += ': ';
                                if (context.parsed.y !== null) label += formatCurrency(context.parsed.y);
                                return label;
                            }
                        }
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        title: { display: true, text: 'ポートフォリオ価値 (万円)' },
                        ticks: { callback: (value) => formatNumber(value / 10000) }
                    },
                    x: {
                        title: { display: true, text: '年/月 (年齢)' },
                        grid: {
                            display: true,
                            color: (context) => (context.tick && context.tick.label && context.tick.label.includes('/1 ')) ? 'rgba(0, 0, 0, 0.1)' : 'rgba(0, 0, 0, 0.05)',
                            lineWidth: 1
                        },
                        ticks: {
                            maxRotation: 45, minRotation: 45, autoSkip: false,
                            callback: function(value, index) {
                                const label = this.getLabelForValue(value);
                                return (index === 0 || label.includes('/1 ')) ? label : null;
                            }
                        }
                    },
                },
            },
        });
    };

    const updateChartData = () => {
        const now = new Date();
        const startYear = now.getFullYear();
        const startMonth = now.getMonth() + 1;
        const currentAge = parseInt(document.getElementById('current_age').value, 10) || 0;

        const chartLabels = globalMonthlyResults.map(res => {
            const totalMonths = startMonth - 1 + res.month;
            const y = startYear + Math.floor(totalMonths / 12);
            const m = (totalMonths % 12) + 1;
            const age = currentAge + Math.floor(totalMonths / 12);
            return `${y}/${m} (${age}歳)`;
        });

        const datasets = [
            {
                label: '中央値 (50%)', data: globalMonthlyResults.map(res => res.median),
                borderColor: 'rgb(75, 192, 192)', backgroundColor: 'rgba(75, 192, 192, 0.5)',
                tension: 0.1, pointRadius: 0,
            },
            {
                label: '90% パーセンタイル', data: globalMonthlyResults.map(res => res.p90),
                borderColor: 'rgb(53, 162, 235)', backgroundColor: 'rgba(53, 162, 235, 0.5)',
                tension: 0.1, pointRadius: 0,
            },
            {
                label: '10% パーセンタイル', data: globalMonthlyResults.map(res => res.p10),
                borderColor: 'rgb(255, 99, 132)', backgroundColor: 'rgba(255, 99, 132, 0.5)',
                tension: 0.1, pointRadius: 0,
            },
        ];

        if (toggleSamplePaths.checked && globalSamplePaths.length > 0) {
            const numPathsToShow = parseInt(samplePathsSlider.value, 10);
            for (let i = 0; i < Math.min(numPathsToShow, globalSamplePaths.length); i++) {
                const r = Math.floor(Math.random() * 200) + 50;
                const g = Math.floor(Math.random() * 200) + 50;
                const b = Math.floor(Math.random() * 200) + 50;
                datasets.push({
                    label: `サンプルパス ${i + 1}`, data: globalSamplePaths[i],
                    borderColor: `rgba(${r}, ${g}, ${b}, 0.7)`, backgroundColor: `rgba(${r}, ${g}, ${b}, 0.3)`,
                    borderDash: [5, 5], tension: 0.1, pointRadius: 0, hidden: false,
                });
            }
        }

        simulationChart.data.labels = chartLabels;
        simulationChart.data.datasets = datasets;
        simulationChart.update();
    };

    initChart();
    loadDefaultValues();

    addLifeEventButton.addEventListener('click', () => addLifeEventItem(lifeEventsContainer));
    lifeEventsContainer.addEventListener('click', (event) => {
        if (event.target.classList.contains('remove-life-event')) event.target.closest('.life-event-item').remove();
    });

    addInv1ChangeSettingButton.addEventListener('click', () => addChangeSettingItem(inv1ChangeSettingsContainer));
    inv1ChangeSettingsContainer.addEventListener('click', (event) => {
        if (event.target.classList.contains('remove-change-setting')) event.target.closest('.change-setting-item').remove();
    });

    addInv2ChangeSettingButton.addEventListener('click', () => addChangeSettingItem(inv2ChangeSettingsContainer));
    inv2ChangeSettingsContainer.addEventListener('click', (event) => {
        if (event.target.classList.contains('remove-change-setting')) event.target.closest('.change-setting-item').remove();
    });

    downloadChartButton.addEventListener('click', () => {
        if (simulationChart) {
            const image = simulationChart.toBase64Image('image/jpeg', 1.0);
            const a = document.createElement('a');
            a.href = image; a.download = 'simulation_chart.jpeg';
            document.body.appendChild(a); a.click(); document.body.removeChild(a);
        }
    });

    toggleSamplePaths.addEventListener('change', updateChartData);
    document.getElementById('current_age').addEventListener('input', updateChartData);
    samplePathsSlider.addEventListener('input', () => {
        samplePathsCountSpan.textContent = samplePathsSlider.value;
        updateChartData();
    });

    simulateButton.addEventListener('click', async () => {
        simulateButton.disabled = true;
        downloadChartButton.style.display = 'none';

        const inv1Data = getInvestmentData('inv1');
        const inv2Data = getInvestmentData('inv2');
        const savings = parseFloat(document.getElementById('existing_savings').value);
        const totalPeriod = parseInt(document.getElementById('total_period').value, 10);
        const lifeEvents = getLifeEventsData();
        const crashEnabled = document.getElementById('crash_enabled').checked;
        const wMonthly = parseFloat(document.getElementById('withdrawal_monthly').value);
        const wStart = parseInt(document.getElementById('withdrawal_start').value, 10);

        try {
            const response = await fetch('/simulate', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    investment1: inv1Data, investment2: inv2Data,
                    existing_savings: savings, total_period: totalPeriod,
                    life_events: lifeEvents, market_event_enabled: crashEnabled,
                    withdrawal_monthly: wMonthly, withdrawal_start: wStart
                }),
            });

            if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);

            const output = await response.json();
            globalMonthlyResults = output.monthly_results;
            globalSamplePaths = output.sample_paths;

            resultsTableBody.innerHTML = '';
            const now = new Date();
            const startYear = now.getFullYear();
            const startMonth = now.getMonth() + 1;
            const currentAgeValue = parseInt(document.getElementById('current_age').value, 10) || 0;

            globalMonthlyResults.forEach(res => {
                const totalMonths = startMonth - 1 + res.month;
                const y = startYear + Math.floor(totalMonths / 12);
                const m = (totalMonths % 12) + 1;
                const age = currentAgeValue + Math.floor(totalMonths / 12);

                if (res.month === 0 || m === 1) {
                    const row = resultsTableBody.insertRow();
                    row.insertCell().textContent = `${y}/${m} (${age}歳)`;
                    row.insertCell().textContent = formatCurrency(res.min * 10000);
                    row.insertCell().textContent = formatCurrency(res.p10 * 10000);
                    row.insertCell().textContent = formatCurrency(res.median * 10000);
                    row.insertCell().textContent = formatCurrency(res.p90 * 10000);
                    row.insertCell().textContent = formatCurrency(res.max * 10000);
                    row.insertCell().textContent = formatCurrency(res.average * 10000);
                }
            });

            summaryTableBody.innerHTML = '';
            if (globalMonthlyResults.length > 0) {
                const last = globalMonthlyResults[globalMonthlyResults.length - 1];
                const row = summaryTableBody.insertRow();
                row.insertCell().textContent = formatCurrency(last.min * 10000);
                row.insertCell().textContent = formatCurrency(last.p10 * 10000);
                row.insertCell().textContent = formatCurrency(last.median * 10000);
                row.insertCell().textContent = formatCurrency(last.p90 * 10000);
                row.insertCell().textContent = formatCurrency(last.max * 10000);
                row.insertCell().textContent = formatCurrency(last.average * 10000);
            }
            updateChartData();
            downloadChartButton.style.display = 'block';
        } catch (error) {
            console.error('Error:', error);
            alert('エラーが発生しました: ' + error.message);
        } finally { simulateButton.disabled = false; }
    });
});
