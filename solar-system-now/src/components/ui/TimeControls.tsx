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
    <div className="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-4 bg-black/40 backdrop-blur-md p-6 rounded-2xl border border-white/10 shadow-2xl z-10 w-96">
      <div className="text-white text-xl font-mono tracking-wider flex items-center gap-3">
        <Calendar className="w-5 h-5 opacity-70" />
        {formattedDate}
      </div>
      
      <div className="flex items-center gap-6">
        <button 
          onClick={() => setTimeSpeed(Math.max(-1000, timeSpeed - 10))}
          className="p-2 text-white/70 hover:text-white transition-colors"
          title="Slower / Reverse"
        >
          <Rewind className="w-6 h-6" />
        </button>

        <button 
          onClick={() => setIsPaused(!isPaused)}
          className="w-12 h-12 bg-white/20 backdrop-blur-md rounded-full flex items-center justify-center border border-white/30 hover:bg-white/30 hover:scale-105 transition-all active:scale-95 shadow-lg"
          title={isPaused ? "Play" : "Pause"}
        >
          {isPaused ? <Play className="w-6 h-6 text-white ml-1" /> : <Pause className="w-6 h-6 text-white" />}
        </button>

        <button 
          onClick={() => setTimeSpeed(Math.min(1000, timeSpeed + 10))}
          className="p-2 text-white hover:text-white transition-colors"
          title="Faster"
        >
          <FastForward className="w-6 h-6" />
        </button>
      </div>

      <div className="w-full flex flex-col gap-1">
        <div className="flex justify-between text-[10px] text-white font-mono uppercase tracking-tighter">
          <span>シミュレーション速度</span>
          <span>{timeSpeed.toFixed(1)} 日/秒</span>
        </div>
        <input 
          type="range" 
          min="-100" 
          max="100" 
          step="1"
          value={timeSpeed}
          onChange={(e) => setTimeSpeed(parseFloat(e.target.value))}
          className="w-full accent-white h-1 bg-white/40 rounded-lg cursor-pointer appearance-none"
        />
      </div>

      <button 
        onClick={() => {
          setCurrentDate(new Date());
          setTimeSpeed(1);
        }}
        className="text-[10px] text-white hover:text-white underline underline-offset-4 uppercase tracking-widest transition-colors"
      >
        現在時刻にリセット
      </button>
    </div>
  );
};

export default TimeControls;
