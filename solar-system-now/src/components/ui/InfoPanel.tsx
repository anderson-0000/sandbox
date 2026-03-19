import React from 'react';
import { useStore } from '../../hooks/useStore';
import { PLANETS_DATA, MOONS_DATA } from '../../engine/kepler';
import { X } from 'lucide-react';

const InfoPanel: React.FC = () => {
  const selectedPlanet = useStore((state) => state.selectedPlanet);
  const setSelectedPlanet = useStore((state) => state.setSelectedPlanet);

  if (!selectedPlanet) return null;

  let data = null;
  let isMoon = false;
  
  if (selectedPlanet === 'Sun') {
    data = null;
  } else {
    data = PLANETS_DATA[selectedPlanet as string];
    if (!data) {
      // Look in moons
      for (const parent in MOONS_DATA) {
        const found = MOONS_DATA[parent].find((m: any) => m.id === selectedPlanet);
        if (found) {
          data = found;
          isMoon = true;
          break;
        }
      }
    }
  }

  return (
    <div className="absolute top-8 left-8 w-72 bg-black/80 backdrop-blur-xl p-6 rounded-2xl border border-white/20 text-white shadow-2xl z-10 transition-all pointer-events-auto">
      <div className="flex justify-between items-center mb-6">
        <div className="flex flex-col">
          <h2 className="text-2xl font-bold tracking-tight">{selectedPlanet === 'Sun' ? '太陽' : data?.name}</h2>
          {isMoon && (data as any).parent && (
            <p className="text-[10px] uppercase opacity-100 font-mono tracking-widest text-white">
              {(PLANETS_DATA[(data as any).parent]?.name || (data as any).parent)}の衛星
            </p>
          )}
        </div>
        <button 
          onClick={() => setSelectedPlanet(null)}
          className="p-1 hover:bg-white/10 rounded-lg transition-colors"
        >
          <X className="w-5 h-5 opacity-100 hover:opacity-100 text-white" />
        </button>
      </div>

      <div className="space-y-4">
        {selectedPlanet === 'Sun' ? (
          <>
            <p className="text-sm opacity-100 leading-relaxed text-white">
              太陽系の中心にある恒星。太陽系の質量の約99.8%を占める。
            </p>
            <div className="grid grid-cols-2 gap-4 pt-4 border-t border-white/10">
              <div>
                <p className="text-[10px] uppercase opacity-100 font-mono tracking-widest text-white">分類</p>
                <p className="text-sm font-medium">G型主系列星</p>
              </div>
              <div>
                <p className="text-[10px] uppercase opacity-100 font-mono tracking-widest text-white">表面温度</p>
                <p className="text-sm font-medium">5,778 K</p>
              </div>
            </div>
          </>
        ) : data ? (
          <>
            <p className="text-sm opacity-100 leading-relaxed text-white">
              {data.name}は太陽から約{data.a.toFixed(2)} AUの距離を公転する惑星です。
            </p>
            <div className="space-y-4 pt-4 border-t border-white/10">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-[10px] uppercase opacity-100 font-mono tracking-widest text-white">太陽からの距離</p>
                  <p className="text-sm font-medium">{data.a.toFixed(3)} AU</p>
                </div>
                <div>
                  <p className="text-[10px] uppercase opacity-100 font-mono tracking-widest text-white">離心率</p>
                  <p className="text-sm font-medium">{data.e.toFixed(4)}</p>
                </div>
              </div>
              <div>
                <p className="text-[10px] uppercase opacity-100 font-mono tracking-widest text-white">軌道傾斜角</p>
                <p className="text-sm font-medium">{data.i.toFixed(3)}°</p>
              </div>
            </div>
          </>
        ) : null}
      </div>
      
      <div className="mt-8">
        <div className="h-1.5 w-full bg-white/10 rounded-full overflow-hidden">
          <div 
            className="h-full bg-white/60 transition-all duration-500" 
            style={{ width: selectedPlanet === 'Sun' ? '100%' : `${Math.min(100, (data?.a || 0) * 3)}%` }}
          />
        </div>
      </div>
    </div>
  );
};

export default InfoPanel;
