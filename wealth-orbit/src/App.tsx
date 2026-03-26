import { useState, useMemo } from 'react';
import { Canvas } from '@react-three/fiber';
import { SpaceScene } from './components/SpaceScene';
import { calculateYearsToTarget } from './utils/finance';
import type { FinanceParams } from './utils/finance';
import './App.css';

function App() {
  const [params, setParams] = useState<FinanceParams>({
    initialInvestment: 1000000,
    monthlySavings: 50000,
    annualRate: 5,
    targetAmount: 30000000,
  });

  const [currentAmount, setCurrentAmount] = useState(params.initialInvestment);
  const [cameraTargetX, setCameraTargetX] = useState(0);

  const yearsToTarget = useMemo(() => calculateYearsToTarget(params), [params]);
  
  const progress = Math.min(Math.max(currentAmount / params.targetAmount, 0), 1);

  const handleParamChange = (key: keyof FinanceParams, value: number) => {
    setParams((prev) => ({ ...prev, [key]: value }));
  };

  return (
    <div className="app-container">
      <div className="sidebar">
        <h1>ウェルス・オービット</h1>
        <p>資産を増やして目標の惑星へ到達しましょう</p>
        
        <div className="input-group">
          <label>初期投資額 (円)</label>
          <input 
            type="number" 
            value={params.initialInvestment} 
            onChange={(e) => handleParamChange('initialInvestment', Number(e.target.value))} 
          />
        </div>

        <div className="input-group">
          <label>毎月の積立額 (円)</label>
          <input 
            type="number" 
            value={params.monthlySavings} 
            onChange={(e) => handleParamChange('monthlySavings', Number(e.target.value))} 
          />
        </div>

        <div className="input-group">
          <label>想定利回り (年利 %)</label>
          <input 
            type="number" 
            value={params.annualRate} 
            onChange={(e) => handleParamChange('annualRate', Number(e.target.value))} 
          />
        </div>

        <div className="input-group">
          <label>目標金額 (円)</label>
          <input 
            type="number" 
            value={params.targetAmount} 
            onChange={(e) => handleParamChange('targetAmount', Number(e.target.value))} 
          />
        </div>

        <hr />

        <div className="input-group">
          <label>現在の資産シミュレーション</label>
          <input 
            type="range" 
            min={0} 
            max={params.targetAmount} 
            value={currentAmount} 
            onChange={(e) => setCurrentAmount(Number(e.target.value))} 
          />
          <div className="value-display">{currentAmount.toLocaleString()} 円</div>
        </div>

        <div className="camera-controls">
          <label>カメラ移動</label>
          <div className="button-row">
            <button onClick={() => setCameraTargetX(0)}>🌍 開始</button>
            <button onClick={() => setCameraTargetX(progress * 100)}>🚀 現在</button>
            <button onClick={() => setCameraTargetX(100)}>🪐 ゴール</button>
          </div>
        </div>

        <div className="results">
          <h3>到達予想: {yearsToTarget === Infinity ? '到達不能' : `${yearsToTarget.toFixed(1)} 年後`}</h3>
          <p>進捗率: {(progress * 100).toFixed(1)} %</p>
          <p className="hint">※右クリック（または二本指ドラッグ）でカメラを自由に移動できます。</p>
        </div>
      </div>

      <div className="canvas-container">
        <Canvas shadows>
          <SpaceScene progress={progress} targetX={cameraTargetX} />
        </Canvas>
      </div>
    </div>
  );
}

export default App;
