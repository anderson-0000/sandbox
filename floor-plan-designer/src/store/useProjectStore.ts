import { create } from 'zustand';
import type { ProjectData, ProjectObject, Plot, ModuleType, Position } from '../types';

export interface ProjectSettings {
  module: ModuleType;
  gridSizeMm: number;
  wallThicknessMm: number;
}

interface ProjectState extends ProjectData {
  selectedId: string | null;
  stageScale: number;
  stagePos: { x: number; y: number };
  setSelectedId: (id: string | null) => void;
  setStageScale: (scale: number) => void;
  setStagePos: (pos: { x: number; y: number }) => void;
  setModule: (module: ModuleType) => void;
  setWallThickness: (thickness: number) => void;
  addObject: (obj: ProjectObject) => void;
  updateObject: (id: string, updates: Partial<ProjectObject>) => void;
  removeObject: (id: string) => void;
  updatePlotPoint: (index: number, pos: Position) => void;
  setPlot: (plot: Plot) => void;
  importProject: (data: ProjectData) => void;
}

const initialState: ProjectData & { selectedId: string | null; stageScale: number; stagePos: { x: number; y: number } } = {
  version: "1.1",
  selectedId: null,
  stageScale: 1,
  stagePos: { x: 50, y: 50 },
  projectSettings: {
    module: "shaku",
    gridSizeMm: 910,
    wallThicknessMm: 100,
  },
  plot: {
    points: [
      { x: 0, y: 0 },
      { x: 15000, y: 0 },
      { x: 15000, y: 12000 },
      { x: 0, y: 12000 }
    ],
  },
  objects: [],
};

export const useProjectStore = create<ProjectState>((set) => ({
  ...initialState,

  setSelectedId: (id) => set({ selectedId: id }),

  setStageScale: (scale) => set({ stageScale: scale }),

  setStagePos: (pos) => set({ stagePos: pos }),

  setModule: (module) => set((state) => ({
    projectSettings: {
      ...state.projectSettings,
      module,
      gridSizeMm: module === 'shaku' ? 910 : 1000,
    }
  })),

  setWallThickness: (thickness) => set((state) => ({
    projectSettings: {
      ...state.projectSettings,
      wallThicknessMm: thickness,
    }
  })),

  addObject: (obj) => set((state) => ({
    objects: [...state.objects, obj]
  })),

  updateObject: (id, updates) => set((state) => ({
    objects: state.objects.map((obj) => obj.id === id ? { ...obj, ...updates } : obj)
  })),

  removeObject: (id) => set((state) => ({
    objects: state.objects.filter((obj) => obj.id !== id)
  })),

  updatePlotPoint: (index, pos) => set((state) => {
    const newPoints = [...state.plot.points];
    newPoints[index] = pos;
    return { plot: { points: newPoints } };
  }),

  setPlot: (plot) => set({ plot }),

  importProject: (data) => set({ ...data }),
}));
