import { Canvas, useFrame, useThree } from '@react-three/fiber'
import { OrbitControls, Stars } from '@react-three/drei'
import Earth from './Earth'
import Moon from './Moon'
import SatelliteCloud from './SatelliteCloud'
import { type SatelliteData } from '../services/api'

interface OrbitSceneProps {
  onSelect: (sat: SatelliteData | null, alt: number) => void;
  selectedSat: SatelliteData | null;
  onAltitudeChange: (alt: number) => void;
  onDataLoaded: (count: number) => void;
}

function CameraTracker({ onAltitudeChange }: { onAltitudeChange: (alt: number) => void }) {
  const { camera } = useThree()
  useFrame(() => {
    const distance = camera.position.length()
    const altitude = Math.max(0, distance - 6371)
    onAltitudeChange(altitude)
  })
  return null
}

export default function OrbitScene({ onSelect, selectedSat, onAltitudeChange, onDataLoaded }: OrbitSceneProps) {
  return (
    <div style={{ width: '100%', height: '100%', position: 'relative', zIndex: 1 }}>
      <Canvas 
        camera={{ position: [0, 20000, 40000], far: 2000000, near: 10, fov: 45 }}
        raycaster={{
          params: {
            Mesh: { threshold: 10 }
          }
        }}
        onPointerMissed={(e: any) => {
          if (e.target.tagName === 'CANVAS') {
            onSelect(null, 0)
          }
        }}
      >
        <color attach="background" args={['#050505']} />
        <ambientLight intensity={0.2} />
        <pointLight position={[100000, 50000, 50000]} intensity={2.0} />
        <Stars radius={500000} depth={100000} count={10000} factor={4} saturation={0} fade speed={0} />
        <OrbitControls minDistance={6500} maxDistance={1000000} makeDefault />
        
        <group raycast={() => null}>
          <Earth />
          <Moon />
        </group>
        
        <SatelliteCloud 
          onSelect={onSelect} 
          selectedSat={selectedSat} 
          onDataLoaded={onDataLoaded}
        />
        
        <CameraTracker onAltitudeChange={onAltitudeChange} />
        
        <gridHelper args={[800000, 40, 0x222222, 0x111111]} position={[0, -5000, 0]} />
      </Canvas>
    </div>
  )
}
