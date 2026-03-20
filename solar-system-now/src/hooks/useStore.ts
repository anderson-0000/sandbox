import { create } from 'zustand';
import { Vector3 } from 'three';
import { SUN_GALACTIC_VELOCITY } from '../engine/kepler';

interface SolarSystemState {
  currentDate: Date;
  timeSpeed: number; // multiplier for time passage
  isPaused: boolean;
  selectedObjectName: string | null;
  surfaceTargetName: string | null;
  cameraTarget: Vector3;
  sunPosition: Vector3;
  viewMode: 'orbit' | 'surface' | 'galactic';
  orreryMode: boolean;
  showOrbits: boolean;
  showMoonOrbits: boolean;
  showSunOrbit: boolean;
  showLabels: boolean;
  zoomDistance: number;
  
  // Actions
  setCurrentDate: (date: Date) => void;
  setTimeSpeed: (speed: number) => void;
  setIsPaused: (paused: boolean) => void;
  setSelectedObjectName: (name: string | null) => void;
  setSurfaceTargetName: (name: string | null) => void;
  setCameraTarget: (target: Vector3) => void;
  setViewMode: (mode: 'orbit' | 'surface' | 'galactic') => void;
  setOrreryMode: (active: boolean) => void;
  setShowOrbits: (show: boolean) => void;
  setShowMoonOrbits: (show: boolean) => void;
  setShowSunOrbit: (show: boolean) => void;
  setShowLabels: (show: boolean) => void;
  setZoomDistance: (distance: number) => void;
  advanceTime: (deltaTimeSeconds: number) => void;
}

export const useStore = create<SolarSystemState>((set) => ({
  currentDate: new Date(),
  timeSpeed: 1, // 1 day per second
  isPaused: false,
  selectedObjectName: null,
  surfaceTargetName: null,
  cameraTarget: new Vector3(0, 0, 0),
  sunPosition: new Vector3(0, 0, 0),
  viewMode: 'orbit',
  orreryMode: false,
  showOrbits: true,
  showMoonOrbits: true,
  showSunOrbit: true,
  showLabels: true,
  zoomDistance: 5000,

  setCurrentDate: (date) => set({ currentDate: date }),
  setTimeSpeed: (speed) => set({ timeSpeed: speed }),
  setIsPaused: (paused) => set({ isPaused: paused }),
  setSelectedObjectName: (name) => set((state) => ({ 
    selectedObjectName: name, 
    viewMode: state.viewMode === 'galactic' ? 'galactic' : 'orbit', 
    surfaceTargetName: null 
  })), 
  setSurfaceTargetName: (name) => set({ surfaceTargetName: name }),
  setCameraTarget: (target) => set({ cameraTarget: target }),
  setViewMode: (mode) => set({ viewMode: mode }),
  setOrreryMode: (active) => set({ orreryMode: active }),
  setShowOrbits: (show) => set({ showOrbits: show }),
  setShowMoonOrbits: (show) => set({ showMoonOrbits: show }),
  setShowSunOrbit: (show) => set({ showSunOrbit: show }),
  setShowLabels: (show) => set({ showLabels: show }),
  setZoomDistance: (distance) => set({ zoomDistance: distance }),
  
  advanceTime: (deltaTimeSeconds) => set((state) => {
    if (state.isPaused) return {};
    const deltaDays = deltaTimeSeconds * (state.timeSpeed / 1.0);
    const newDate = new Date(state.currentDate.getTime() + deltaDays * 24 * 60 * 60 * 1000);
    
    const displacement = SUN_GALACTIC_VELOCITY.clone().multiplyScalar(deltaDays);
    const newSunPos = state.sunPosition.clone().add(displacement);
    
    return { currentDate: newDate, sunPosition: newSunPos };
  }),
}));