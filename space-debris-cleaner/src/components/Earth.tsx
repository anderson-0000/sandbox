import { useRef } from 'react';
import { Sphere } from '@react-three/drei';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';

export const Earth = () => {
  const meshRef = useRef<THREE.Mesh>(null);

  useFrame(() => {
    if (meshRef.current) {
      meshRef.current.rotation.y += 0.001;
    }
  });

  return (
    <Sphere ref={meshRef} args={[1, 64, 64]}>
      <meshStandardMaterial color="#2266aa" wireframe />
    </Sphere>
  );
};

interface DebrisFieldProps {
  debris: any[];
  onCollect: (id: string) => void;
  updatePositions: () => void;
}

export const DebrisField = ({ debris, onCollect, updatePositions }: DebrisFieldProps) => {
  // Update positions every frame
  useFrame(() => {
    updatePositions();
  });

  return (
    <>
      {debris.map((d) => (
        !d.isCollected && (
          <mesh 
            key={d.id} 
            position={d.position} 
            onClick={() => onCollect(d.id)}
          >
            <sphereGeometry args={[0.01, 8, 8]} />
            <meshBasicMaterial color="#ffcc00" />
          </mesh>
        )
      ))}
    </>
  );
};
