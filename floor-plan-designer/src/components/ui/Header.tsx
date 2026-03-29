import React from 'react';
import { useProjectStore } from '../../store/useProjectStore';
import { calculateFloorAreaM2, calculateBuildingAreaM2 } from '../../utils/areaCalc';
import { formatArea } from '../../utils/unitConverter';
import { Save, FolderOpen, Box, ChevronDown, Sparkles } from 'lucide-react';

const Header: React.FC = () => {
  const { projectSettings, setModule, objects } = useProjectStore();
  
  const floorAreaM2 = calculateFloorAreaM2(objects);
  const buildingAreaM2 = calculateBuildingAreaM2(objects);

  return (
    <header className="h-16 bg-white/80 backdrop-blur-xl border-b border-white/50 px-8 flex items-center justify-between z-50 shadow-[0_4px_30px_rgba(0,0,0,0.03)]">
      <div className="flex items-center gap-10">
        <div className="flex items-center gap-3 group cursor-pointer">
          <div className="w-10 h-10 bg-gradient-to-br from-indigo-500 via-purple-500 to-rose-500 rounded-2xl flex items-center justify-center shadow-xl shadow-indigo-200 group-hover:rotate-12 transition-all duration-500">
            <Box size={22} className="text-white" />
          </div>
          <div className="flex flex-col -gap-1">
            <h1 className="text-[16px] font-black tracking-tight text-slate-900 leading-tight">FloorPlan<span className="bg-gradient-to-r from-indigo-600 to-rose-600 bg-clip-text text-transparent font-black italic">.studio</span></h1>
            <span className="text-[9px] font-black text-indigo-400 tracking-[0.2em] uppercase flex items-center gap-1">
              <Sparkles size={10} /> REIWA EDITION
            </span>
          </div>
        </div>

        {/* Improved Toggle Switch */}
        <div className="flex bg-slate-100/50 p-1 rounded-2xl border border-slate-200/50 w-72 h-11 relative shadow-inner gap-1">
          <div 
            className={`absolute top-1 bottom-1 w-[calc(50%-4px)] bg-gradient-to-r from-slate-800 to-slate-950 rounded-[14px] shadow-lg shadow-slate-300 transition-all duration-500 ease-[cubic-bezier(0.34,1.56,0.64,1)] ${
              projectSettings.module === 'meter' ? 'translate-x-full' : 'translate-x-0'
            }`}
          />
          
          <button
            onClick={() => setModule('shaku')}
            className={`relative z-10 flex-1 flex items-center justify-center gap-2 text-[11px] font-black tracking-wider transition-colors duration-300 ${
              projectSettings.module === 'shaku' ? 'text-white' : 'text-slate-400 hover:text-slate-600'
            }`}
          >
            {projectSettings.module === 'shaku' && (
              <span className="w-1.5 h-1.5 bg-emerald-400 rounded-full shadow-[0_0_8px_rgba(52,211,153,0.8)] animate-pulse" />
            )}
            尺 (910mm)
          </button>
          <button
            onClick={() => setModule('meter')}
            className={`relative z-10 flex-1 flex items-center justify-center gap-2 text-[11px] font-black tracking-wider transition-colors duration-300 ${
              projectSettings.module === 'meter' ? 'text-white' : 'text-slate-400 hover:text-slate-600'
            }`}
          >
            {projectSettings.module === 'meter' && (
              <span className="w-1.5 h-1.5 bg-emerald-400 rounded-full shadow-[0_0_8px_rgba(52,211,153,0.8)] animate-pulse" />
            )}
            メーター (1m)
          </button>
        </div>
      </div>

      <div className="flex items-center gap-12">
        <div className="flex gap-10">
          <div className="flex flex-col">
            <span className="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] mb-0.5">Total Floor</span>
            <span className="text-[16px] font-black text-indigo-600 tabular-nums tracking-tight">{formatArea(floorAreaM2)}</span>
          </div>
          <div className="flex flex-col border-l border-slate-100 pl-10">
            <span className="text-[9px] font-black text-slate-400 uppercase tracking-[0.2em] mb-0.5">Build Area</span>
            <span className="text-[16px] font-black text-rose-500 tabular-nums tracking-tight">{formatArea(buildingAreaM2)}</span>
          </div>
        </div>

        <div className="flex gap-3">
          <button className="group flex items-center gap-2.5 px-6 py-2.5 bg-gradient-to-r from-indigo-600 to-indigo-800 text-white rounded-2xl text-[13px] font-bold hover:shadow-[0_8px_25px_-5px_rgba(79,70,229,0.4)] active:scale-95 transition-all shadow-lg shadow-indigo-100 border border-indigo-500/20">
            <Save size={16} className="text-indigo-200 group-hover:text-white transition-colors" />
            保存
          </button>
          <button className="flex items-center gap-2.5 px-6 py-2.5 bg-white border border-slate-200 text-slate-900 rounded-2xl text-[13px] font-bold hover:border-indigo-600 hover:text-indigo-600 active:scale-95 transition-all shadow-sm">
            <FolderOpen size={16} className="text-slate-400 group-hover:text-indigo-400" />
            読込
            <ChevronDown size={14} className="text-slate-300 ml-1" />
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;
