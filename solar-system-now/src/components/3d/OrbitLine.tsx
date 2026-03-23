import React, { useMemo } from 'react';
import { Line } from '@react-three/drei';
import { type OrbitalElements, getPlanetPosition } from '../../engine/kepler';
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
      
      pts.push(relPos);
    }
    return pts;
  }, [data, isMoon, currentDate]); // Re-calculate when currentDate changes to "animate" the spiral

  return (
    <group>
      <Line
        points={points.map(p => [p.x, p.y, p.z])}
        color={data.color}
        lineWidth={isMoon ? 1.5 : 3.0} // Increased thickness
        transparent
        opacity={isMoon ? 0.4 : 0.8} // Increased opacity
      />
    </group>
  );
};

export default OrbitLine;