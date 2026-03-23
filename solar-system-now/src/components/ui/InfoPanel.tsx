import React from 'react';
import { useStore, type ViewMode } from '../../hooks/useStore';
import { PLANETS_DATA, MOONS_DATA } from '../../engine/kepler';
import { X, Eye, Globe, Navigation, Settings2, Sparkles, CheckCircle2, Circle, Search, Compass } from 'lucide-react';

const INTERESTING_EVENTS = [
  { name: '皆既日食 (2026)', date: new Date('2026-08-12T18:00:00Z'), target: 'Earth' },
  { name: 'アポロ11号月面着陸', date: new Date('1969-07-20T20:17:00Z'), target: 'Moon' },
  { name: '惑星直列 (2040)', date: new Date('2040-09-08T12:00:00Z'), target: 'Jupiter' },
  { name: '火星大接近 (2035)', date: new Date('2035-09-15T12:00:00Z'), target: 'Mars' },
];

const VIEW_MODE_DESCRIPTIONS: Record<string, { title: string, description: string }> = {
  earth: { title: '地球視点', description: '太陽系第3惑星。唯一生命が確認されている惑星です。このモードでは地球を追跡し、その周囲の動きを観察します。' },
  solar_system: { title: '太陽系', description: '太陽とその周りを回る惑星・衛星の集まり。銀河系内を秒速約230kmで公転しています。' },
  orion_arm: { title: 'オリオン腕', description: '天の川銀河の螺旋状の「腕」の一つ。私たちの太陽系はこの腕の内側に位置しています。' },
  milky_way: { title: '天の川銀河', description: '約2000億個の恒星が集まる棒渦巻銀河。中心には巨大ブラックホール「いて座A*」が存在します。' },
  local_group: { title: '局所銀河群', description: '天の川銀河、アンドロメダ銀河、さんかく座銀河などを含む約50個以上の銀河の集まりです。' },
  virgo_supercluster: { title: 'おとめ座超銀河団', description: '局所銀河群を含む巨大な銀河の集団。重力の中心であるグレート・アトラクターへ向かっています。' },
  galactic: { title: '銀河追走', description: '太陽の銀河内での移動方向に基づき、太陽系を前方からダイナミックに追従します。' },
};

const MIN_ZOOM = 0.001;
const MAX_ZOOM = 2000000000;
const MIN_LOG = Math.log(MIN_ZOOM);
const MAX_LOG = Math.log(MAX_ZOOM);

const ToggleSwitch: React.FC<{ label: string, active: boolean, onClick: () => void }> = ({ label, active, onClick }) => (
  <button 
    onClick={onClick}
    className={`w-full flex items-center justify-between px-3 py-2 rounded-xl text-[10px] font-bold transition-all border ${
      active 
        ? 'bg-blue-600/20 border-blue-500/50 text-blue-100 shadow-[inset_0_0_10px_rgba(59,130,246,0.2)]' 
        : 'bg-white/5 border-white/10 text-white/40 hover:bg-white/10'
    }`}
  >
    <div className="flex items-center gap-2">
      {active ? <CheckCircle2 className="w-3 h-3 text-blue-400" /> : <Circle className="w-3 h-3 opacity-40" />}
      {label}
    </div>
    <div className={`text-[8px] px-1 py-0.5 rounded uppercase tracking-tighter ${active ? 'bg-blue-500 text-white' : 'bg-white/10 text-white/40'}`}>
      {active ? 'ON' : 'OFF'}
    </div>
  </button>
);

