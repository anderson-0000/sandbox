// static/script.js
import { formatCurrency, formatNumber } from './utils.js';
import { populateForm, addLifeEventItem, addChangeSettingItem, getInvestmentData, getLifeEventsData } from './dom_handlers.js';

document.addEventListener('DOMContentLoaded', () => {
    const simulateButton = document.getElementById('simulateButton');
    const downloadChartButton = document.getElementById('downloadChartButton');
    const resultsTableBody = document.querySelector('#resultsTable tbody');
    const summaryTableBody = document.querySelector('#summaryTable tbody');
    const chartCanvas = document.getElementById('simulationChart');

    const toggleSamplePaths = document.getElementById('toggleSamplePaths'); // 追加
    const samplePathsSlider = document.getElementById('samplePathsSlider'); // 追加
    const samplePathsCountSpan = document.getElementById('samplePathsCount'); // 追加

    // Life Event elements
    const addLifeEventButton = document.getElementById('addLifeEventButton');
    const lifeEventsContainer = document.getElementById('life_events_container');

    // Investment 1 Change Setting elements
    const addInv1ChangeSettingButton = document.getElementById('addInv1ChangeSettingButton');
    const inv1ChangeSettingsContainer = document.getElementById('inv1_change_settings_container');

    // Investment 2 Change Setting elements
    const addInv2ChangeSettingButton = document.getElementById('addInv2ChangeSettingButton');
    const inv2ChangeSettingsContainer = document.getElementById('inv2_change_settings_container');


    // Define and register custom plugin for chart background
    const customCanvasBackgroundColor = {
        id: 'customCanvasBackgroundColor',
        beforeDraw: (chart, args, options) => {
            const {ctx} = chart;
            ctx.save();
            ctx.globalCompositeOperation = 'destination-over';
            ctx.fillStyle = options.color || '#fff'; // Use option color or default to white
            ctx.fillRect(0, 0, chart.width, chart.height);
            ctx.restore();
        }
    };
    Chart.register(customCanvasBackgroundColor); // Register the plugin

    let simulationChart; // To hold the Chart.js instance
    let globalMonthlyResults = []; // 全体の結果を保持 (月次)
    let globalSamplePaths = [];   // サンプルパスを保持

    // Function to load default values
    const loadDefaultValues = async () => {
        try {
            const response = await fetch('/defaults');
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const defaults = await response.json();
            populateForm('inv1', defaults.inv1);
            populateForm('inv2', defaults.inv2);
            // Set existing savings default
            document.getElementById('existing_savings').value = defaults.existing_savings || 0;
            document.getElementById('current_age').value = defaults.current_age || 30;
            document.getElementById('crash_enabled').checked = defaults.crash_enabled || false;
            document.getElementById('crash_year').value = defaults.crash_year || 5;
            document.getElementById('crash_interval').value = defaults.crash_interval || 10;
            document.getElementById('crash_rate').value = defaults.crash_rate || 30;

            // Populate life events
            lifeEventsContainer.innerHTML = ''; // Clear existing
            if (defaults.life_events && defaults.life_events.length > 0) {
                defaults.life_events.forEach(event => {
                    addLifeEventItem(lifeEventsContainer, event.year, event.amount);
                });
            } else {
                addLifeEventItem(lifeEventsContainer); // Add one empty by default
            }

            // populateForm で積立額変更設定も初期表示されるようになったため、ここでの処理は不要
        } catch (error) {
            console.error('Error loading default values:', error);
            alert('デフォルト値の読み込み中にエラーが発生しました。');
        }
    };

    // Initialize Chart.js
    const initChart = () => {
        if (simulationChart) {
            simulationChart.destroy(); // Destroy previous chart if exists
        }
        simulationChart = new Chart(chartCanvas, {
            type: 'line',
            data: {
                labels: [], // Will be filled with YYYY/MM
                datasets: [] // Will be filled with percentile data
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    customCanvasBackgroundColor: { // Plugin configuration
                        color: 'white'
                    },
                    legend: {
                        position: 'top',
                    },
                    title: {
                        display: true,
                        text: 'ポートフォリオ価値の推移',
                    },
                    tooltip: {
                        callbacks: {
                            label: function(context) {
                                let label = context.dataset.label || '';
                                if (label) {
                                    label += ': ';
                                }
                                if (context.parsed.y !== null) {
                                    label += formatCurrency(context.parsed.y);
                                }
                                return label;
                            }
                        }
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: 'ポートフォリオ価値 (万円)',
                        },
                        ticks: {
                            callback: function(value) {
                                return formatNumber(value / 10000);
                            }
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: '年/月 (年齢)',
                        },
                        grid: {
                            display: true,
                            color: function(context) {
                                if (context.tick && context.tick.label) {
                                    // ラベルが "/1 " を含む場合（1月）に線を強調 (例: 2027/1 (31歳))
                                    return context.tick.label.includes('/1 ') ? 'rgba(0, 0, 0, 0.1)' : 'rgba(0, 0, 0, 0.05)';
                                }
                                return 'rgba(0, 0, 0, 0.05)';
                            },
                            lineWidth: 1
                        },
                        ticks: {
                            maxRotation: 45,
                            minRotation: 45,
                            autoSkip: false,
                            callback: function(value, index, values) {
                                const label = this.getLabelForValue(value);
                                // 最初のデータ（現在月）または1月のラベルのみ表示
                                if (index === 0 || label.includes('/1 ')) {
                                    return label;
                                }
                                return null;
                            }
                        }
                    },
                },
            },
        });
    };

    // Function to update chart datasets based on results and controls
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
                label: '中央値 (50%)',
                data: globalMonthlyResults.map(res => res.median),
                borderColor: 'rgb(75, 192, 192)',
                backgroundColor: 'rgba(75, 192, 192, 0.5)',
                tension: 0.1,
                pointRadius: 0, // 点を非表示にしてスッキリさせる
            },
            {
                label: '90% パーセンタイル',
                data: globalMonthlyResults.map(res => res.p90),
                borderColor: 'rgb(53, 162, 235)',
                backgroundColor: 'rgba(53, 162, 235, 0.5)',
                tension: 0.1,
                pointRadius: 0,
            },
            {
                label: '10% パーセンタイル',
                data: globalMonthlyResults.map(res => res.p10),
                borderColor: 'rgb(255, 99, 132)',
                backgroundColor: 'rgba(255, 99, 132, 0.5)',
                tension: 0.1,
                pointRadius: 0,
            },
        ];

        if (toggleSamplePaths.checked && globalSamplePaths.length > 0) {
            const numPathsToShow = parseInt(samplePathsSlider.value, 10);
            for (let i = 0; i < Math.min(numPathsToShow, globalSamplePaths.length); i++) {
                const r = Math.floor(Math.random() * 200) + 50;
                const g = Math.floor(Math.random() * 200) + 50;
                const b = Math.floor(Math.random() * 200) + 50;

                datasets.push({
                    label: `サンプルパス ${i + 1}`,
                    data: globalSamplePaths[i],
                    borderColor: `rgba(${r}, ${g}, ${b}, 0.7)`,
                    backgroundColor: `rgba(${r}, ${g}, ${b}, 0.3)`,
                    borderDash: [5, 5],
                    tension: 0.1,
                    pointRadius: 0,
                    hidden: false,
                });
            }
        }

        simulationChart.data.labels = chartLabels;
        simulationChart.data.datasets = datasets;
        simulationChart.update();
    };

    // Call initChart once on load
    initChart();
    // Load default values on page load
    loadDefaultValues();

    // Event listener for adding life event items
    addLifeEventButton.addEventListener('click', () => addLifeEventItem(lifeEventsContainer)); // lifeEventsContainer を渡す

    // Event listener for removing life event items (delegated)
    lifeEventsContainer.addEventListener('click', (event) => {
        if (event.target.classList.contains('remove-life-event')) {
            event.target.closest('.life-event-item').remove();
        }
    });

    // Event listener for adding investment 1 change setting items
    addInv1ChangeSettingButton.addEventListener('click', () => addChangeSettingItem(inv1ChangeSettingsContainer));

    // Event listener for removing investment 1 change setting items (delegated)
    inv1ChangeSettingsContainer.addEventListener('click', (event) => {
        if (event.target.classList.contains('remove-change-setting')) {
            event.target.closest('.change-setting-item').remove();
        }
    });

    // Event listener for adding investment 2 change setting items
    addInv2ChangeSettingButton.addEventListener('click', () => addChangeSettingItem(inv2ChangeSettingsContainer));

    // Event listener for removing investment 2 change setting items (delegated)
    inv2ChangeSettingsContainer.addEventListener('click', (event) => {
        if (event.target.classList.contains('remove-change-setting')) {
            event.target.closest('.change-setting-item').remove();
        }
    });

    // Event listener for downloading chart
    downloadChartButton.addEventListener('click', () => {
        if (simulationChart) {
            const image = simulationChart.toBase64Image('image/jpeg', 1.0);
            const a = document.createElement('a');
            a.href = image;
            a.download = 'simulation_chart.jpeg';
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    });

    // Event listeners for sample paths control
    toggleSamplePaths.addEventListener('change', updateChartData);
    document.getElementById('current_age').addEventListener('input', updateChartData);
    samplePathsSlider.addEventListener('input', () => {
        samplePathsCountSpan.textContent = samplePathsSlider.value;
        updateChartData();
    });

    simulateButton.addEventListener('click', async () => {
        simulateButton.disabled = true; // Disable button during simulation
        downloadChartButton.style.display = 'none'; // Hide download button until new results are ready

        const investment1Data = getInvestmentData('inv1');
        const investment2Data = getInvestmentData('inv2');
        const existingSavings = parseFloat(document.getElementById('existing_savings').value);
        const lifeEventsData = getLifeEventsData(); // Get life events data
        
        // 暴落設定の取得
        const marketEvent = {
            enabled: document.getElementById('crash_enabled').checked,
            start_year: parseInt(document.getElementById('crash_year').value, 10),
            interval_years: parseInt(document.getElementById('crash_interval').value, 10),
            rate: parseFloat(document.getElementById('crash_rate').value)
        };

        // Basic validation for investment data
        if (isNaN(investment1Data.initial) || isNaN(investment1Data.monthly) || isNaN(investment1Data.return) || isNaN(investment1Data.risk) || isNaN(investment1Data.period) ||
            isNaN(investment2Data.initial) || isNaN(investment2Data.monthly) || isNaN(investment2Data.return) || isNaN(investment2Data.risk) || isNaN(investment2Data.period) ||
            isNaN(existingSavings)) {
            alert('すべての入力フィールドに有効な数値を入力してください。');
            simulateButton.disabled = false;
            return;
        }
        // Basic validation for life events
        for (const event of lifeEventsData) {
            if (isNaN(event.year) || event.year <= 0 || isNaN(event.amount)) {
                alert('ライフイベントの「何年後」は1以上の数値を、「金額」は数値を入力してください。');
                simulateButton.disabled = false;
                return;
            }
        }
        
        // Basic validation for change settings
        const validateChangeSettings = (settings, investmentPrefix) => {
            for (const setting of settings) {
                if (isNaN(setting.year) || setting.year <= 0 || isNaN(setting.monthly)) {
                    alert(`${investmentPrefix}の積立額変更設定で、「何年後」は1以上の数値を、「変更後の積立額」は数値を入力してください。`);
                    return false;
                }
            }
            return true;
        };

        if (!validateChangeSettings(investment1Data.change_settings, '投資信託1') ||
            !validateChangeSettings(investment2Data.change_settings, '投資信託2')) {
            simulateButton.disabled = false;
            return;
        }

        try {
            const response = await fetch('/simulate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    investment1: investment1Data,
                    investment2: investment2Data,
                    existing_savings: existingSavings,
                    life_events: lifeEventsData,
                    market_event: marketEvent,
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const simulationOutput = await response.json();
            console.log("Simulation results:", simulationOutput);

            globalMonthlyResults = simulationOutput.monthly_results; // 結果をグローバル変数に保存
            globalSamplePaths = simulationOutput.sample_paths;   // サンプルパスをグローバル変数に保存

            // Render table results
            resultsTableBody.innerHTML = ''; // Clear previous results
            const now = new Date();
            const startYear = now.getFullYear();
            const startMonth = now.getMonth() + 1;
            const currentAgeValue = parseInt(document.getElementById('current_age').value, 10) || 0;

            globalMonthlyResults.forEach(res => {
                const totalMonths = startMonth - 1 + res.month;
                const y = startYear + Math.floor(totalMonths / 12);
                const m = (totalMonths % 12) + 1;
                const age = currentAgeValue + Math.floor(totalMonths / 12);

                // 最初の月、または1月のデータのみテーブルに表示
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

            // Render summary table
            summaryTableBody.innerHTML = '';
            if (globalMonthlyResults.length > 0) {
                const lastResult = globalMonthlyResults[globalMonthlyResults.length - 1];
                const row = summaryTableBody.insertRow();
                row.insertCell().textContent = formatCurrency(lastResult.min * 10000);
                row.insertCell().textContent = formatCurrency(lastResult.p10 * 10000);
                row.insertCell().textContent = formatCurrency(lastResult.median * 10000);
                row.insertCell().textContent = formatCurrency(lastResult.p90 * 10000);
                row.insertCell().textContent = formatCurrency(lastResult.max * 10000);
                row.insertCell().textContent = formatCurrency(lastResult.average * 10000);
            }
            
            updateChartData(); // グラフ更新関数を呼び出す
            downloadChartButton.style.display = 'block'; // Show download button

        } catch (error) {
            console.error('Error during simulation:', error);
            alert('シミュレーション中にエラーが発生しました: ' + error.message);
        } finally {
            simulateButton.disabled = false; // Re-enable button
        }
    });
});
