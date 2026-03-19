import React, { useMemo } from 'react';
import { Line } from '@react-three/drei';
import { type OrbitalElements, getPlanetPosition } from '../../engine/kepler';

interface OrbitLineProps {
  data: OrbitalElements;
}

const OrbitLine: React.FC<OrbitLineProps> = ({ data }) => {
  const points = useMemo(() => {
    const pts = [];
    const segments = 128;
    const baseDate = new Date('2000-01-01T12:00:00Z');
    
    // We sample positions over one orbital period (roughly)
    // For Neptune this is 165 years, for Mercury 88 days.
    // To keep it simple and visually correct for an ellipse:
    // We can use the getPlanetPosition with varying mean longitude.
    // Actually, getPlanetPosition uses date, so we'll just step through a year or more.
    
    const orbitalPeriodDays = 365.25 * Math.pow(data.a, 1.5);
    const step = orbitalPeriodDays / segments;

    for (let i = 0; i <= segments; i++) {
      const d = new Date(baseDate.getTime() + i * step * 24 * 60 * 60 * 1000);
      const pos = getPlanetPosition(data, d);
      pts.push(pos);
    }
    return pts;
  }, [data]);

  return (
    <Line
      points={points}
      color={data.color}
      lineWidth={0.5}
      transparent
      opacity={0.3}
    />
  );
};

export default OrbitLine;
