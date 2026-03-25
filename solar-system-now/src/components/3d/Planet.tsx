import React, { useRef, useState } from 'react';
import { useFrame } from '@react-three/fiber';
import { Html } from '@react-three/drei';
import * as THREE from 'three';
import { type OrbitalElements, getPlanetPosition, getPlanetVelocity, MOONS_DATA, getJulianDate } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';
import Moon from './Moon';
import OrbitLine from './OrbitLine';
import MovementVector from './MovementVector';

interface PlanetProps {
  data: OrbitalElements;
}

const Planet: React.FC<PlanetProps> = ({ data }) => {
  const meshRef = useRef<THREE.Mesh>(null);
  const groupRef = useRef<THREE.Group>(null);
  const [velocity, setVelocity] = useState(new THREE.Vector3());
  const currentDate = useStore((state) => state.currentDate);
  const setSelectedObjectName = useStore((state) => state.setSelectedObjectName);
  const selectedObjectName = useStore((state) => state.selectedObjectName);
  const showLabels = useStore((state) => state.showLabels);
  const showMoonOrbits = useStore((state) => state.showMoonOrbits);
  const orreryMode = useStore((state) => state.orreryMode);
  const viewMode = useStore((state) => state.viewMode);
  const zoomDistance = useStore((state) => state.zoomDistance);

  // Use a consistent scale to make planets visible
  // To keep Moon-Planet ratio correct, we MUST apply this scale to Moon DISTANCE as well.
  const visualScale = (orreryMode || viewMode === 'solar_system' || viewMode === 'galactic') ? 100 : 1.0;

  useFrame(() => {
    if (groupRef.current) {
      const pos = getPlanetPosition(data, currentDate);
      if (!isNaN(pos.x) && !isNaN(pos.y) && !isNaN(pos.z)) {
        groupRef.current.position.copy(pos);
      }

      const vel = getPlanetVelocity(data, currentDate);
      if (!isNaN(vel.x) && !isNaN(vel.y) && !isNaN(vel.z)) {
        setVelocity(vel);
      }
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
  const planetRadius = data.radius * visualScale;

  return (
    <group ref={groupRef} name={data.id} userData={{ radius: planetRadius }}>
      <MovementVector velocity={velocity} color={data.color} scale={50} />
      
      {/* 
        IMPORTANT: We pass visualScale to children so they can scale their distance correctly.
        This keeps the Earth-Moon visual ratio accurate even when they are enlarged.
      */}
      {showMoonOrbits && moons.map((moon) => (
        <OrbitLine key={`${moon.id}-orbit`} data={moon} isMoon parentScale={visualScale} />
      ))}

      <mesh
        ref={meshRef}
        userData={{ radius: planetRadius }}
        rotation={[data.axialTilt ? data.axialTilt * (Math.PI / 180) : 0, 0, 0]}
        onClick={(e) => {
          e.stopPropagation();
          setSelectedObjectName(data.id);
        }}
      >
        <sphereGeometry args={[planetRadius || 1, 32, 32]} />
        <meshBasicMaterial 
          color={data.color} 
        />
        
        {isSelected && (
          <mesh scale={[1.3, 1.3, 1.3]} rotation={[0, 0, 0]}>
            <sphereGeometry args={[planetRadius || 1, 32, 32]} />
            <meshBasicMaterial color="white" wireframe transparent opacity={0.6} />
          </mesh>
        )}
      </mesh>

      {showLabels && (
        <Html distanceFactor={zoomDistance * 2} position={[0, planetRadius * 1.5, 0]}>
          <div 
            className="text-white text-xs font-bold whitespace-nowrap bg-black/80 px-2 py-0.5 rounded border border-white/20 pointer-events-none select-none shadow-lg"
            style={{ color: '#ffffff' }}
          >
            {data.name}
          </div>
        </Html>
      )}

      {/* Render moons with the same scale factor for their distance */}
      {moons.map((moon) => (
        <Moon key={moon.name} data={moon} parentScale={visualScale} />
      ))}
    </group>
  );
};

export default Planet;
