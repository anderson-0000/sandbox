import React, { useRef, useMemo } from 'react';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';
import { useStore } from '../../hooks/useStore';

const SpaceDust: React.FC = () => {
  const count = 1500;
  const meshRef = useRef<THREE.Points>(null);
  const sunPosition = useStore((state) => state.sunPosition);

  // Create a large box of random points
  const [positions, sizes] = useMemo(() => {
    const pos = new Float32Array(count * 3);
    const sz = new Float32Array(count);
    for (let i = 0; i < count; i++) {
      pos[i * 3] = (Math.random() - 0.5) * 50000;
      pos[i * 3 + 1] = (Math.random() - 0.5) * 50000;
      pos[i * 3 + 2] = (Math.random() - 0.5) * 50000;
      sz[i] = Math.random() * 2;
    }
    return [pos, sz];
  }, [count]);

  useFrame(() => {
    if (meshRef.current) {
      // Keep the dust field centered around the sun's position
      meshRef.current.position.set(
        Math.floor(sunPosition.x / 1000) * 1000,
        Math.floor(sunPosition.y / 1000) * 1000,
        Math.floor(sunPosition.z / 1000) * 1000
      );
    }
  });

  return (
    <points ref={meshRef}>
      <bufferGeometry>
        <bufferAttribute
          attach="attributes-position"
          args={[positions, 3]}
        />
        <bufferAttribute
          attach="attributes-size"
          args={[sizes, 1]}
        />
      </bufferGeometry>
      <pointsMaterial 
        size={5} 
        sizeAttenuation 
        color="#ffffff" 
        transparent 
        opacity={0.4} 
        blending={THREE.AdditiveBlending}
      />
    </points>
  );
};

export default SpaceDust;