import { useRef } from 'react'
import { useFrame } from '@react-three/fiber'
import { Mesh } from 'three'

export default function Moon() {
  const moonRef = useRef<Mesh>(null!)
  // Moon distance from Earth: ~384,400 km
  const moonDistance = 384400

  useFrame(({ clock }) => {
    // Very slow orbit
    const t = clock.getElapsedTime() * 0.05
    if (moonRef.current) {
      moonRef.current.position.x = Math.cos(t) * moonDistance
      moonRef.current.position.z = Math.sin(t) * moonDistance
      moonRef.current.rotation.y += 0.001
    }
  })

  return (
    <mesh ref={moonRef} position={[moonDistance, 0, 0]}>
      <sphereGeometry args={[1737, 32, 32]} />
      <meshStandardMaterial color="#cccccc" wireframe />
    </mesh>
  )
}
