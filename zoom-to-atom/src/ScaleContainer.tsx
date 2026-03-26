import React, { useRef } from 'react'
import { useFrame } from '@react-three/fiber'
import * as THREE from 'three'

interface ScaleContainerProps {
  children: React.ReactNode
  range: [number, number] // 可視範囲 [minZoom, maxZoom]
  currentZoom: number
  baseScale?: number // そのスケールが「等倍」で見えるべき中心のzoomLevel
}

/**
 * ズームレベルに応じて中身の透明度と「スケール」を制御するコンポーネント。
 */
export const ScaleContainer: React.FC<ScaleContainerProps> = ({ children, range, currentZoom, baseScale = 0 }) => {
  const groupRef = useRef<THREE.Group>(null!)
  const [min, max] = range

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

    // スケール計算: zoomLevel が 1 上がるごとに 10倍大きくする
    // baseScale において 1倍 (10^0) になるようにオフセット
    const s = Math.pow(10, currentZoom - baseScale)
    groupRef.current.scale.set(s, s, s)

    // 子要素のマテリアルの透明度を一括操作する
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
