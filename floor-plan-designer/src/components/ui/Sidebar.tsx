import React from 'react';
import { useProjectStore } from '../../store/useProjectStore';
import { Trash2, Settings2, Sliders, Palette, Maximize2, Check, Ruler } from 'lucide-react';

const Sidebar: React.FC = () => {
  const { 
    selectedId, objects, updateObject, removeObject, setSelectedId,
    projectSettings, setWallThickness 
  } = useProjectStore();
  
  const selectedObject = objects.find(obj => obj.id === selectedId);

  const handleChange = (key: string, value: any) => {
    if (selectedObject) {
      updateObject(selectedObject.id, { [key]: value });
    }
  };

  const handleDimensionChange = (key: string, value: number) => {
    if (selectedObject) {
      updateObject(selectedObject.id, {
        dimensions: { ...selectedObject.dimensions, [key]: value }
      });
    }
  };

  const colors = [
    { value: '#ffffff', name: 'White', bg: 'bg-white' },
    { value: '#f8fafc', name: 'Cloud', bg: 'bg-slate-50' },
    { value: '#e0f2fe', name: 'Sky', bg: 'bg-sky-100' },
    { value: '#dcfce7', name: 'Mint', bg: 'bg-emerald-100' },
    { value: '#fef3c7', name: 'Amber', bg: 'bg-amber-100' },
    { value: '#fee2e2', name: 'Rose', bg: 'bg-rose-100' },
    { value: '#f3e8ff', name: 'Grape', bg: 'bg-purple-100' },
    { value: '#1e293b', name: 'Deep', bg: 'bg-slate-800' },
  ];

  return (
    <div className="w-80 bg-white/70 backdrop-blur-xl border-l border-white/50 flex flex-col z-40 overflow-hidden shadow-[-10px_0_30px_rgba(0,0,0,0.02)]">
      <div className="flex-1 overflow-y-auto p-7 space-y-10 scrollbar-hide">
        {/* Global Settings */}
        <div className="space-y-5">
          <h2 className="text-[11px] font-black text-indigo-500 uppercase tracking-[0.2em] flex items-center gap-2">
            <Settings2 size={16} className="text-indigo-400" /> Project Settings
          </h2>
          <div className="bg-gradient-to-br from-indigo-50 to-white p-5 rounded-[24px] space-y-5 border border-indigo-100/50 shadow-sm">
            <div className="space-y-4">
              <div className="flex justify-between items-end">
                <label className="text-[11px] font-black text-slate-600 flex items-center gap-1.5">
                  <Ruler size={12} /> 壁の厚み
                </label>
                <span className="text-[12px] font-black text-indigo-600 bg-white px-2.5 py-1 rounded-lg border border-indigo-100 shadow-sm">{projectSettings.wallThicknessMm}mm</span>
              </div>
              <input
                type="range"
                min="50"
                max="300"
                step="10"
                value={projectSettings.wallThicknessMm}
                onChange={(e) => setWallThickness(Number(e.target.value))}
                className="w-full h-2 bg-indigo-100 rounded-full appearance-none cursor-pointer accent-indigo-600"
              />
            </div>
          </div>
        </div>

        {/* Object Properties */}
        <div className="space-y-5">
          <h2 className="text-[11px] font-black text-rose-500 uppercase tracking-[0.2em] flex items-center gap-2">
            <Sliders size={16} className="text-rose-400" /> Properties
          </h2>
          
          {!selectedObject ? (
            <div className="bg-white/50 border-2 border-dashed border-slate-200 rounded-[2.5rem] p-12 flex flex-col items-center justify-center gap-5 text-center transition-all duration-500 hover:border-indigo-300 hover:bg-indigo-50/30 group">
              <div className="w-16 h-16 bg-white rounded-3xl flex items-center justify-center shadow-lg text-slate-300 group-hover:text-indigo-400 group-hover:scale-110 group-hover:rotate-12 transition-all duration-500">
                <Maximize2 size={24} />
              </div>
              <p className="text-[12px] font-bold text-slate-400 leading-relaxed tracking-tight group-hover:text-slate-600 transition-colors">
                キャンバス上のオブジェクトを<br/>選択して編集を開始
              </p>
            </div>
          ) : (
            <div className="space-y-8 animate-in fade-in slide-in-from-right-8 duration-500">
              <div className="bg-white/80 border border-slate-200/50 rounded-[2rem] p-6 shadow-sm space-y-6">
                <div className="space-y-2">
                  <label className="text-[10px] font-black text-slate-400 uppercase ml-1 tracking-wider">Object Name</label>
                  <input
                    type="text"
                    value={selectedObject.name}
                    onChange={(e) => handleChange('name', e.target.value)}
                    className="w-full px-5 py-3.5 bg-slate-50 rounded-2xl text-sm font-bold focus:ring-2 focus:ring-indigo-500 focus:bg-white outline-none transition-all border border-transparent shadow-inner"
                  />
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-[10px] font-black text-slate-400 uppercase ml-1 tracking-wider">Width (mm)</label>
                    <input
                      type="number"
                      value={selectedObject.dimensions.widthMm}
                      onChange={(e) => handleDimensionChange('widthMm', Number(e.target.value))}
                      className="w-full px-5 py-3.5 bg-slate-50 rounded-2xl text-sm font-bold outline-none border border-transparent focus:bg-white focus:border-indigo-100 transition-all shadow-inner"
                    />
                  </div>
                  <div className="space-y-2">
                    <label className="text-[10px] font-black text-slate-400 uppercase ml-1 tracking-wider">Depth (mm)</label>
                    <input
                      type="number"
                      value={selectedObject.dimensions.heightMm}
                      onChange={(e) => handleDimensionChange('heightMm', Number(e.target.value))}
                      className="w-full px-5 py-3.5 bg-slate-50 rounded-2xl text-sm font-bold outline-none border border-transparent focus:bg-white focus:border-indigo-100 transition-all shadow-inner"
                    />
                  </div>
                </div>

                <div className="space-y-3 pt-2">
                  <div className="flex justify-between items-center mb-1">
                    <label className="text-[10px] font-black text-slate-400 uppercase ml-1 tracking-wider">Rotation</label>
                    <span className="text-[11px] font-black text-indigo-600 bg-indigo-50 px-2.5 py-1 rounded-lg border border-indigo-100">{selectedObject.rotation}°</span>
                  </div>
                  <input
                    type="range"
                    min="0"
                    max="315"
                    step="45"
                    value={selectedObject.rotation}
                    onChange={(e) => handleChange('rotation', Number(e.target.value))}
                    className="w-full h-2 bg-slate-100 rounded-full appearance-none cursor-pointer accent-indigo-600"
                  />
                </div>
              </div>

              {/* Enhanced Color Swatch */}
              <div className="space-y-5">
                <h3 className="text-[11px] font-black text-slate-500 uppercase flex items-center gap-2 ml-1 tracking-widest">
                  <Palette size={16} className="text-amber-400" /> Color Swatch
                </h3>
                <div className="grid grid-cols-4 gap-4">
                  {colors.map(color => (
                    <button
                      key={color.value}
                      onClick={() => handleChange('properties', { ...selectedObject.properties, color: color.value })}
                      className={`relative aspect-square rounded-2xl border-2 transition-all duration-300 flex items-center justify-center group ${
                        selectedObject.properties.color === color.value 
                          ? 'border-indigo-600 bg-white shadow-xl scale-110 z-10' 
                          : 'border-white hover:border-indigo-200 hover:scale-105 shadow-sm'
                      }`}
                      style={{ backgroundColor: color.value }}
                      title={color.name}
                    >
                      {selectedObject.properties.color === color.value && (
                        <Check size={18} className={`${color.value === '#1e293b' ? 'text-white' : 'text-indigo-600'} animate-in zoom-in duration-300`} />
                      )}
                    </button>
                  ))}
                </div>
              </div>

              <div className="pt-6">
                <button
                  onClick={() => {
                    removeObject(selectedObject.id);
                    setSelectedId(null);
                  }}
                  className="w-full flex items-center justify-center gap-3 py-4 bg-rose-50 text-rose-600 rounded-2xl text-[13px] font-black hover:bg-rose-100 active:scale-[0.98] transition-all border border-rose-100"
                >
                  <Trash2 size={18} />
                  Delete Object
                </button>
              </div>
            </div>
          )}
        </div>
      </div>
      
      {/* Footer info */}
      <div className="p-7 border-t border-white/50 bg-white/30 backdrop-blur-md">
        <div className="flex items-center justify-between text-[10px] font-black text-slate-400 uppercase tracking-[0.2em]">
          <span>REIWA V2.0</span>
          <span className="flex items-center gap-2">
            <span className="w-2.5 h-2.5 bg-emerald-400 rounded-full animate-pulse shadow-[0_0_10px_rgba(52,211,153,0.6)]" />
            System Active
          </span>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
