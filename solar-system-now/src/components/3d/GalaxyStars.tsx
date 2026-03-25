import React, { useMemo, useRef } from 'react';
import * as THREE from 'three';
import { Points, PointMaterial } from '@react-three/drei';

const GalaxyStars: React.FC = () => {
  const pointsRef = useRef<THREE.Points>(null);
  
  const count = 30000; // Increase count for better density
  
  const [positions, colors] = useMemo(() => {
    const pos = new Float32Array(count * 3);
    const cols = new Float32Array(count * 3);
    
    for (let i = 0; i < count; i++) {
      // Create spiral galaxy distribution with more volume
      const radius = Math.pow(Math.random(), 1.5) * 600000000;
      const spin = radius * 0.0000006;
      const angle = Math.random() * Math.PI * 2;
      const branch = (i % 4) * (Math.PI / 2); // 4 arms
      
      // Add more dispersion for 3D volume
      const x = Math.cos(angle + spin + branch) * radius + (Math.random() - 0.5) * radius * 0.3;
      const z = Math.sin(angle + spin + branch) * radius + (Math.random() - 0.5) * radius * 0.3;
      
      // Thickness: Thicker at center (Bulge), thinner at edges (Disk)
      const thicknessBase = 40000000; // Base thickness
      const bulgeEffect = Math.exp(-radius / 100000000) * 80000000; // Extra thickness at center
      const y = (Math.random() - 0.5) * (thicknessBase + bulgeEffect);
      
      pos[i * 3] = x;
      pos[i * 3 + 1] = y;
      pos[i * 3 + 2] = z;
      
      // Color based on distance and temperature
      const distRatio = radius / 600000000;
      const r = 0.8 + Math.random() * 0.2;
      const g = 0.7 + Math.random() * 0.3 - distRatio * 0.3;
      const b = 0.6 + Math.random() * 0.4 - distRatio * 0.2;
      
      cols[i * 3] = r;
      cols[i * 3 + 1] = g;
      cols[i * 3 + 2] = b;
    }
    return [pos, cols];
  }, [count]);

  return (
    <group>
      <Points ref={pointsRef} positions={positions} colors={colors} stride={3}>
        <PointMaterial
          transparent
          vertexColors
          size={5000} // Drastically reduced from 100,000
          sizeAttenuation={true}
          depthWrite={false}
          blending={THREE.AdditiveBlending}
          opacity={0.8}
        />
      </Points>
    </group>
  );
};

export default GalaxyStars;
