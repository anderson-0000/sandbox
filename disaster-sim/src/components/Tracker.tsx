import { useFrame } from '@react-three/fiber';
import { useStore } from '../store/useStore';

/**
 * R3F 内で Zustand の状態を同期するためのコンポーネント。
 * useFrame を使用してシミュレーション時間を更新し、
 * 必要に応じてグローバルな状態（カメラ座標等）をストアに流す。
 */
export const Tracker = () => {
  const updateTime = useStore((state) => state.updateTime);
  const isActive = useStore((state) => state.isActive);

  useFrame((_state, delta) => {
    if (isActive) {
      // delta は秒単位。シミュレーターの精度に応じてスケール調整可能。
      updateTime(delta);
    }
  });

  return null;
};
