document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('bp-form');
    const dateInput = document.getElementById('date');
    const timeInput = document.getElementById('time');
    const chartCanvas = document.getElementById('bp-chart');
    const dataTableContainer = document.getElementById('data-table-container');

    let chart;
    let bloodPressureData = [];
    let editingRowId = null; // インライン編集中の行のID

    // 日付と時刻の初期値を今日/現在時刻に設定
    const now = new Date();
    const yyyy = now.getFullYear();
    const mm = String(now.getMonth() + 1).padStart(2, '0');
    const dd = String(now.getDate()).padStart(2, '0');
    const hh = String(now.getHours()).padStart(2, '0');
    const min = String(now.getMinutes()).padStart(2, '0');
    dateInput.value = `${yyyy}-${mm}-${dd}`;
    timeInput.value = `${hh}:${min}`;

    /**
     * データテーブルを描画または更新する関数
     * @param {Array} data - 表示するデータの配列
     */
    const renderTable = (data) => {
        // データが空の場合はメッセージを表示
        if (data.length === 0) {
            dataTableContainer.innerHTML = '<p>記録されたデータはありません。</p>';
            return;
        }

        // 新しいデータが上に来るように逆順のコピーを作成
        const reversedData = [...data].reverse();

        // テーブルのHTMLを生成
        const table = `
            <table>
                <thead>
                    <tr>
                        <th>日付</th>
                        <th>時刻</th>
                        <th>最高血圧 (mmHg)</th>
                        <th>最低血圧 (mmHg)</th>
                        <th>心拍数 (bpm)</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    ${reversedData.map(d => {
                        const dt = new Date(d.date);
                        const dateStr = `${dt.getFullYear()}-${String(dt.getMonth() + 1).padStart(2, '0')}-${String(dt.getDate()).padStart(2, '0')}`;
                        const timeStr = `${String(dt.getHours()).padStart(2, '0')}:${String(dt.getMinutes()).padStart(2, '0')}`;

                        if (editingRowId === d.date) {
                            // 編集中行の表示 (インライン編集モード)
                            return `
                                <tr data-id="${d.date}">
                                    <td><input type="date" class="edit-date" value="${dateStr}"></td>
                                    <td><input type="time" class="edit-time" value="${timeStr}"></td>
                                    <td><input type="number" class="edit-systolic" value="${d.systolic}"></td>
                                    <td><input type="number" class="edit-diastolic" value="${d.diastolic}"></td>
                                    <td><input type="number" class="edit-pulse" value="${d.pulse}"></td>
                                    <td>
                                        <button class="save-btn" data-id="${d.date}">保存</button>
                                        <button class="cancel-btn" data-id="${d.date}">キャンセル</button>
                                    </td>
                                </tr>
                            `;
                        } else {
                            // 通常行の表示
                            return `
                                <tr data-id="${d.date}">
                                    <td>${dateStr}</td>
                                    <td>${timeStr}</td>
                                    <td>${d.systolic}</td>
                                    <td>${d.diastolic}</td>
                                    <td>${d.pulse}</td>
                                    <td>
                                        <button class="edit-btn" data-id="${d.date}">編集</button>
                                        <button class="delete-btn" data-id="${d.date}">削除</button>
                                    </td>
                                </tr>
                            `;
                        }
                    }).join('')}
                </tbody>
            </table>
        `;
        dataTableContainer.innerHTML = table;
    };

    /**
     * グラフを描画または更新する関数
     * @param {Array} data - 表示するデータの配列
     */
            const renderChart = (data) => {
                const labels = data.map(d => {
                    const dt = new Date(d.date);
                    const month = String(dt.getMonth() + 1).padStart(2, '0');
                    const day = String(dt.getDate()).padStart(2, '0');
                    const hours = String(dt.getHours()).padStart(2, '0');
                    const minutes = String(dt.getMinutes()).padStart(2, '0');
                    return `${month}/${day} ${hours}:${minutes}`;
                });
    
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
            renderTable(bloodPressureData);
        } catch (error) {
            console.error('データの読み込みに失敗しました:', error);
            // エラーが発生しても空のグラフとテーブルを描画
            renderChart([]);
            renderTable([]);
        }
    };

    /**
     * フォーム送信時の処理
     */
    form.addEventListener('submit', (e) => {
        e.preventDefault();

        const newRecord = {
            date: `${dateInput.value}T${timeInput.value}:00`, // ISO 8601形式で結合
            systolic: parseInt(document.getElementById('systolic').value, 10),
            diastolic: parseInt(document.getElementById('diastolic').value, 10),
            pulse: parseInt(document.getElementById('pulse').value, 10),
        };

        // 新規データを追加
        bloodPressureData.push(newRecord);
        
        bloodPressureData.sort((a, b) => new Date(a.date) - new Date(b.date));
        
        // 更新されたデータをJSONファイルとしてダウンロード
        triggerDownload(bloodPressureData);
        
        // グラフとテーブルを再描画
        renderChart(bloodPressureData);
        renderTable(bloodPressureData);
        
        form.reset();
        // 日付と時刻を再度今日/現在時刻に設定
        const now = new Date();
        const yyyy = now.getFullYear();
        const mm = String(now.getMonth() + 1).padStart(2, '0');
        const dd = String(now.getDate()).padStart(2, '0');
        const hh = String(now.getHours()).padStart(2, '0');
        const min = String(now.getMinutes()).padStart(2, '0');
        dateInput.value = `${yyyy}-${mm}-${dd}`;
        timeInput.value = `${hh}:${min}`;
    });

    /**
     * データ配列をJSONファイルとしてダウンロードさせる
     * @param {Array} data ダウンロードさせるデータ
     */
    const triggerDownload = (data) => {
        const jsonString = JSON.stringify(data, null, 2);
        const blob = new Blob([jsonString], { type: 'application/json' });
        const url = URL.createObjectURL(blob);

        const a = document.createElement('a');
        a.href = url;
        a.download = 'data.json';
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
    };

    /**
     * テーブル内のボタンクリックを処理する
     */
    dataTableContainer.addEventListener('click', (e) => {
        const target = e.target;
        const id = target.dataset.id;

        if (target.classList.contains('delete-btn')) {
            if (confirm('このデータを削除してもよろしいですか？')) {
                // データ配列から該当IDのものを削除
                bloodPressureData = bloodPressureData.filter(d => d.date !== id);
                // 表示を更新
                renderChart(bloodPressureData);
                renderTable(bloodPressureData);
                // 更新後のデータをダウンロード
                triggerDownload(bloodPressureData);
            }
        } else if (target.classList.contains('edit-btn')) {
            // インライン編集モードに入る
            editingRowId = id;
            renderTable(bloodPressureData); // UIを更新して編集モードにする
        } else if (target.classList.contains('save-btn')) {
            const row = target.closest('tr');
            const updatedDate = row.querySelector('.edit-date').value;
            const updatedTime = row.querySelector('.edit-time').value;
            const updatedSystolic = parseInt(row.querySelector('.edit-systolic').value, 10);
            const updatedDiastolic = parseInt(row.querySelector('.edit-diastolic').value, 10);
            const updatedPulse = parseInt(row.querySelector('.edit-pulse').value, 10);

            const updatedRecord = {
                date: `${updatedDate}T${updatedTime}:00`,
                systolic: updatedSystolic,
                diastolic: updatedDiastolic,
                pulse: updatedPulse,
            };

            // データ配列内で該当IDのものを更新
            bloodPressureData = bloodPressureData.map(d => 
                d.date === id ? updatedRecord : d
            );
            
            // 編集モードを終了
            editingRowId = null;
            
            // 表示を更新
            bloodPressureData.sort((a, b) => new Date(a.date) - new Date(b.date)); // ソートし直す
            renderChart(bloodPressureData);
            renderTable(bloodPressureData);
            // 更新後のデータをダウンロード
            triggerDownload(bloodPressureData);
        } else if (target.classList.contains('cancel-btn')) {
            // 編集モードを終了
            editingRowId = null;
            renderTable(bloodPressureData); // UIを更新して通常モードに戻す
        }
    });

    // 初期データの読み込み
    loadData();

    // グラフダウンロード処理
    const downloadBtn = document.getElementById('download-chart-btn');
    downloadBtn.addEventListener('click', () => {
        if (!chart) return; // グラフがなければ何もしない

        const originalCanvas = chart.canvas;
        
        // 新しいcanvasを作成して背景を白にする
        const newCanvas = document.createElement('canvas');
        newCanvas.width = originalCanvas.width;
        newCanvas.height = originalCanvas.height;
        const newCtx = newCanvas.getContext('2d');
        
        // 背景を白で塗りつぶす
        newCtx.fillStyle = '#ffffff';
        newCtx.fillRect(0, 0, newCanvas.width, newCanvas.height);
        
        // 既存のグラフを新しいcanvasに描画
        newCtx.drawImage(originalCanvas, 0, 0);

        // ダウンロードリンクを作成
        const a = document.createElement('a');
        a.href = newCanvas.toDataURL('image/jpeg', 1.0);
        a.download = 'blood-pressure-chart.jpeg';
        a.click();
    });
});
