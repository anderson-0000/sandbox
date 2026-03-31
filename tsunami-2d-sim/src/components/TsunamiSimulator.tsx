import React, { useState, useMemo, useEffect, useRef } from 'react';
import { useTsunamiPhysics } from '../hooks/useTsunamiPhysics';
import './TsunamiSimulator.css';

const TsunamiSimulator: React.FC = () => {
  // Scenario settings
  const [magnitude, setMagnitude] = useState(8.5);
  const [sourceDepth, setSourceDepth] = useState(4000);
  const [sourceDistance, setSourceDistance] = useState(200); // km
  const [shelfWidth, setShelfWidth] = useState(50); // km
  const [shelfDepth, setShelfDepth] = useState(200); // m
  
  // Animation state
  const [isPlaying, setIsPlaying] = useState(false);
  const [elapsedTime, setElapsedTime] = useState(0); 
  const [playbackSpeed, setPlaybackSpeed] = useState(10); 
  
  const { calculateInitialHeight, calculateHeightAtDepth, calculateArrivalTimeNumerical } = useTsunamiPhysics();
  const requestRef = useRef<number | null>(null);
  const lastTimeRef = useRef<number | null>(null);

  // Depth Profile: Source -> Shelf Edge -> Shore
  const getDepthAtX = (x: number) => {
    if (x >= shelfWidth) {
      const slope = (sourceDepth - shelfDepth) / (sourceDistance - shelfWidth);
      return shelfDepth + (x - shelfWidth) * slope;
    } else if (x > 0) {
      const slope = shelfDepth / shelfWidth;
      return x * slope;
    }
    return 0; // Shore
  };

  const animate = (time: number) => {
    if (lastTimeRef.current !== null) {
      const deltaTime = (time - lastTimeRef.current) / 1000; 
      setElapsedTime(prev => prev + deltaTime * playbackSpeed);
    }
    lastTimeRef.current = time;
    requestRef.current = requestAnimationFrame(animate);
  };

  useEffect(() => {
    if (isPlaying) {
      requestRef.current = requestAnimationFrame(animate);
    } else {
      if (requestRef.current !== null) cancelAnimationFrame(requestRef.current);
      lastTimeRef.current = null;
    }
    return () => {
      if (requestRef.current !== null) cancelAnimationFrame(requestRef.current);
    };
  }, [isPlaying, playbackSpeed]);

  const initialHeight = useMemo(() => calculateInitialHeight(magnitude), [magnitude, calculateInitialHeight]);

  const arrivalTimes = useMemo(() => {
    const points = 200;
    const times = [];
    for (let i = 0; i <= points; i++) {
      const x = (i / points) * sourceDistance;
      times.push(calculateArrivalTimeNumerical(x, sourceDistance, getDepthAtX));
    }
    return times;
  }, [sourceDistance, shelfWidth, shelfDepth, sourceDepth]);

  const plotData = useMemo(() => {
    const points = 200;
    const data = [];
    const baseWaveWidth = sourceDistance * 0.1; 

    for (let i = 0; i <= points; i++) {
      const x = (i / points) * sourceDistance;
      const currentDepth = getDepthAtX(x);
      const peakArrivalTime = arrivalTimes[i];
      const timeDiff = elapsedTime - peakArrivalTime;
      
      const speedKmS = Math.sqrt(9.81 * Math.max(currentDepth, 1)) / 1000;
      const localWidth = baseWaveWidth * (speedKmS / (Math.sqrt(9.81 * sourceDepth) / 1000));
      const distanceDiff = timeDiff * speedKmS;
      const gaussian = Math.exp(-Math.pow(distanceDiff / Math.max(localWidth, 0.1), 2));
      const localMaxHeight = calculateHeightAtDepth(initialHeight, sourceDepth, currentDepth);
      const waveHeight = localMaxHeight * gaussian;
      
      data.push({
        x,
        depth: -currentDepth,
        waveHeight: waveHeight,
      });
    }
    return data;
  }, [magnitude, sourceDepth, sourceDistance, shelfWidth, shelfDepth, elapsedTime, initialHeight, arrivalTimes]);

  // SVG configuration
  const width = 800;
  const height = 500;
  const margin = { top: 60, right: 40, bottom: 60, left: 60 };
  const innerWidth = width - margin.left - margin.right;
  const innerHeight = height - margin.top - margin.bottom;

  const displayXRange = sourceDistance + 20; // 20km land
  const xScale = (km: number) => ((km + 20) / displayXRange) * innerWidth;
  
  const yScale = (m: number) => {
    const totalMeters = sourceDepth * 1.5;
    return innerHeight - ((m + sourceDepth) / totalMeters) * innerHeight;
  };

  // Paths
  const landX = -20;
  const landHeight = 50; 
  const bottomY = innerHeight;

  // 1. Seabed & Land Path (Ground)
  // Close the path at the BOTTOM of the SVG (innerHeight)
  const seabedPath = `M ${margin.left + xScale(landX)} ${margin.top + yScale(landHeight)} ` +
    `L ${margin.left + xScale(0)} ${margin.top + yScale(0)} ` +
    plotData.map(p => `L ${margin.left + xScale(p.x)} ${margin.top + yScale(p.depth)}`).join(' ') +
    ` L ${margin.left + xScale(sourceDistance)} ${margin.top + yScale(-sourceDepth)}` +
    ` L ${margin.left + xScale(sourceDistance)} ${margin.top + bottomY}` +
    ` L ${margin.left + xScale(landX)} ${margin.top + bottomY} Z`;

  // 2. Still Water Path (Area between seabed and 0m)
  const seaStillPath = `M ${margin.left + xScale(0)} ${margin.top + yScale(0)} ` +
    plotData.map(p => `L ${margin.left + xScale(p.x)} ${margin.top + yScale(0)}`).join(' ') +
    ` L ${margin.left + xScale(sourceDistance)} ${margin.top + yScale(0)} ` +
    ` L ${margin.left + xScale(sourceDistance)} ${margin.top + yScale(-sourceDepth)} ` +
    plotData.slice().reverse().map(p => `L ${margin.left + xScale(p.x)} ${margin.top + yScale(p.depth)}`).join(' ') +
    ` Z`;

  // 3. Tsunami Path (Area between 0m and wave profile)
  const tsunamiPath = `M ${margin.left + xScale(0)} ${margin.top + yScale(0)} ` +
    plotData.map(p => `L ${margin.left + xScale(p.x)} ${margin.top + yScale(p.waveHeight)}`).join(' ') +
    ` L ${margin.left + xScale(sourceDistance)} ${margin.top + yScale(0)} Z`;

  const minutes = Math.floor(elapsedTime / 60);
  const seconds = Math.floor(elapsedTime % 60);

  return (
    <div className="simulator-container">
      <h2>2D 津波断面シミュレーター (視覚的バグ修正版)</h2>
      
      <div className="controls-section">
        <div className="control-column">
          <h3>震源設定</h3>
          <div className="control-group">
            <label>マグニチュード (Mw): {magnitude.toFixed(1)}</label>
            <input type="range" min="5" max="9.5" step="0.1" value={magnitude} onChange={(e) => setMagnitude(parseFloat(e.target.value))} />
          </div>
          <div className="control-group">
            <label>震源水深 (m): {sourceDepth}</label>
            <input type="number" step="100" value={sourceDepth} onChange={(e) => setSourceDepth(parseInt(e.target.value))} />
          </div>
          <div className="control-group">
            <label>震源距離 (km): {sourceDistance}</label>
            <input type="number" step="10" value={sourceDistance} onChange={(e) => setSourceDistance(parseInt(e.target.value))} />
          </div>
        </div>

        <div className="control-column">
          <h3>海底地形設定</h3>
          <div className="control-group">
            <label>大陸棚の幅 (km): {shelfWidth}</label>
            <input type="range" min="0" max={sourceDistance} step="5" value={shelfWidth} onChange={(e) => setShelfWidth(parseInt(e.target.value))} />
          </div>
          <div className="control-group">
            <label>大陸棚の水深 (m): {shelfDepth}</label>
            <input type="range" min="10" max={1000} step="10" value={shelfDepth} onChange={(e) => setShelfDepth(parseInt(e.target.value))} />
          </div>
        </div>
      </div>

      <div className="animation-controls">
        <button onClick={() => setIsPlaying(!isPlaying)} className={isPlaying ? 'pause' : 'play'}>
          {isPlaying ? '一時停止' : '再生'}
        </button>
        <button onClick={() => { setElapsedTime(0); setIsPlaying(false); }}>
          リセット
        </button>
        <div className="control-group">
          <label>再生速度: x{playbackSpeed}</label>
          <input type="range" min="1" max="100" step="1" value={playbackSpeed} onChange={(e) => setPlaybackSpeed(parseInt(e.target.value))} />
        </div>
        <div className="timer">
          経過時間: {minutes}分{seconds}秒
        </div>
      </div>

      <div className="plot-container">
        <svg width={width} height={height} viewBox={`0 0 ${width} ${height}`}>
          <defs>
            <linearGradient id="skyGradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stopColor="#bae6fd" />
              <stop offset="100%" stopColor="#fff" />
            </linearGradient>
            <linearGradient id="groundGradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stopColor="#475569" />
              <stop offset="100%" stopColor="#1e293b" />
            </linearGradient>
          </defs>

          {/* Background / Sky */}
          <rect x={margin.left} y={margin.top} width={innerWidth} height={innerHeight} fill="url(#skyGradient)" />

          {/* 1. Ground (Seabed & Land) */}
          <path d={seabedPath} fill="url(#groundGradient)" stroke="#1e2937" strokeWidth="2" />

          {/* 2. Still Water (Quiet Sea) */}
          <path d={seaStillPath} fill="rgba(37, 99, 235, 0.4)" />

          {/* 3. Tsunami (Wave) */}
          <path d={tsunamiPath} fill="#0ea5e9" stroke="#fff" strokeWidth="2" strokeLinejoin="round" />

          {/* 4. Land Surface Highlight */}
          <path d={`M ${margin.left + xScale(landX)} ${margin.top + yScale(landHeight)} L ${margin.left + xScale(0)} ${margin.top + yScale(0)}`} stroke="#15803d" strokeWidth="4" />
          <path d={`M ${margin.left + xScale(landX)} ${margin.top + yScale(landHeight)} L ${margin.left + xScale(landX)} ${margin.top + yScale(0)} L ${margin.left + xScale(0)} ${margin.top + yScale(0)} Z`} fill="#15803d" opacity="0.4" />

          {/* Labels */}
          <text x={margin.left + xScale(landX / 2)} y={margin.top + yScale(landHeight + 10)} textAnchor="middle" fill="#15803d" fontWeight="bold">陸地</text>
          <text x={margin.left + xScale(sourceDistance / 2)} y={margin.top + yScale(100)} textAnchor="middle" fill="#1d4ed8" opacity="0.6">静かな海</text>
          
          {/* Axes & Grids */}
          <line x1={margin.left} y1={margin.top + yScale(0)} x2={margin.left + innerWidth} y2={margin.top + yScale(0)} stroke="#94a3b8" strokeDasharray="4" />
          <text x={margin.left} y={margin.top + yScale(0)} textAnchor="end" dx="-5" fontSize="12">海面 0m</text>
          <text x={margin.left} y={margin.top + yScale(-sourceDepth)} textAnchor="end" dx="-5" fontSize="12">-{sourceDepth}m</text>
          
          <text x={margin.left + xScale(0)} y={height - 20} textAnchor="middle" fontSize="12">海岸</text>
          <text x={margin.left + xScale(sourceDistance)} y={height - 20} textAnchor="middle" fontSize="12">震源 {sourceDistance}km</text>
        </svg>
      </div>

      <div className="explanation">
        <p>※グレーのグラデーション部分は「海底」および「地下」を表しています。0mより上の青い部分は「海（波）」です。</p>
        <p>※陸地（緑）は海岸線より左側にあり、津波が到達すると波が乗り上げる様子を確認できます。</p>
      </div>
    </div>
  );
};

export default TsunamiSimulator;
