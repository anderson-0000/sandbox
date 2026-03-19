import React from 'react';
import { Html } from '@react-three/drei';
import { useStore } from '../../hooks/useStore';
import { SUN_RADIUS } from '../../engine/kepler';

const Sun: React.FC = () => {
  const setSelectedPlanet = useStore((state) => state.setSelectedPlanet);

  return (
    <mesh onClick={() => setSelectedPlanet('Sun')}>
      <sphereGeometry args={[SUN_RADIUS, 64, 64]} />
      <meshStandardMaterial
        emissive="#FDB813"
        emissiveIntensity={2}
        color="#FDB813"
      />
      <pointLight 
        intensity={1000000} 
        decay={2} 
        distance={100000}
        color="#FDB813" 
        castShadow 
        shadow-mapSize={[2048, 2048]}
      />
      
      <Html distanceFactor={150}>
        <div 
          className="text-white text-xs font-bold whitespace-nowrap bg-black/80 px-2 py-0.5 rounded border border-white/20 pointer-events-none select-none shadow-lg"
          style={{ color: '#ffffff' }}
        >
          太陽
        </div>
      </Html>
    </mesh>
  );
};

export default Sun;
