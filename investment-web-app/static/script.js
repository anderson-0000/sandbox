// static/script.js
import { formatCurrency } from './utils.js';
import { populateForm, addLifeEventItem, getInvestmentData, getLifeEventsData } from './dom_handlers.js';

document.addEventListener('DOMContentLoaded', () => {
    const simulateButton = document.getElementById('simulateButton');
    const downloadChartButton = document.getElementById('downloadChartButton');
    const resultsTableBody = document.querySelector('#resultsTable tbody');
    const summaryTableBody = document.querySelector('#summaryTable tbody');
    const chartCanvas = document.getElementById('simulationChart');

    // Life Event elements
    const addLifeEventButton = document.getElementById('addLifeEventButton');
    const lifeEventsContainer = document.getElementById('life_events_container');

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

            // Populate life events
            lifeEventsContainer.innerHTML = ''; // Clear existing
            if (defaults.life_events && defaults.life_events.length > 0) {
                defaults.life_events.forEach(event => {
                    addLifeEventItem(lifeEventsContainer, event.year, event.amount); // lifeEventsContainer を渡す
                });
            } else {
                addLifeEventItem(lifeEventsContainer); // Add one empty by default, lifeEventsContainer を渡す
            }

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
                labels: [], // Will be filled with years
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
                        text: 'ポートフォリオ価値の年次推移',
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
                            text: 'ポートフォリオ価値 (円)',
                        },
                        ticks: {
                            callback: function(value) {
                                return formatCurrency(value);
                            }
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: '年',
                        },
                    },
                },
            },
        });
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

    simulateButton.addEventListener('click', async () => {
        simulateButton.disabled = true; // Disable button during simulation
        downloadChartButton.style.display = 'none'; // Hide download button until new results are ready

        const investment1Data = getInvestmentData('inv1');
        const investment2Data = getInvestmentData('inv2');
        const existingSavings = parseFloat(document.getElementById('existing_savings').value);
        const lifeEventsData = getLifeEventsData(); // Get life events data

        // Basic validation for investment data
        if (Object.values(investment1Data).some(isNaN) || Object.values(investment2Data).some(isNaN) || isNaN(existingSavings)) {
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
                    life_events: lifeEventsData, // Add life events
                }),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const results = await response.json();
            console.log("Simulation results:", results);

            // Render table results
            resultsTableBody.innerHTML = ''; // Clear previous results
            results.forEach(res => {
                const row = resultsTableBody.insertRow();
                row.insertCell().textContent = res.year;
                row.insertCell().textContent = formatCurrency(res.min);
                row.insertCell().textContent = formatCurrency(res.p10);
                row.insertCell().textContent = formatCurrency(res.median);
                row.insertCell().textContent = formatCurrency(res.p90);
                row.insertCell().textContent = formatCurrency(res.max);
                row.insertCell().textContent = formatCurrency(res.average);
            });

            // Render summary table
            summaryTableBody.innerHTML = '';
            if (results.length > 0) {
                const lastResult = results[results.length - 1];
                const row = summaryTableBody.insertRow();
                row.insertCell().textContent = formatCurrency(lastResult.min); // Final min
                row.insertCell().textContent = formatCurrency(lastResult.p10); // Final 10%
                row.insertCell().textContent = formatCurrency(lastResult.median); // Final median
                row.insertCell().textContent = formatCurrency(lastResult.p90); // Final 90%
                row.insertCell().textContent = formatCurrency(lastResult.max); // Final max
                row.insertCell().textContent = formatCurrency(lastResult.average); // Final average
            }
            
            // Update Chart.js data
            const chartLabels = results.map(res => res.year);
            simulationChart.data.labels = chartLabels;
            simulationChart.data.datasets = [
                {
                    label: '中央値 (50%)',
                    data: results.map(res => res.median),
                    borderColor: 'rgb(75, 192, 192)',
                    backgroundColor: 'rgba(75, 192, 192, 0.5)',
                    tension: 0.1,
                },
                {
                    label: '90% パーセンタイル',
                    data: results.map(res => res.p90),
                    borderColor: 'rgb(53, 162, 235)',
                    backgroundColor: 'rgba(53, 162, 235, 0.5)',
                    tension: 0.1,
                },
                {
                    label: '10% パーセンタイル',
                    data: results.map(res => res.p10),
                    borderColor: 'rgb(255, 99, 132)',
                    backgroundColor: 'rgba(255, 99, 132, 0.5)',
                    tension: 0.1,
                },
            ];
            simulationChart.update(); // Update chart to re-render with new data
            downloadChartButton.style.display = 'block'; // Show download button

        } catch (error) {
            console.error('Error during simulation:', error);
            alert('シミュレーション中にエラーが発生しました: ' + error.message);
        } finally {
            simulateButton.disabled = false; // Re-enable button
        }
    });
});