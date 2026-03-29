import React from 'react';
import { useProjectStore } from '../../store/useProjectStore';
import { 
  Square, Car, Video, Map, Plus, 
  Layers, Utensils, Bath, DoorOpen, Layout, Frame, MousePointer2, Sparkles
} from 'lucide-react';
import type { ObjectType } from '../../types';

const Toolbar: React.FC = () => {
  const { addObject, setSelectedId, selectedId } = useProjectStore();

  const handleAddObject = (type: ObjectType) => {
    const id = crypto.randomUUID();
    let name = '新規オブジェクト';
    let dimensions = { widthMm: 1820, heightMm: 1820 };
    let color = '#f8fafc';

    switch (type) {
      case 'room':
        name = '部屋';
        dimensions = { widthMm: 3640, heightMm: 3640 };
        color = '#f1f5f9';
        break;
      case 'garage':
        name = 'ガレージ';
        dimensions = { widthMm: 5460, heightMm: 6000 };
        color = '#f8fafc';
        break;
      case 'vehicle':
        name = '大型SUV';
        dimensions = { widthMm: 1920, heightMm: 4750 };
        color = '#fbbf24';
        break;
      case 'equipment':
        name = 'プロジェクター';
        dimensions = { widthMm: 400, heightMm: 300 };
        color = '#f87171';
        break;
      case 'stairs':
        name = '階段';
        dimensions = { widthMm: 910, heightMm: 1820 };
        break;
      case 'kitchen':
        name = 'キッチン';
        dimensions = { widthMm: 2550, heightMm: 650 };
        break;
      case 'bath':
        name = 'お風呂';
        dimensions = { widthMm: 1600, heightMm: 1600 };
        color = '#f0f9ff';
        break;
      case 'entrance':
        name = '玄関';
        dimensions = { widthMm: 1820, heightMm: 1820 };
        break;
      case 'door':
        name = 'ドア';
        dimensions = { widthMm: 800, heightMm: 100 };
        break;
      case 'window':
        name = '窓';
        dimensions = { widthMm: 1600, heightMm: 100 };
        break;
    }

    addObject({
      id,
      type,
      name,
      dimensions,
      position: { x: 5000, y: 5000 },
      rotation: 0,
      properties: { color, floorLevel: 1 },
    });
    setSelectedId(id);
  };

  const sections = [
    {
      title: "Basic",
      color: "from-blue-500 to-indigo-600",
      tools: [
        { type: 'room' as ObjectType, icon: Square, label: '部屋' },
        { type: 'entrance' as ObjectType, icon: DoorOpen, label: '玄関' },
        { type: 'stairs' as ObjectType, icon: Layers, label: '階段' },
      ]
    },
    {
      title: "Water",
      color: "from-cyan-400 to-blue-500",
      tools: [
        { type: 'kitchen' as ObjectType, icon: Utensils, label: 'キッチン' },
        { type: 'bath' as ObjectType, icon: Bath, label: 'お風呂' },
      ]
    },
    {
      title: "Opening",
      color: "from-amber-400 to-orange-500",
      tools: [
        { type: 'door' as ObjectType, icon: Layout, label: 'ドア' },
        { type: 'window' as ObjectType, icon: Frame, label: '窓' },
      ]
    },
    {
      title: "Other",
      color: "from-rose-400 to-pink-600",
      tools: [
        { type: 'garage' as ObjectType, icon: Map, label: 'ガレージ' },
        { type: 'vehicle' as ObjectType, icon: Car, label: '車両' },
        { type: 'equipment' as ObjectType, icon: Video, label: 'シアター' },
      ]
    }
  ];

  return (
    <div className="w-24 bg-white/60 backdrop-blur-xl border-r border-white/50 flex flex-col items-center py-8 gap-8 z-40 overflow-y-auto scrollbar-hide shadow-[10px_0_30px_rgba(0,0,0,0.02)]">
      <button 
        onClick={() => setSelectedId(null)}
        className={`w-14 h-14 flex items-center justify-center rounded-[20px] transition-all duration-300 group relative ${
          !selectedId 
            ? 'bg-gradient-to-br from-indigo-600 to-purple-700 text-white shadow-lg shadow-indigo-200' 
            : 'bg-white text-slate-400 hover:text-indigo-600 hover:shadow-md border border-slate-100'
        }`}
        title="選択モード"
      >
        <MousePointer2 size={24} className={!selectedId ? 'animate-pulse' : ''} />
        {!selectedId && <Sparkles size={12} className="absolute top-2 right-2 text-indigo-200" />}
      </button>

      <div className="w-12 h-[2px] bg-gradient-to-r from-transparent via-slate-200 to-transparent" />

      {sections.map((section, idx) => (
        <div key={idx} className="flex flex-col items-center gap-4 w-full px-4">
          <span className="text-[10px] font-black text-slate-400 uppercase tracking-[0.15em]">{section.title}</span>
          <div className="flex flex-col gap-3">
            {section.tools.map((tool) => (
              <button
                key={tool.type}
                onClick={() => handleAddObject(tool.type)}
                className="group w-14 h-14 flex items-center justify-center bg-white border border-slate-100 rounded-[20px] hover:border-transparent hover:shadow-[0_10px_20px_-5px_rgba(0,0,0,0.1)] active:scale-90 transition-all duration-300 relative overflow-hidden"
                title={tool.label}
              >
                <div className={`absolute inset-0 bg-gradient-to-br ${section.color} opacity-0 group-hover:opacity-10 transition-opacity duration-300`} />
                <tool.icon size={22} strokeWidth={1.5} className="text-slate-500 group-hover:text-slate-900 group-hover:scale-110 transition-all duration-300 z-10" />
              </button>
            ))}
          </div>
        </div>
      ))}

      <button className="mt-auto w-12 h-12 flex items-center justify-center text-slate-300 hover:text-indigo-500 hover:scale-125 transition-all">
        <Plus size={24} />
      </button>
    </div>
  );
};

export default Toolbar;
