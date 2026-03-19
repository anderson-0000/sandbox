import React, { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import { Html } from '@react-three/drei';
import * as THREE from 'three';
import { type MoonOrbitalElements, getPlanetPosition } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';

interface MoonProps {
  data: MoonOrbitalElements;
}

const Moon: React.FC<MoonProps> = ({ data }) => {
  const meshRef = useRef<THREE.Mesh>(null);
  const currentDate = useStore((state) => state.currentDate);
  const setSelectedPlanet = useStore((state) => state.setSelectedPlanet);

  useFrame(() => {
    if (meshRef.current) {
      const pos = getPlanetPosition(data, currentDate, true);
      meshRef.current.position.copy(pos);
    }
  });

  return (
    <mesh 
      ref={meshRef} 
      castShadow 
      receiveShadow
      onClick={(e) => {
        e.stopPropagation();
        setSelectedPlanet(data.id);
      }}
    >
      <sphereGeometry args={[data.radius, 16, 16]} />
      <meshStandardMaterial 
        color={data.color} 
        emissive={data.color}
        emissiveIntensity={0.2}
        roughness={0.8} 
        metalness={0.1} 
      />
      
      <Html distanceFactor={15}>
        <div className="text-white text-[10px] font-bold whitespace-nowrap bg-black/40 px-1 rounded pointer-events-none select-none">
          {data.name}
        </div>
      </Html>
    </mesh>
  );
};

export default Moon;
