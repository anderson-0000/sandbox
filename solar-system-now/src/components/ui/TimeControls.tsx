import React from 'react';
import { Play, Pause, FastForward, Rewind, Calendar } from 'lucide-react';
import { useStore } from '../../hooks/useStore';

const TimeControls: React.FC = () => {
  const { currentDate, timeSpeed, isPaused, setIsPaused, setTimeSpeed, setCurrentDate } = useStore();

  const formattedDate = currentDate.toLocaleDateString('ja-JP', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });

  return (
    <div className="fixed bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-4 bg-black/60 backdrop-blur-xl p-6 rounded-3xl border border-white/20 shadow-[0_0_50px_rgba(0,0,0,0.5)] z-[150] w-[420px] pointer-events-auto">
      <div className="text-white text-xl font-mono tracking-wider flex items-center gap-3 bg-white/5 px-4 py-2 rounded-full border border-white/5">
        <Calendar className="w-5 h-5 opacity-70" />
        {formattedDate}
      </div>
      
      <div className="flex items-center gap-10">
        <button 
          onClick={() => setTimeSpeed(Math.max(-1000, timeSpeed - 10))}
          className="p-3 text-white/50 hover:text-white transition-all hover:scale-110 active:scale-90"
          title="Slower / Reverse"
        >
          <Rewind className="w-7 h-7" />
        </button>

        <button 
          onClick={() => setIsPaused(!isPaused)}
          className="w-16 h-16 bg-blue-600/30 backdrop-blur-xl rounded-full flex items-center justify-center border border-blue-400/50 hover:bg-blue-600/50 hover:scale-110 transition-all active:scale-95 shadow-[0_0_20px_rgba(37,99,235,0.4)]"
          title={isPaused ? "Play" : "Pause"}
        >
          {isPaused ? <Play className="w-8 h-8 text-white fill-white ml-1" /> : <Pause className="w-8 h-8 text-white fill-white" />}
        </button>

        <button 
          onClick={() => setTimeSpeed(Math.min(1000, timeSpeed + 10))}
          className="p-3 text-white/50 hover:text-white transition-all hover:scale-110 active:scale-90"
          title="Faster"
        >
          <FastForward className="w-7 h-7" />
        </button>
      </div>

      <div className="w-full space-y-3 pt-2">
        <div className="flex justify-between items-end">
          <span className="text-[10px] font-black uppercase tracking-[0.2em] text-white/40">Simulation Speed</span>
          <span className="text-sm font-mono font-bold text-blue-400">{timeSpeed.toFixed(1)} <span className="text-[10px] text-white/40">Days/Sec</span></span>
        </div>
        <div className="relative h-6 flex items-center">
          <input 
            type="range" 
            min="-100" 
            max="100" 
            step="1"
            value={timeSpeed}
            onChange={(e) => setTimeSpeed(parseFloat(e.target.value))}
            className="w-full accent-blue-500 h-1.5 bg-white/10 rounded-lg cursor-pointer appearance-none hover:bg-white/20 transition-all"
          />
        </div>
      </div>

      <button 
        onClick={() => {
          setCurrentDate(new Date());
          setTimeSpeed(1);
          setIsPaused(false);
        }}
        className="mt-2 text-[10px] text-white/60 hover:text-white underline underline-offset-8 uppercase tracking-[0.3em] font-bold transition-all hover:tracking-[0.4em]"
      >
        現在時刻にリセット
      </button>
    </div>
  );
};

export default TimeControls;
