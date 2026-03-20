import React, { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import { Html } from '@react-three/drei';
import * as THREE from 'three';
import { type OrbitalElements, getPlanetPosition, MOONS_DATA, getJulianDate } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';
import Moon from './Moon';
import OrbitLine from './OrbitLine';

interface PlanetProps {
  data: OrbitalElements;
}

const Planet: React.FC<PlanetProps> = ({ data }) => {
  const meshRef = useRef<THREE.Mesh>(null);
  const groupRef = useRef<THREE.Group>(null);
  const currentDate = useStore((state) => state.currentDate);
  const setSelectedObjectName = useStore((state) => state.setSelectedObjectName);
  const selectedObjectName = useStore((state) => state.selectedObjectName);
  const showLabels = useStore((state) => state.showLabels);
  const showMoonOrbits = useStore((state) => state.showMoonOrbits);

  // Update position based on time
  useFrame(() => {
    if (groupRef.current) {
      const pos = getPlanetPosition(data, currentDate);
      groupRef.current.position.copy(pos);
    }
    
    // Update rotation
    if (meshRef.current && data.rotationPeriod) {
      const rotationSpeed = (2 * Math.PI) / data.rotationPeriod;
      const jd = getJulianDate(currentDate);
      meshRef.current.rotation.y = rotationSpeed * (jd - 2451545.0);
    }
  });

  const isSelected = selectedObjectName === data.id;
  const moons = MOONS_DATA[data.id] || [];

  return (
    <group ref={groupRef} name={data.id}>
      {/* Render moon orbits relative to planet */}
      {showMoonOrbits && moons.map((moon) => (
        <OrbitLine key={`${moon.id}-orbit`} data={moon} isMoon />
      ))}

      <mesh
        ref={meshRef}
        userData={{ radius: data.radius }}
        castShadow
        receiveShadow
        rotation={[data.axialTilt ? data.axialTilt * (Math.PI / 180) : 0, 0, 0]}
        onClick={(e) => {
          e.stopPropagation();
          setSelectedObjectName(data.id);
        }}
      >
        <sphereGeometry args={[data.radius, 32, 32]} />
        <meshStandardMaterial color={data.color} roughness={0.7} metalness={0.1} />
        
        {isSelected && (
          <mesh scale={[1.2, 1.2, 1.2]} rotation={[0, 0, 0]}>
            <sphereGeometry args={[data.radius, 32, 32]} />
            <meshBasicMaterial color="white" wireframe transparent opacity={0.3} />
          </mesh>
        )}

        {showLabels && (
          <Html distanceFactor={25}>
            <div 
              className="text-white text-xs font-bold whitespace-nowrap bg-black/80 px-2 py-0.5 rounded border border-white/20 pointer-events-none select-none shadow-lg"
              style={{ color: '#ffffff' }}
            >
              {data.name}
            </div>
          </Html>
        )}
      </mesh>

      {/* Render moons */}
      {moons.map((moon) => (
        <Moon key={moon.name} data={moon} />
      ))}
    </group>
  );
};

export default Planet;