const InfoPanel: React.FC = () => {
  const { 
    selectedObjectName, setSelectedObjectName, 
    viewMode, setViewMode,
    surfaceTargetName, setSurfaceTargetName,
    setCurrentDate,
    showOrbits, setShowOrbits,
    showMoonOrbits, setShowMoonOrbits,
    showSunOrbit, setShowSunOrbit,
    showLabels, setShowLabels,
    showVectors, setShowVectors,
    zoomDistance, setZoomDistance
  } = useStore();

  const handleToggleViewMode = (mode: ViewMode) => {
    if (mode === 'earth') {
        setViewMode(viewMode === 'earth' ? 'solar_system' : 'earth');
    } else {
        setViewMode(viewMode === mode ? 'solar_system' : mode);
    }
  };

  const handleZoomChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const val = parseFloat(e.target.value);
    const distance = Math.exp(MIN_LOG + (val / 100) * (MAX_LOG - MIN_LOG));
    setZoomDistance(distance);
  };

  const currentZoomProgress = ((Math.log(zoomDistance) - MIN_LOG) / (MAX_LOG - MIN_LOG)) * 100;

  const allTargets: { id: string, name: string }[] = [{ id: 'Sun', name: '太陽' }];
  Object.values(PLANETS_DATA).forEach(p => {
    if (p.id !== selectedObjectName) allTargets.push({ id: p.id, name: p.name });
    const moons = MOONS_DATA[p.id] || [];
    moons.forEach(m => {
      if (m.id !== selectedObjectName) allTargets.push({ id: m.id, name: m.name });
    });
  });

  const renderDisplaySettings = () => (
    <div className="space-y-1.5">
      <ToggleSwitch label="太陽の軌道" active={showSunOrbit} onClick={() => setShowSunOrbit(!showSunOrbit)} />
      <ToggleSwitch label="惑星の軌道" active={showOrbits} onClick={() => setShowOrbits(!showOrbits)} />
      <ToggleSwitch label="衛星の軌道" active={showMoonOrbits} onClick={() => setShowMoonOrbits(!showMoonOrbits)} />
      <ToggleSwitch label="移動方向 (矢印)" active={showVectors} onClick={() => setShowVectors(!showVectors)} />
      <ToggleSwitch label="ラベル表示" active={showLabels} onClick={() => setShowLabels(!showLabels)} />
    </div>
  );

  if (!selectedObjectName) {
    const info = VIEW_MODE_DESCRIPTIONS[viewMode] || { title: '宇宙', description: '広大な宇宙のシミュレーションです。' };
    
    return (
      <div className="absolute bottom-8 left-8 w-80 bg-black/80 backdrop-blur-xl p-6 rounded-2xl border border-white/20 text-white shadow-2xl z-[110] transition-all pointer-events-auto max-h-[85vh] overflow-y-auto scrollbar-hide">
        <h2 className="text-2xl font-bold mb-2 flex items-center gap-2">
          <Sparkles className="w-5 h-5 text-yellow-400" />
          {info.title}
        </h2>
        <p className="text-sm opacity-100 leading-relaxed mb-6">
          {info.description}
        </p>

        <div className="space-y-2 mb-6">
          <p className="text-[10px] uppercase font-bold opacity-60 flex items-center gap-2 mb-1">
            <Sparkles className="w-3.5 h-3.5" /> 天文イベント
          </p>
          {INTERESTING_EVENTS.map((event) => (
            <button
              key={event.name}
              onClick={() => {
                setCurrentDate(event.date);
                setSelectedObjectName(event.target);
              }}
              className="w-full text-left px-3 py-2 rounded-lg bg-white/5 hover:bg-white/10 border border-white/10 transition-colors text-xs"
            >
              <div className="font-bold">{event.name}</div>
              <div className="text-[10px] opacity-60 font-mono">{event.date.toLocaleDateString()}</div>
            </button>
          ))}
        </div>

        <div className="mt-6 pt-6 border-t border-white/10 space-y-4">
          <button 
            onClick={() => handleToggleViewMode('galactic')}
            className={`w-full flex items-center justify-center gap-2 px-4 py-3 rounded-xl text-xs font-bold transition-all ${
              viewMode === 'galactic' ? 'bg-indigo-600 text-white shadow-[0_0_20px_rgba(79,70,229,0.4)]' : 'bg-white/5 hover:bg-white/10'
            }`}
          >
            <Compass className="w-4 h-4" />
            銀河追走モード {viewMode === 'galactic' ? '解除' : '開始'}
          </button>

          <h3 className="text-[10px] uppercase font-bold opacity-60 flex items-center gap-2">
            <Search className="w-3.5 h-3.5" /> ズーム
          </h3>
          <input type="range" min="0" max="100" step="0.1" value={currentZoomProgress} onChange={handleZoomChange} className="w-full h-1 bg-white/10 rounded-lg appearance-none cursor-pointer accent-blue-500" />

          <h3 className="text-[10px] uppercase font-bold opacity-60 flex items-center gap-2 pt-2">
            <Settings2 className="w-3.5 h-3.5" /> 表示設定
          </h3>
          {renderDisplaySettings()}
        </div>
      </div>
    );
  }

  let data = null;
  let isMoon = false;
  if (selectedObjectName !== 'Sun') {
    data = PLANETS_DATA[selectedObjectName as string];
    if (!data) {
      for (const parent in MOONS_DATA) {
        const found = MOONS_DATA[parent].find((m: any) => m.id === selectedObjectName);
        if (found) { data = found; isMoon = true; break; }
      }
    }
  }

  return (
    <div className="absolute bottom-8 left-8 w-80 bg-black/80 backdrop-blur-xl p-6 rounded-2xl border border-white/20 text-white shadow-2xl z-[110] transition-all pointer-events-auto max-h-[85vh] overflow-y-auto scrollbar-hide">
      <div className="flex justify-between items-center mb-6">
        <div className="flex flex-col">
          <h2 className="text-2xl font-bold tracking-tight">{selectedObjectName === 'Sun' ? '太陽' : data?.name}</h2>
          {isMoon && (data as any).parent && <p className="text-[10px] uppercase opacity-60 font-mono tracking-widest">{(PLANETS_DATA[(data as any).parent]?.name || (data as any).parent)}の衛星</p>}
        </div>
        <button onClick={() => setSelectedObjectName(null)} className="p-1 hover:bg-white/10 rounded-lg transition-colors"><X className="w-5 h-5 opacity-60 hover:opacity-100" /></button>
      </div>

      <div className="space-y-6">
        <div className="space-y-4">
          <p className="text-sm opacity-100 leading-relaxed">
            {selectedObjectName === 'Sun' ? '太陽系の中心にある恒星。太陽系の質量の約99.8%を占める。' : `${data?.name}は太陽から約${data?.a.toFixed(2)} AUの距離を公転する${isMoon ? '衛星' : '惑星'}です。`}
          </p>
        </div>

        <div className="pt-4 border-t border-white/10 space-y-3">
          {(selectedObjectName === 'Earth' || selectedObjectName === 'Sun') && (
            <button 
              onClick={() => handleToggleViewMode('earth')}
              className={`w-full flex items-center justify-center gap-2 px-4 py-3 rounded-2xl text-sm font-bold transition-all ${
                viewMode === 'earth' ? 'bg-blue-600 text-white shadow-[0_0_20px_rgba(37,99,235,0.4)]' : 'bg-white/10 hover:bg-white/20'
              }`}
            >
              {viewMode === 'earth' ? <Globe className="w-4 h-4" /> : <Navigation className="w-4 h-4" />}
              {viewMode === 'earth' ? '軌道に戻る' : (selectedObjectName === 'Sun' ? '光球に降りる' : '地表に降りる')}
            </button>
          )}

          <button 
            onClick={() => handleToggleViewMode('galactic')}
            className={`w-full flex items-center justify-center gap-2 px-4 py-3 rounded-xl text-xs font-bold transition-all ${
              viewMode === 'galactic' ? 'bg-indigo-600 text-white shadow-[0_0_20px_rgba(79,70,229,0.4)]' : 'bg-white/5 hover:bg-white/10'
            }`}
          >
            <Compass className="w-4 h-4" />
            銀河追走モード {viewMode === 'galactic' ? '解除' : '開始'}
          </button>

          {viewMode === 'earth' && (
            <div className="space-y-2 animate-in fade-in slide-in-from-top-2 duration-300">
              <p className="text-[10px] uppercase font-bold opacity-60 tracking-wider flex items-center gap-1"><Eye className="w-3 h-3" /> 追尾ターゲットを選択</p>
              <div className="flex flex-wrap gap-1.5 max-h-40 overflow-y-auto pr-2 scrollbar-hide border border-white/5 rounded-lg p-2 bg-white/5">
                <button onClick={() => setSurfaceTargetName(null)} className={`px-2 py-1 rounded-md text-[10px] border transition-all ${surfaceTargetName === null ? 'bg-white text-black border-white' : 'bg-transparent border-white/20 hover:border-white/40'}`}>真上</button>
                {allTargets.map(target => (
                  <button key={target.id} onClick={() => setSurfaceTargetName(target.id)} className={`px-2 py-1 rounded-md text-[10px] border transition-all ${surfaceTargetName === target.id ? 'bg-white text-black border-white' : 'bg-transparent border-white/20 hover:border-white/40'}`}>{target.name}</button>
                ))}
              </div>
            </div>
          )}
        </div>

        <div className="pt-6 border-t border-white/10 space-y-4">
          <h3 className="text-[10px] uppercase font-bold opacity-60 tracking-wider flex items-center gap-1"><Search className="w-3.5 h-3.5" /> ズーム</h3>
          <input type="range" min="0" max="100" step="0.1" value={currentZoomProgress} onChange={handleZoomChange} className="w-full h-1 bg-white/10 rounded-lg appearance-none cursor-pointer accent-blue-500" />
          <h3 className="text-[10px] uppercase font-bold opacity-60 tracking-wider flex items-center gap-1 pt-2"><Settings2 className="w-3.5 h-3.5" /> 表示設定</h3>
          {renderDisplaySettings()}
        </div>
      </div>
    </div>
  );
};

export default InfoPanel;