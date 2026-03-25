import { useThree, useFrame } from '@react-three/fiber'
import { useEffect } from 'react'
import { useZoomStore } from './store'

/**
 * R3FのCanvas内でのみ呼び出し可能なフックを使用し、
 * 親のZustandストア（Canvas外からでもアクセス可能）にデータを同期させるためのコンポーネント。
 */
export const Tracker = () => {
  const { camera } = useThree()
  const setCameraPosition = useZoomStore((state) => state.setCameraPosition)

  // 毎フレームごとにカメラの位置を同期させる（必要に応じてスロットリングする）
  useFrame(() => {
    setCameraPosition([camera.position.x, camera.position.y, camera.position.z])
  })

  // カメラの設定が変更された際の初期同期
  useEffect(() => {
    setCameraPosition([camera.position.x, camera.position.y, camera.position.z])
  }, [camera, setCameraPosition])

  return null
}
