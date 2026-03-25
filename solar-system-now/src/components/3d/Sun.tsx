import React, { useRef } from 'react';
import { Html } from '@react-three/drei';
import * as THREE from 'three';
import { useStore } from '../../hooks/useStore';
import { SUN_RADIUS, SUN_GALACTIC_VELOCITY } from '../../engine/kepler';
import MovementVector from './MovementVector';

const Sun: React.FC = () => {
  const meshRef = useRef<THREE.Mesh>(null);
  const setSelectedObjectName = useStore((state) => state.setSelectedObjectName);
  const showLabels = useStore((state) => state.showLabels);
  const orreryMode = useStore((state) => state.orreryMode);
  const viewMode = useStore((state) => state.viewMode);

  // Sync with visual scale of other bodies (100x)
  const visualScale = (orreryMode || viewMode === 'solar_system' || viewMode === 'galactic') ? 100 : 1.0;
  const radius = SUN_RADIUS * visualScale;

  return (
    <group>
      <mesh 
        ref={meshRef}
        name="Sun" 
        userData={{ radius: radius }}
        onClick={(e) => {
          e.stopPropagation();
          setSelectedObjectName('Sun');
        }}
      >
        <MovementVector velocity={SUN_GALACTIC_VELOCITY} color="#FDB813" scale={10} />
        <sphereGeometry args={[radius, 64, 64]} />
        <meshBasicMaterial color="#FDB813" />
        <pointLight 
          intensity={10000000} 
          decay={1.5} 
          distance={0}
          color="#FDB813" 
          castShadow 
          shadow-mapSize={[2048, 2048]}
        />
        
        {/* Glow effect */}
        <mesh scale={[1.2, 1.2, 1.2]}>
          <sphereGeometry args={[radius, 32, 32]} />
          <meshBasicMaterial color="#FDB813" transparent opacity={0.3} />
        </mesh>

        {showLabels && (
          <Html position={[0, radius * 1.5, 0]}>
            <div 
              className="text-white text-xs font-bold whitespace-nowrap bg-black/80 px-2 py-0.5 rounded border border-white/20 pointer-events-none select-none shadow-lg"
              style={{ color: '#ffffff' }}
            >
              太陽
            </div>
          </Html>
        )}
      </mesh>
    </group>
  );
};

export default Sun;
