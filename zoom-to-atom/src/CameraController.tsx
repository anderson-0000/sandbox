import { useThree, useFrame } from '@react-three/fiber'
import { useEffect, useRef } from 'react'
import { useZoomStore } from './store'
import * as THREE from 'three'

export const CameraController = () => {
  const { gl } = useThree()
  const zoomLevel = useZoomStore((state) => state.zoomLevel)
  const setZoomLevel = useZoomStore((state) => state.setZoomLevel)
  
  const targetZoom = useRef(zoomLevel)
  const lastTouchDistance = useRef<number | null>(null)

  useEffect(() => {
    const handleWheel = (e: WheelEvent) => {
      e.preventDefault()
      // 感度調整：スクロールで zoomLevel を増減させる
      const sensitivity = 0.002
      targetZoom.current += e.deltaY * sensitivity
      // 0倍（人間）から 18倍（クォーク）まで
      targetZoom.current = Math.min(Math.max(targetZoom.current, 0), 18.0)
    }

    const handleTouchMove = (e: TouchEvent) => {
      if (e.touches.length === 2) {
        e.preventDefault()
        const dx = e.touches[0].pageX - e.touches[1].pageX
        const dy = e.touches[0].pageY - e.touches[1].pageY
        const distance = Math.sqrt(dx * dx + dy * dy)

        if (lastTouchDistance.current !== null) {
          const delta = distance - lastTouchDistance.current
          const touchSensitivity = 0.01
          targetZoom.current -= delta * touchSensitivity
          targetZoom.current = Math.min(Math.max(targetZoom.current, 0), 18.0)
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
    // スムーズな数値遷移
    const newZoom = THREE.MathUtils.lerp(zoomLevel, targetZoom.current, 0.1)
    setZoomLevel(newZoom)
    
    // カメラ位置は固定（z=5）することで、浮動小数点の精度問題を回避
  })

  return null
}
