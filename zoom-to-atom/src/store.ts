import { create } from 'zustand'

interface ZoomState {
  zoomLevel: number
  setZoomLevel: (level: number) => void
  cameraPosition: [number, number, number]
  setCameraPosition: (pos: [number, number, number]) => void
}

export const useZoomStore = create<ZoomState>((set) => ({
  zoomLevel: 1, // 初期ズーム（人間スケールを1とする）
  setZoomLevel: (level) => set({ zoomLevel: level }),
  cameraPosition: [0, 0, 5],
  setCameraPosition: (pos) => set({ cameraPosition: pos }),
}))
