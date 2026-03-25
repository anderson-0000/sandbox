import React, { useRef } from 'react'
import { useFrame } from '@react-three/fiber'
import * as THREE from 'three'

interface ScaleContainerProps {
  children: React.ReactNode
  range: [number, number] // 可視範囲 [minZoom, maxZoom]
  currentZoom: number
}

/**
 * ズームレベルに応じて中身の透明度を制御するコンポーネント。
 */
export const ScaleContainer: React.FC<ScaleContainerProps> = ({ children, range, currentZoom }) => {
  const groupRef = useRef<THREE.Group>(null!)
  const [min, max] = range

  // ズームに応じた透明度の計算
  // min に近い時は 0 -> 1 に、中央で 1、max に近づくと 1 -> 0 に。
  const calculateOpacity = (zoom: number) => {
    if (zoom < min || zoom > max) return 0
    const center = (min + max) / 2
    const halfWidth = (max - min) / 2
    const distance = Math.abs(zoom - center)
    return Math.max(0, 1 - distance / halfWidth)
  }

  useFrame(() => {
    const opacity = calculateOpacity(currentZoom)
    groupRef.current.visible = opacity > 0.01

    // 子要素のマテリアルの透明度を一括操作する（簡易的な実装）
    groupRef.current.traverse((obj) => {
      if ((obj as THREE.Mesh).isMesh) {
        const material = (obj as THREE.Mesh).material as THREE.Material
        if (material.transparent) {
          material.opacity = opacity
        }
      }
    })
  })

  return <group ref={groupRef}>{children}</group>
}
