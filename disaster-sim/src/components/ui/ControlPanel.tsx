import React from 'react';
import { useStore } from '../../store/useStore';
import { Play, Pause, RotateCcw, Activity, Waves } from 'lucide-react';

export const ControlPanel: React.FC = () => {
  const { 
    isActive, currentTime, epicenter, toggleActive, reset,
    magnitude, setMagnitude, playbackSpeed, setPlaybackSpeed 
  } = useStore();

  const formatTime = (seconds: number) => {
    const mins = Math.floor(seconds / 60);
    const secs = Math.floor(seconds % 60);
    return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
  };

  return (
    <div className="absolute top-6 left-6 w-80 space-y-4 z-10">
      <div className="bg-[#0f172a]/80 backdrop-blur-xl border border-white/10 rounded-2xl p-6 shadow-2xl">
        <div className="flex items-center justify-between mb-6">
          <div className="flex items-center gap-2">
            <div className="w-2 h-2 bg-red-500 rounded-full animate-pulse shadow-[0_0_8px_rgba(239,68,68,0.8)]" />
            <h1 className="text-white font-black tracking-widest text-xs uppercase">Disaster Monitor</h1>
          </div>
          <div className="text-[#94a3b8] font-mono text-[10px] tabular-nums">
            {new Date().toLocaleTimeString()}
          </div>
        </div>

        <div className="space-y-6">
          {/* Timer Display */}
          <div className="bg-black/40 rounded-xl p-4 border border-white/5">
            <div className="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-1 flex items-center gap-1.5">
              <Activity size={10} /> Elapsed Time
            </div>
            <div className="text-3xl font-black font-mono text-cyan-400 tabular-nums">
              {formatTime(currentTime)}
            </div>
          </div>

          {/* Controls */}
          <div className="flex gap-2">
            <button
              onClick={toggleActive}
              className={`flex-1 flex items-center justify-center gap-2 py-3 rounded-xl font-bold transition-all ${
                isActive 
                  ? 'bg-amber-500/20 text-amber-500 border border-amber-500/20' 
                  : 'bg-indigo-600 text-white shadow-lg shadow-indigo-500/20'
              }`}
            >
              {isActive ? <Pause size={18} /> : <Play size={18} />}
              {isActive ? 'Pause' : 'Start'}
            </button>
            <button
              onClick={reset}
              className="px-4 bg-slate-800 text-slate-400 rounded-xl hover:bg-slate-700 transition-colors"
            >
              <RotateCcw size={18} />
            </button>
          </div>

          {/* Settings */}
          <div className="space-y-4 pt-2">
            <div className="space-y-2">
              <div className="flex justify-between text-[10px] font-bold uppercase tracking-widest text-slate-400">
                <span>Magnitude</span>
                <span className="text-cyan-400">M{magnitude.toFixed(1)}</span>
              </div>
              <input
                type="range"
                min="5"
                max="9"
                step="0.1"
                value={magnitude}
                onChange={(e) => setMagnitude(Number(e.target.value))}
                className="w-full h-1.5 bg-slate-700 rounded-lg appearance-none cursor-pointer accent-cyan-500"
              />
            </div>

            <div className="space-y-2">
              <div className="flex justify-between text-[10px] font-bold uppercase tracking-widest text-slate-400">
                <span>Playback Speed</span>
                <span className="text-indigo-400">{playbackSpeed}x</span>
              </div>
              <input
                type="range"
                min="1"
                max="100"
                step="1"
                value={playbackSpeed}
                onChange={(e) => setPlaybackSpeed(Number(e.target.value))}
                className="w-full h-1.5 bg-slate-700 rounded-lg appearance-none cursor-pointer accent-indigo-500"
              />
            </div>
          </div>
        </div>
      </div>

      {/* Info Card */}
      {epicenter && (
        <div className="bg-[#0f172a]/60 backdrop-blur-md border border-white/5 rounded-2xl p-5 animate-in slide-in-from-left duration-500">
          <div className="text-[10px] font-bold text-slate-400 uppercase tracking-widest mb-3 flex items-center gap-1.5">
            <Waves size={12} className="text-cyan-500" /> Wave Velocity Info
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-1">
              <span className="text-[9px] text-slate-500 uppercase font-black">P-Wave</span>
              <p className="text-xs text-blue-400 font-bold">~7.0 km/s</p>
            </div>
            <div className="space-y-1">
              <span className="text-[9px] text-slate-500 uppercase font-black">S-Wave</span>
              <p className="text-xs text-red-400 font-bold">~4.0 km/s</p>
            </div>
            <div className="space-y-1">
              <span className="text-[9px] text-slate-500 uppercase font-black">Tsunami</span>
              <p className="text-xs text-cyan-400 font-bold">~200 m/s</p>
            </div>
            <div className="space-y-1">
              <span className="text-[9px] text-slate-500 uppercase font-black">Distance</span>
              <p className="text-xs text-white font-bold">
                {(Math.sqrt(epicenter.x ** 2 + epicenter.z ** 2) / 1000).toFixed(1)} km
              </p>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
