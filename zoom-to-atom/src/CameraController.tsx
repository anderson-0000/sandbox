import { useThree, useFrame } from '@react-three/fiber'
import { useEffect, useRef } from 'react'
import { useZoomStore } from './store'
import * as THREE from 'three'

export const CameraController = () => {
  const { camera, gl } = useThree()
  const zoomLevel = useZoomStore((state) => state.zoomLevel)
  const setZoomLevel = useZoomStore((state) => state.setZoomLevel)
  
  const targetZoom = useRef(zoomLevel)
  const lastTouchDistance = useRef<number | null>(null)

  useEffect(() => {
    // --- マウスホイール / トラックパッド スクロール ---
    const handleWheel = (e: WheelEvent) => {
      e.preventDefault()
      // ズーム感度の調整
      // 上にスクロール / ピンチアウト (deltaY < 0) でズームイン (targetZoom 増加)
      const sensitivity = 0.005
      targetZoom.current -= e.deltaY * sensitivity
      
      // 範囲制限 (0: 人間, 14.0: 原子核の深淵)
      targetZoom.current = Math.min(Math.max(targetZoom.current, 0), 14.0)
    }

    // --- タッチ操作 (ピンチイン/アウト) ---
    const handleTouchMove = (e: TouchEvent) => {
      if (e.touches.length === 2) {
        e.preventDefault()
        const dx = e.touches[0].pageX - e.touches[1].pageX
        const dy = e.touches[0].pageY - e.touches[1].pageY
        const distance = Math.sqrt(dx * dx + dy * dy)

        if (lastTouchDistance.current !== null) {
          const delta = distance - lastTouchDistance.current
          // ピンチアウト (delta > 0) でズームイン (targetZoom 増加)
          const touchSensitivity = 0.01
          targetZoom.current += delta * touchSensitivity
          targetZoom.current = Math.min(Math.max(targetZoom.current, 0), 14.0)
        }
        lastTouchDistance.current = distance
      }
    }

    const handleTouchEnd = () => {
      lastTouchDistance.current = null
    }

    const domElement = gl.domElement
    domElement.addEventListener('wheel', handleWheel, { passive: false })
    domElement.addEventListener('touchmove', handleTouchMove, { passive: false })
    domElement.addEventListener('touchend', handleTouchEnd)
    
    return () => {
      domElement.removeEventListener('wheel', handleWheel)
      domElement.removeEventListener('touchmove', handleTouchMove)
      domElement.removeEventListener('touchend', handleTouchEnd)
    }
  }, [gl])

  useFrame(() => {
    // スムーズなズーム遷移
    const newZoom = THREE.MathUtils.lerp(zoomLevel, targetZoom.current, 0.1)
    setZoomLevel(newZoom)

    // ズームレベルに応じたカメラ位置の計算
    // z = 5 * (0.1 ^ zoomLevel)
    const zPos = 5 * Math.pow(0.1, newZoom)
    camera.position.z = zPos
    
    // Near/Far の調整（極小スケールに対応するため、より小さい値を設定）
    camera.near = Math.min(0.1, zPos * 0.001)
    camera.far = 100
    camera.updateProjectionMatrix()
  })

  return null
}
