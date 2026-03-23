import React, { useMemo } from 'react';
import { Line } from '@react-three/drei';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';
import { SUN_GALACTIC_VELOCITY } from '../../engine/kepler';
import { useStore } from '../../hooks/useStore';

const SunOrbitLine: React.FC = () => {
  const lineRef = React.useRef<THREE.Group>(null);
  const sunPosition = useStore((state) => state.sunPosition);

  useFrame(() => {
    if (lineRef.current) {
        lineRef.current.position.copy(sunPosition);
    }
  });

  const points = useMemo(() => {
    const pts = [];
    const dir = SUN_GALACTIC_VELOCITY.clone().normalize();
    pts.push(dir.clone().multiplyScalar(-500000));
    pts.push(dir.clone().multiplyScalar(500000));
    return pts;
  }, []);

  return (
    <group ref={lineRef}>
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
    </group>
  );
};

export default SunOrbitLine;