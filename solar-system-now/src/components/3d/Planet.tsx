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

  // Update position based on time
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

  // Determine a visual scale factor so planets are actually visible in wide views
  // In solar_system view, planets are too small to see at true scale
  const visualScale = (orreryMode || viewMode === 'solar_system') 
    ? (data.id === 'Jupiter' || data.id === 'Saturn' ? 15 : 40) // Scale factor adjustment
    : 1.0;

  const planetRadius = data.radius * visualScale;

  return (
    <group ref={groupRef} name={data.id} userData={{ radius: planetRadius }}>
      <MovementVector velocity={velocity} color={data.color} scale={50} />
      {/* Render moon orbits relative to planet */}
      {showMoonOrbits && moons.map((moon) => (
        <OrbitLine key={`${moon.id}-orbit`} data={moon} isMoon />
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

      {/* Render moons */}
      {moons.map((moon) => (
        <Moon key={moon.name} data={moon} />
      ))}
    </group>
  );
};

export default Planet;