import React, { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import { Html } from '@react-three/drei';
import * as THREE from 'three';
import { type MoonOrbitalElements, getPlanetPosition, getJulianDate } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';

interface MoonProps {
  data: MoonOrbitalElements;
}

const Moon: React.FC<MoonProps> = ({ data }) => {
  const meshRef = useRef<THREE.Mesh>(null);
  const currentDate = useStore((state) => state.currentDate);
  const setSelectedObjectName = useStore((state) => state.setSelectedObjectName);
  const showLabels = useStore((state) => state.showLabels);
  const orreryMode = useStore((state) => state.orreryMode);
  const viewMode = useStore((state) => state.viewMode);

  useFrame(() => {
    if (meshRef.current) {
      const pos = getPlanetPosition(data, currentDate, true);
      if (!isNaN(pos.x) && !isNaN(pos.y) && !isNaN(pos.z)) {
        meshRef.current.position.copy(pos);
      }
      
      // Update rotation
      if (data.rotationPeriod) {
        const rotationSpeed = (2 * Math.PI) / data.rotationPeriod;
        const jd = getJulianDate(currentDate);
        meshRef.current.rotation.y = rotationSpeed * (jd - 2451545.0);
      }
    }
  });

  const visualScale = (orreryMode || viewMode === 'solar_system') ? 25 : 1.0;
  const radius = data.radius * visualScale;

  return (
    <group>
      <mesh 
        ref={meshRef}
        name={data.id}
        userData={{ radius: radius }}
        castShadow 
        receiveShadow
        rotation={[data.axialTilt ? data.axialTilt * (Math.PI / 180) : 0, 0, 0]}
        onClick={(e) => {
          e.stopPropagation();
          setSelectedObjectName(data.id);
        }}
      >
        <sphereGeometry args={[radius, 16, 16]} />
        <meshStandardMaterial 
          color={data.color} 
          emissive={data.color}
          emissiveIntensity={0.8}
          roughness={0.8} 
          metalness={0.1} 
        />
        
        {showLabels && (
          <Html position={[0, radius * 1.5, 0]}>
            <div className="text-white text-[10px] font-bold whitespace-nowrap bg-black/40 px-1 rounded pointer-events-none select-none">
              {data.name}
            </div>
          </Html>
        )}
      </mesh>
    </group>
  );
};

export default Moon;
