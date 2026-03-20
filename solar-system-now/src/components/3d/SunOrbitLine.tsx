import React, { useMemo } from 'react';
import { Line } from '@react-three/drei';
import { SUN_GALACTIC_VELOCITY } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';

const SunOrbitLine: React.FC = () => {
  const points = useMemo(() => {
    const pts = [];
    // A long line along the velocity vector
    const dir = SUN_GALACTIC_VELOCITY.clone().normalize();
    
    // Line spanning from -1 million to +1 million units relative to current sun position
    // But since the group moves with the sun, we just need a line through origin in group space
    pts.push(dir.clone().multiplyScalar(-500000));
    pts.push(dir.clone().multiplyScalar(500000));
    
    return pts;
  }, []);

  return (
    <Line
      points={points}
      color="#FDB813"
      lineWidth={1}
      transparent
      opacity={0.2}
      dashed
      dashScale={50}
      dashSize={10}
      gapSize={5}
    />
  );
};

export default SunOrbitLine;