import { useRef } from 'react'
import { useFrame } from '@react-three/fiber'
import { Mesh } from 'three'

export default function Earth() {
  const earthRef = useRef<Mesh>(null!)

  useFrame(() => {
    if (earthRef.current) {
      earthRef.current.rotation.y += 0.0005
    }
  })

  return (
    <mesh ref={earthRef}>
      <sphereGeometry args={[6371, 64, 64]} />
      <meshStandardMaterial color="#2233ff" wireframe />
    </mesh>
  )
}
