import { useRef, useEffect } from 'react';
import { useFrame } from '@react-three/fiber';
import { Sphere, MeshDistortMaterial, Stars, Text, PerspectiveCamera, OrbitControls } from '@react-three/drei';
import * as THREE from 'three';

interface ProgressArrowProps {
  progress: number;
}

const ProgressArrow = ({ progress }: ProgressArrowProps) => {
  const meshRef = useRef<THREE.Mesh>(null);
  const length = progress * 100;
  
  useFrame(() => {
    if (meshRef.current) {
      const targetLength = Math.max(length, 0.01);
      meshRef.current.scale.x = THREE.MathUtils.lerp(meshRef.current.scale.x, targetLength, 0.05);
      meshRef.current.position.x = meshRef.current.scale.x / 2;
    }
  });

  return (
    <group>
      <mesh ref={meshRef}>
        <boxGeometry args={[1, 0.5, 0.5]} />
        <meshStandardMaterial color="#44ff88" emissive="#44ff88" emissiveIntensity={0.5} transparent opacity={0.8} />
      </mesh>
      
      <mesh position={[length, 0, 0]} rotation={[0, 0, -Math.PI / 2]}>
        <coneGeometry args={[0.8, 2, 8]} />
        <meshStandardMaterial color="#44ff88" />
      </mesh>
    </group>
  );
};

export const SpaceScene = ({ progress, targetX }: { progress: number; targetX: number }) => {
  const controlsRef = useRef<any>(null);

  // targetX が変わったときに、OrbitControls の target をスムーズに移動させる
  useFrame(() => {
    if (controlsRef.current) {
      const currentTarget = controlsRef.current.target;
      currentTarget.x = THREE.MathUtils.lerp(currentTarget.x, targetX, 0.05);
      controlsRef.current.update();
    }
  });

  return (
    <>
      <PerspectiveCamera makeDefault position={[targetX, 20, 50]} />
      <OrbitControls 
        ref={controlsRef}
        makeDefault 
        minDistance={10} 
        maxDistance={300} 
        enablePan={true}
      />
      
      <ambientLight intensity={0.5} />
      <pointLight position={[10, 10, 10]} intensity={1.5} />
      <Stars radius={100} depth={50} count={5000} factor={4} saturation={0} fade speed={1} />

      {/* 地球 */}
      <group position={[0, 0, 0]}>
        <Sphere args={[2, 32, 32]}>
          <meshStandardMaterial color="#4488ff" wireframe />
        </Sphere>
        <Text position={[0, -3, 0]} fontSize={1} color="white">
          地球 (開始)
        </Text>
      </group>

      {/* 目標惑星 */}
      <group position={[100, 0, 0]}>
        <Sphere args={[5, 64, 64]}>
          <MeshDistortMaterial color="#ff8844" speed={2} distort={0.4} />
        </Sphere>
        <Text position={[0, -7, 0]} fontSize={1.5} color="white">
          目標惑星 (ゴール)
        </Text>
      </group>

      <ProgressArrow progress={progress} />
    </>
  );
};
