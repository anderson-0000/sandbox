import React from 'react';
import { useStore } from '../../hooks/useStore';

const AddressDisplay: React.FC = () => {
  const viewMode = useStore((state) => state.viewMode);

  const addressParts = [
    { id: 'virgo_supercluster', name: 'おとめ座超銀河団' },
    { id: 'local_group', name: '局所銀河群' },
    { id: 'milky_way', name: '天の川銀河' },
    { id: 'orion_arm', name: 'オリオン腕' },
    { id: 'solar_system', name: '太陽系' },
    { id: 'earth', name: '地球' },
  ];

  // Determine current "depth"
  const currentIndex = addressParts.findIndex(p => p.id === viewMode);
  
  return (
    <div className="fixed top-8 left-1/2 -translate-x-1/2 flex items-center gap-2 bg-black/40 backdrop-blur-md px-6 py-2 rounded-full border border-white/10 z-[120] pointer-events-none select-none">
      {addressParts.map((part, i) => {
        // We want to show the full hierarchy up to the current level
        const isHighlight = i === currentIndex;

        return (
          <React.Fragment key={part.id}>
            <span 
              className={`text-[10px] font-bold tracking-widest transition-all duration-500 ${
                isHighlight ? 'text-blue-400 scale-110 opacity-100' : 'text-white/40 opacity-70'
              }`}
            >
              {part.name}
            </span>
            {i < addressParts.length - 1 && (
              <span className="text-[8px] text-white/10 mx-1">/</span>
            )}
          </React.Fragment>
        );
      })}
    </div>
  );
};

export default AddressDisplay;