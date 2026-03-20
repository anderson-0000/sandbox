import React from 'react';
import { Html } from '@react-three/drei';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';
import { useStore } from '../../hooks/useStore';
import { SUN_RADIUS } from '../../engine/kepler';

const Sun: React.FC = () => {
  const meshRef = React.useRef<THREE.Mesh>(null);
  const setSelectedObjectName = useStore((state) => state.setSelectedObjectName);
  const sunPosition = useStore((state) => state.sunPosition);
  const showLabels = useStore((state) => state.showLabels);

  useFrame(() => {
    if (meshRef.current) {
      meshRef.current.position.copy(sunPosition);
    }
  });

  return (
    <mesh 
      ref={meshRef}
      name="Sun" 
      userData={{ radius: SUN_RADIUS }}
      onClick={() => setSelectedObjectName('Sun')}
    >
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
      
      {showLabels && (
        <Html distanceFactor={150}>
          <div 
            className="text-white text-xs font-bold whitespace-nowrap bg-black/80 px-2 py-0.5 rounded border border-white/20 pointer-events-none select-none shadow-lg"
            style={{ color: '#ffffff' }}
          >
            太陽
          </div>
        </Html>
      )}
    </mesh>
  );
};

export default Sun;