import React from 'react';
import { Rocket, RefreshCw, AlertCircle } from 'lucide-react';

interface DashboardProps {
  velocity: number;
  angle: number;
  onVelocityChange: (v: number) => void;
  onAngleChange: (a: number) => void;
  onLaunch: () => void;
  onReset: () => void;
  isLaunched: boolean;
  status: 'idle' | 'orbiting' | 'crashed' | 'escaped';
  stats: {
    speed: number;
    altitude: number;
  };
}

export function Dashboard({
  velocity,
  angle,
  onVelocityChange,
  onAngleChange,
  onLaunch,
  onReset,
  isLaunched,
  status,
  stats
}: DashboardProps) {
  const statusLabels = {
    idle: '待機中',
    orbiting: '軌道周回中',
    crashed: '墜落',
    escaped: '重力圏脱出'
  };

  return (
    <div className="dashboard-container">
      <div className="title">
        <h1>Can I Orbit?</h1>
        <p>衛星投入シミュレーター</p>
      </div>

      <div className="controls">
        <div className="input-group">
          <label>発射速度 (km/s)</label>
          <input
            type="range"
            min="1"
            max="15"
            step="0.1"
            value={velocity}
            onChange={(e) => onVelocityChange(parseFloat(e.target.value))}
            disabled={isLaunched}
          />
          <div className="value-display">{velocity.toFixed(1)} km/s</div>
        </div>

        <div className="input-group">
          <label>発射角度 (度)</label>
          <input
            type="range"
            min="0"
            max="90"
            step="1"
            value={angle}
            onChange={(e) => onAngleChange(parseFloat(e.target.value))}
            disabled={isLaunched}
          />
          <div className="value-display">{angle}°</div>
        </div>

        {!isLaunched ? (
          <button className="launch-btn" onClick={onLaunch}>
            <Rocket size={18} /> 発射
          </button>
        ) : (
          <button className="reset-btn" onClick={onReset}>
            <RefreshCw size={18} /> リセット
          </button>
        )}
      </div>

      <div className="status-panel">
        <div className="status-item">
          <span className="label">ステータス:</span>
          <span className={`status-badge ${status}`}>{statusLabels[status]}</span>
        </div>
        <div className="status-item">
          <span className="label">現在の速度:</span>
          <span className="value">{(stats.speed / 1000).toFixed(2)} km/s</span>
        </div>
        <div className="status-item">
          <span className="label">高度:</span>
          <span className="value">{(stats.altitude / 1000).toFixed(0)} km</span>
        </div>
      </div>

      {status === 'crashed' && (
        <div className="alert-message error">
          <AlertCircle size={16} /> 衛星が地球に墜落しました！
        </div>
      )}
      {status === 'escaped' && (
        <div className="alert-message warning">
          <AlertCircle size={16} /> 衛星が地球の重力圏を脱出しました！
        </div>
      )}

      <style>{`
        .dashboard-container {
          position: absolute;
          top: 20px;
          left: 20px;
          width: 300px;
          background: rgba(0, 0, 0, 0.8);
          color: white;
          padding: 20px;
          border-radius: 12px;
          font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
          border: 1px solid rgba(255, 255, 255, 0.1);
          backdrop-filter: blur(8px);
          z-index: 1000;
        }
        .title h1 { margin: 0; font-size: 24px; color: #00ffff; }
        .title p { margin: 5px 0 20px 0; font-size: 14px; color: #888; }
        .controls { margin-bottom: 20px; }
        .input-group { margin-bottom: 15px; }
        .input-group label { display: block; font-size: 12px; color: #aaa; margin-bottom: 5px; }
        input[type="range"] { width: 100%; cursor: pointer; }
        .value-display { font-weight: bold; font-size: 16px; margin-top: 5px; }
        button {
          width: 100%;
          padding: 10px;
          border: none;
          border-radius: 6px;
          font-weight: bold;
          cursor: pointer;
          display: flex;
          align-items: center;
          justify-content: center;
          gap: 8px;
          transition: 0.2s;
        }
        .launch-btn { background: #00ffff; color: #000; }
        .launch-btn:hover { background: #00cccc; }
        .reset-btn { background: #333; color: #fff; }
        .reset-btn:hover { background: #444; }
        .status-panel { background: rgba(255, 255, 255, 0.05); padding: 10px; border-radius: 6px; }
        .status-item { display: flex; justify-content: space-between; margin-bottom: 5px; font-size: 14px; }
        .status-badge { font-weight: bold; border-radius: 4px; padding: 2px 6px; font-size: 10px; }
        .status-badge.idle { background: #555; }
        .status-badge.orbiting { background: #22c55e; }
        .status-badge.crashed { background: #ef4444; }
        .status-badge.escaped { background: #f59e0b; }
        .alert-message { margin-top: 15px; padding: 10px; border-radius: 6px; font-size: 12px; display: flex; align-items: center; gap: 8px; }
        .alert-message.error { background: rgba(239, 68, 68, 0.2); color: #ef4444; border: 1px solid #ef4444; }
        .alert-message.warning { background: rgba(245, 158, 11, 0.2); color: #f59e0b; border: 1px solid #f59e0b; }
      `}</style>
    </div>
  );
}
