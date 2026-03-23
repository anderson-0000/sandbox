import React, { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';
import { useStore } from '../../hooks/useStore';

interface MovementVectorProps {
  velocity: THREE.Vector3;
  color?: string;
  scale?: number;
}

const MovementVector: React.FC<MovementVectorProps> = ({ velocity, color = '#ffffff', scale = 1 }) => {
  const arrowRef = useRef<THREE.ArrowHelper>(null);
  const showVectors = useStore((state) => state.showVectors);

  useFrame(() => {
    if (arrowRef.current) {
      const length = velocity.length();
      if (length > 0) {
        arrowRef.current.setDirection(velocity.clone().normalize());
        arrowRef.current.setLength(length * scale, length * scale * 0.2, length * scale * 0.1);
      }
      arrowRef.current.visible = showVectors;
    }
  });

  return showVectors ? (
    <arrowHelper
      ref={arrowRef}
      args={[new THREE.Vector3(0, 1, 0), new THREE.Vector3(0, 0, 0), 1, color]}
    />
  ) : null;
};

export default MovementVector;