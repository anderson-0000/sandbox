import React from 'react';
import { useStore, type ViewMode } from '../../hooks/useStore';
import { Layers, Globe, Sun, Wind, Aperture, Box, Compass } from 'lucide-react';

const ViewModeSelector: React.FC = () => {
  const { viewMode, setViewMode, showVectors, setShowVectors } = useStore();

  const modes: { id: ViewMode; name: string; icon: React.ReactNode }[] = [
    { id: 'earth', name: '地球', icon: <Globe className="w-4 h-4" /> },
    { id: 'solar_system', name: '太陽系', icon: <Sun className="w-4 h-4" /> },
    { id: 'orion_arm', name: 'オリオン腕', icon: <Wind className="w-4 h-4" /> },
    { id: 'milky_way', name: '天の川銀河', icon: <Layers className="w-4 h-4" /> },
    { id: 'local_group', name: '局所銀河群', icon: <Aperture className="w-4 h-4" /> },
    { id: 'virgo_supercluster', name: 'おとめ座超銀河団', icon: <Box className="w-4 h-4" /> },
    { id: 'galactic', name: '銀河追尾', icon: <Compass className="w-4 h-4" /> },
  ];

  return (
    <div className="absolute top-24 left-8 flex flex-col gap-4 z-[110] pointer-events-auto max-h-[calc(100vh-12rem)] overflow-y-auto scrollbar-hide pr-2">
      <div className="flex flex-col gap-1.5 p-4 bg-black/60 backdrop-blur-xl rounded-3xl border border-white/20 shadow-2xl">
        <div className="text-white text-[10px] uppercase tracking-[0.2em] font-black mb-1 opacity-50 px-2">
          視点モード
        </div>
        {modes.map((mode) => (
          <button
            key={mode.id}
            onClick={() => setViewMode(mode.id)}
            className={`flex items-center gap-3 px-4 py-2 rounded-2xl text-xs font-bold transition-all border ${
              viewMode === mode.id
                ? 'bg-blue-600 text-white border-blue-400 shadow-[0_0_20px_rgba(34,113,179,0.4)]'
                : 'bg-white/5 text-white/70 border-white/5 hover:bg-white/10 hover:text-white'
            }`}
          >
            {mode.icon}
            {mode.name}
          </button>
        ))}
      </div>

      <div className="flex flex-col gap-1.5 p-4 bg-black/60 backdrop-blur-xl rounded-3xl border border-white/20 shadow-2xl">
        <div className="text-white text-[10px] uppercase tracking-[0.2em] font-black mb-1 opacity-50 px-2">
          オーバーレイ
        </div>
        <button
          onClick={() => setShowVectors(!showVectors)}
          className={`flex items-center gap-3 px-4 py-2 rounded-2xl text-xs font-bold transition-all border ${
            showVectors
              ? 'bg-purple-600 text-white border-purple-400 shadow-[0_0_20px_rgba(147,51,234,0.4)]'
              : 'bg-white/5 text-white/70 border-white/5 hover:bg-white/10 hover:text-white'
          }`}
        >
          <Wind className="w-4 h-4" />
          移動方向を表示 (矢印)
        </button>
      </div>
    </div>
  );
};

export default ViewModeSelector;