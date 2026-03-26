import { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import * as THREE from 'three';
import { useKeyboard } from '../hooks/useKeyboard';

interface ShipProps {
  onUpdate: (position: THREE.Vector3) => void;
}

export const Ship = ({ onUpdate }: ShipProps) => {
  const meshRef = useRef<THREE.Group>(null);
  const keys = useKeyboard();
  const velocity = useRef(new THREE.Vector3());
  const speed = 0.1;

  useFrame(() => {
    if (!meshRef.current) return;

    // 移動ロジック
    const moveVector = new THREE.Vector3();
    if (keys.forward) moveVector.z -= speed;
    if (keys.backward) moveVector.z += speed;
    if (keys.left) moveVector.x -= speed;
    if (keys.right) moveVector.x += speed;
    if (keys.up) moveVector.y += speed;
    if (keys.down) moveVector.y -= speed;

    velocity.current.lerp(moveVector, 0.1);
    meshRef.current.position.add(velocity.current);

    // 親コンポーネントに位置を通知（カメラや衝突判定用）
    onUpdate(meshRef.current.position);
  });

  return (
    <group ref={meshRef}>
      {/* 宇宙船のモデル */}
      <mesh>
        <boxGeometry args={[0.5, 0.2, 1]} />
        <meshStandardMaterial color="#00ffcc" emissive="#00ffcc" emissiveIntensity={0.5} />
      </mesh>
      <mesh position={[0, 0, 0.2]}>
        <boxGeometry args={[1.5, 0.05, 0.4]} />
        <meshStandardMaterial color="#00ffcc" />
      </mesh>
      {/* 前方を示すライト */}
      <pointLight position={[0, 0, -1]} distance={5} intensity={1} color="#00ffcc" />
    </group>
  );
};
