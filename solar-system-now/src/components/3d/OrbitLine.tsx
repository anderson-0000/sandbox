import React, { useMemo } from 'react';
import * as THREE from 'three';
import { Line } from '@react-three/drei';
import { type OrbitalElements, getPlanetPosition, SUN_GALACTIC_VELOCITY } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';

interface OrbitLineProps {
  data: OrbitalElements;
  isMoon?: boolean;
  parentScale?: number;
}

const OrbitLine: React.FC<OrbitLineProps> = ({ data, isMoon = false, parentScale = 1.0 }) => {
  const currentDate = useStore((state) => state.currentDate);
  const sunPosition = useStore((state) => state.sunPosition);
  
  const points = useMemo(() => {
    const pts: THREE.Vector3[] = [];
    const segments = isMoon ? 128 : 1024;
    
    // Orbital period in days
    const orbitalPeriodDays = 365.25 * Math.pow(data.a, 1.5) * (isMoon ? 27.3 : 1); 
    
    // Range for trail
    const startOffset = -orbitalPeriodDays;
    const step = orbitalPeriodDays / segments;

    for (let i = 0; i <= segments; i++) {
      const offsetDays = startOffset + (i * step);
      const sampleDate = new Date(currentDate.getTime() + offsetDays * 24 * 60 * 60 * 1000);
      
      const relPos = getPlanetPosition(data, sampleDate, isMoon);
      
      if (isMoon) {
        // Moon trails are relative to planet group in Planet.tsx
        pts.push(relPos.multiplyScalar(parentScale));
      } else {
        // Planet trails are absolute in galactic space
        // Estimated sun position at sample time relative to current sunPosition
        const sunDisplacement = SUN_GALACTIC_VELOCITY.clone().multiplyScalar(offsetDays);
        const absPos = sunPosition.clone().add(sunDisplacement).add(relPos);
        pts.push(absPos);
      }
    }
    return pts;
  }, [data, isMoon, currentDate, sunPosition, parentScale]); 

  return (
    <group>
      <Line
        points={points.map(p => [p.x, p.y, p.z])}
        color={data.color}
        lineWidth={isMoon ? 1.0 : 2.5}
        transparent
        opacity={isMoon ? 0.3 : 0.6}
      />
    </group>
  );
};

export default OrbitLine;
