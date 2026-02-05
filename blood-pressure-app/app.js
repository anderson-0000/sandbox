document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('bp-form');
    const dateInput = document.getElementById('date');
    const chartCanvas = document.getElementById('bp-chart');
    
    let chart;
    let bloodPressureData = [];

    // 日付の初期値を今日に設定
    const today = new Date();
    const yyyy = today.getFullYear();
    const mm = String(today.getMonth() + 1).padStart(2, '0');
    const dd = String(today.getDate()).padStart(2, '0');
    dateInput.value = `${yyyy}-${mm}-${dd}`;

    /**
     * グラフを描画または更新する関数
     * @param {Array} data - 表示するデータの配列
     */
    const renderChart = (data) => {
        const labels = data.map(d => d.date);
        const systolicData = data.map(d => d.systolic);
        const diastolicData = data.map(d => d.diastolic);
        const pulseData = data.map(d => d.pulse);

        const chartConfig = {
            type: 'line',
            data: {
                labels: labels,
                datasets: [
                    {
                        label: '最高血圧 (mmHg)',
                        data: systolicData,
                        borderColor: 'rgba(255, 99, 132, 1)',
                        backgroundColor: 'rgba(255, 99, 132, 0.2)',
                        fill: false,
                        yAxisID: 'y-axis-pressure',
                    },
                    {
                        label: '最低血圧 (mmHg)',
                        data: diastolicData,
                        borderColor: 'rgba(54, 162, 235, 1)',
                        backgroundColor: 'rgba(54, 162, 235, 0.2)',
                        fill: false,
                        yAxisID: 'y-axis-pressure',
                    },
                    {
                        label: '心拍数 (bpm)',
                        data: pulseData,
                        borderColor: 'rgba(75, 192, 192, 1)',
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        fill: false,
                        yAxisID: 'y-axis-pulse',
                    }
                ]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    'y-axis-pressure': {
                        type: 'linear',
                        position: 'left',
                        title: {
                            display: true,
                            text: '血圧 (mmHg)'
                        },
                        suggestedMin: 50,
                        suggestedMax: 150
                    },
                    'y-axis-pulse': {
                        type: 'linear',
                        position: 'right',
                        title: {
                            display: true,
                            text: '心拍数 (bpm)'
                        },
                        suggestedMin: 40,
                        suggestedMax: 120,
                        grid: {
                            drawOnChartArea: false, // グラフエリアのグリッド線を非表示に
                        },
                    }
                }
            }
        };

        if (chart) {
            chart.destroy();
        }
        chart = new Chart(chartCanvas, chartConfig);
    };
    
    /**
     * data.jsonからデータを読み込む関数
     */
    const loadData = async () => {
        try {
            // キャッシュを無効化して常に最新のファイルを取得
            const response = await fetch(`data.json?_=${new Date().getTime()}`);
            const data = await response.json();
            // 日付でソート
            bloodPressureData = data.sort((a, b) => new Date(a.date) - new Date(b.date));
            renderChart(bloodPressureData);
        } catch (error) {
            console.error('データの読み込みに失敗しました:', error);
            // エラーが発生しても空のグラフを描画
            renderChart([]);
        }
    };

    /**
     * フォーム送信時の処理
     */
    form.addEventListener('submit', (e) => {
        e.preventDefault();

        const newRecord = {
            date: dateInput.value,
            systolic: parseInt(document.getElementById('systolic').value, 10),
            diastolic: parseInt(document.getElementById('diastolic').value, 10),
            pulse: parseInt(document.getElementById('pulse').value, 10),
        };

        // データを追加してソート
        bloodPressureData.push(newRecord);
        bloodPressureData.sort((a, b) => new Date(a.date) - new Date(b.date));
        
        // 更新されたデータをJSONファイルとしてダウンロード
        const jsonString = JSON.stringify(bloodPressureData, null, 2);
        const blob = new Blob([jsonString], { type: 'application/json' });
        const url = URL.createObjectURL(blob);

        const a = document.createElement('a');
        a.href = url;
        a.download = 'data.json';
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
        
        // グラフを再描画
        renderChart(bloodPressureData);
        
        form.reset();
        // 日付を再度今日の日付に設定
        dateInput.value = `${yyyy}-${mm}-${dd}`;
    });

    // 初期データの読み込み
    loadData();
});
