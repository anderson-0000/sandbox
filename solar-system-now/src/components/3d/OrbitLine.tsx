import React, { useMemo } from 'react';
import { Line } from '@react-three/drei';
import { type OrbitalElements, getPlanetPosition, SUN_GALACTIC_VELOCITY } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';

interface OrbitLineProps {
  data: OrbitalElements;
  isMoon?: boolean;
}

const OrbitLine: React.FC<OrbitLineProps> = ({ data, isMoon = false }) => {
  const currentDate = useStore((state) => state.currentDate);
  
  const points = useMemo(() => {
    const pts = [];
    const segments = isMoon ? 128 : 512; // More segments for planets to show the long spiral
    
    // We show the path over one orbital period
    const orbitalPeriodDays = 365.25 * Math.pow(data.a, 1.5) * (isMoon ? 580 : 1);
    
    // To make it look like a trail, we sample from past to future
    // sampling range: half period in past to half period in future
    const startOffset = -orbitalPeriodDays * 0.5;
    const step = orbitalPeriodDays / segments;

    for (let i = 0; i <= segments; i++) {
      const offsetDays = startOffset + (i * step);
      const sampleDate = new Date(currentDate.getTime() + offsetDays * 24 * 60 * 60 * 1000);
      
      // Relative position to the sun at that specific time
      const relPos = getPlanetPosition(data, sampleDate, isMoon);
      
      // Galactic displacement: how far the sun moves between 'currentDate' and 'sampleDate'
      // displacement = velocity * delta_t
      const displacement = SUN_GALACTIC_VELOCITY.clone().multiplyScalar(offsetDays);
      
      pts.push(relPos.add(displacement));
    }
    return pts;
  }, [data, isMoon, currentDate]); // Re-calculate when currentDate changes to "animate" the spiral

  return (
    <Line
      points={points}
      color={data.color}
      lineWidth={isMoon ? 0.3 : 0.8}
      transparent
      opacity={isMoon ? 0.1 : 0.25}
    />
  );
};

export default OrbitLine;