import { create } from 'zustand';

interface Epicenter {
  x: number;
  z: number;
  depthKm: number;
  magnitude: number;
  time: number; // 発震時刻
}

interface SimulationState {
  isActive: boolean;
  currentTime: number;
  epicenter: Epicenter | null;
  playbackSpeed: number;
  // Top-level settings to allow adjustment before/after epicenter creation
  magnitude: number;
  depthKm: number;
  
  // Actions
  setEpicenter: (x: number, z: number) => void;
  toggleActive: () => void;
  reset: () => void;
  updateTime: (delta: number) => void;
  setMagnitude: (val: number) => void;
  setDepth: (val: number) => void;
  setPlaybackSpeed: (val: number) => void;
}

export const useStore = create<SimulationState>((set) => ({
  isActive: false,
  currentTime: 0,
  epicenter: null,
  playbackSpeed: 1,
  magnitude: 7.0,
  depthKm: 10,

  setEpicenter: (x, z) => set((state) => ({
    epicenter: {
      x,
      z,
      depthKm: state.depthKm,
      magnitude: state.magnitude,
      time: state.currentTime,
    },
    isActive: true,
  })),

  toggleActive: () => set((state) => ({ isActive: !state.isActive })),
  
  reset: () => set({
    isActive: false,
    currentTime: 0,
    epicenter: null
  }),

  updateTime: (delta) => set((state) => ({
    currentTime: state.isActive ? state.currentTime + delta * state.playbackSpeed : state.currentTime
  })),

  setMagnitude: (val) => set((state) => ({
    magnitude: val,
    epicenter: state.epicenter ? { ...state.epicenter, magnitude: val } : null
  })),

  setDepth: (val) => set((state) => ({
    depthKm: val,
    epicenter: state.epicenter ? { ...state.epicenter, depthKm: val } : null
  })),

  setPlaybackSpeed: (val) => set({ playbackSpeed: val }),
}));
