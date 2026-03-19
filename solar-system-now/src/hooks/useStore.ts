import { create } from 'zustand';

interface SolarSystemState {
  currentDate: Date;
  timeSpeed: number; // multiplier for time passage (e.g. 1 second real time = 1 day simulation time)
  isPaused: boolean;
  selectedPlanet: string | null;
  
  // Actions
  setCurrentDate: (date: Date) => void;
  setTimeSpeed: (speed: number) => void;
  setIsPaused: (paused: boolean) => void;
  setSelectedPlanet: (name: string | null) => void;
  advanceTime: (deltaTimeSeconds: number) => void;
}

export const useStore = create<SolarSystemState>((set) => ({
  currentDate: new Date(),
  timeSpeed: 1, // 1 day per second
  isPaused: false,
  selectedPlanet: null,

  setCurrentDate: (date) => set({ currentDate: date }),
  setTimeSpeed: (speed) => set({ timeSpeed: speed }),
  setIsPaused: (paused) => set({ isPaused: paused }),
  setSelectedPlanet: (name) => set({ selectedPlanet: name }),
  
  advanceTime: (deltaTimeSeconds) => set((state) => {
    if (state.isPaused) return state;
    const newDate = new Date(state.currentDate.getTime() + deltaTimeSeconds * state.timeSpeed * 24 * 60 * 60 * 1000);
    return { currentDate: newDate };
  }),
}));
