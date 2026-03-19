import React, { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import { Html } from '@react-three/drei';
import * as THREE from 'three';
import { type OrbitalElements, getPlanetPosition, MOONS_DATA } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';
import Moon from './Moon';

interface PlanetProps {
  data: OrbitalElements;
}

const Planet: React.FC<PlanetProps> = ({ data }) => {
  const meshRef = useRef<THREE.Mesh>(null);
  const groupRef = useRef<THREE.Group>(null);
  const currentDate = useStore((state) => state.currentDate);
  const setSelectedPlanet = useStore((state) => state.setSelectedPlanet);
  const selectedPlanet = useStore((state) => state.selectedPlanet);

  // Update position based on time
  useFrame(() => {
    if (groupRef.current) {
      const pos = getPlanetPosition(data, currentDate);
      groupRef.current.position.copy(pos);
    }
  });

  const isSelected = selectedPlanet === data.id;
  const moons = MOONS_DATA[data.id] || [];

  return (
    <group ref={groupRef}>
      <mesh
        ref={meshRef}
        castShadow
        receiveShadow
        onClick={(e) => {
          e.stopPropagation();
          setSelectedPlanet(data.id);
        }}
      >
        <sphereGeometry args={[data.radius, 32, 32]} />
        <meshStandardMaterial color={data.color} roughness={0.7} metalness={0.1} />
        
        {isSelected && (
          <mesh scale={[1.2, 1.2, 1.2]}>
            <sphereGeometry args={[data.radius, 32, 32]} />
            <meshBasicMaterial color="white" wireframe transparent opacity={0.3} />
          </mesh>
        )}

        <Html distanceFactor={25}>
          <div 
            className="text-white text-xs font-bold whitespace-nowrap bg-black/80 px-2 py-0.5 rounded border border-white/20 pointer-events-none select-none shadow-lg"
            style={{ color: '#ffffff' }}
          >
            {data.name}
          </div>
        </Html>
      </mesh>

      {/* Render moons */}
      {moons.map((moon) => (
        <Moon key={moon.name} data={moon} />
      ))}
    </group>
  );
};

export default Planet;
