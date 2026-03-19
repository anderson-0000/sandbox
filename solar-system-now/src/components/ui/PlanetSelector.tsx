import React from 'react';
import { useStore } from '../../hooks/useStore';
import { PLANETS_DATA, MOONS_DATA } from '../../engine/kepler';

const PlanetSelector: React.FC = () => {
  const { selectedPlanet, setSelectedPlanet } = useStore();

  return (
    <div className="absolute top-8 right-8 flex flex-col items-end gap-2 z-[60] pointer-events-auto h-[80vh] overflow-y-auto pr-4 scrollbar-hide">
      <div className="text-white text-[10px] uppercase tracking-widest font-bold mr-2 mb-1 drop-shadow-lg">
        フォーカス対象
      </div>
      
      <div className="flex flex-col gap-4 items-end">
        {/* Sun */}
        <button
          onClick={() => setSelectedPlanet(selectedPlanet === 'Sun' ? null : 'Sun')}
          className={`px-4 py-2 rounded-full text-xs font-bold transition-all border ${
            selectedPlanet === 'Sun'
              ? 'bg-yellow-500 text-white border-yellow-400 shadow-[0_0_20px_rgba(253,184,19,0.8)] scale-110'
              : 'bg-black/60 text-white border-white/20 hover:bg-black/80'
          }`}
        >
          太陽
        </button>

        {/* Planets and Moons */}
        {Object.values(PLANETS_DATA).map((planet) => {
          const moons = MOONS_DATA[planet.id] || [];
          const isSelected = selectedPlanet === planet.id;

          return (
            <div key={planet.id} className="flex flex-col items-end gap-2">
              <button
                onClick={() => setSelectedPlanet(isSelected ? null : planet.id)}
                className={`px-4 py-2 rounded-full text-xs font-bold transition-all border ${
                  isSelected
                    ? 'bg-blue-600 text-white border-blue-400 shadow-[0_0_20px_rgba(34,113,179,0.8)] scale-110'
                    : 'bg-black/60 text-white border-white/20 hover:bg-black/80'
                }`}
              >
                {planet.name}
              </button>

              {/* Moons indented */}
              {moons.length > 0 && (
                <div className="flex flex-wrap gap-1.5 justify-end max-w-[200px] mr-2">
                  {moons.map((moon) => {
                    const isMoonSelected = selectedPlanet === moon.id;
                    return (
                      <button
                        key={moon.id}
                        onClick={() => setSelectedPlanet(isMoonSelected ? null : moon.id)}
                        className={`px-2 py-1 rounded-md text-[9px] font-bold transition-all border ${
                          isMoonSelected
                            ? 'bg-white/30 text-white border-white scale-105 backdrop-blur-md'
                            : 'bg-white/10 text-white border-white/10 hover:bg-white/20'
                        }`}
                      >
                        {moon.name}
                      </button>
                    );
                  })}
                </div>
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default PlanetSelector;
