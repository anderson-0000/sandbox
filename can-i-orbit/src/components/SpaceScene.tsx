import { useRef, useMemo } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { OrbitControls, Stars, Sphere, Trail, Float } from '@react-three/drei';
import * as THREE from 'three';
import { EARTH_RADIUS } from '../physics/orbital_mechanics';
import type { State } from '../physics/orbital_mechanics';

// 描画用のスケール (1 = EARTH_RADIUS)
const SCALE = 1 / EARTH_RADIUS;

interface SatelliteProps {
  state: State;
}

function Satellite({ state }: SatelliteProps) {
  const meshRef = useRef<THREE.Mesh>(null);

  useFrame(() => {
    if (meshRef.current) {
      meshRef.current.position.set(
        state.position.x * SCALE,
        state.position.z * SCALE, // Y-up に合わせる
        -state.position.y * SCALE
      );
    }
  });

  return (
    <Trail
      width={0.8}
      color={'#00ffff'}
      length={50} // 軌道を長く見せる
      decay={1}
      local={false}
      stride={0}
    >
      <Sphere ref={meshRef} args={[0.03, 16, 16]}>
        <meshStandardMaterial color="#00ffff" emissive="#00ffff" emissiveIntensity={5} />
      </Sphere>
    </Trail>
  );
}

function Earth() {
  return (
    <group>
      {/* 地球本体: 裏側も見えるように DoubleSide を適用 */}
      <Sphere args={[1, 64, 64]}>
        <meshStandardMaterial
          color="#0033aa" // より鮮やかで深いロイヤルブルー
          emissive="#001144"
          emissiveIntensity={1}
          roughness={0.3}
          metalness={0.4}
          transparent={true}
          opacity={0.65} // 青さを強調するために少し不透明度を上げる
          side={THREE.DoubleSide}
        />
      </Sphere>
      
      {/* 大陸をイメージした第2レイヤー（うっすらとした緑） */}
      <Sphere args={[0.995, 32, 32]}>
        <meshStandardMaterial
          color="#1e5d1e" // 少し深い緑
          transparent={true}
          opacity={0.2}
          side={THREE.DoubleSide}
        />
      </Sphere>
      
      {/* 経緯度線（グリッド）: 裏側も見えるようにする */}
      <Sphere args={[1.001, 32, 32]}>
        <meshBasicMaterial
          color="#ffffff"
          wireframe={true}
          transparent={true}
          opacity={0.1}
          side={THREE.DoubleSide} // 裏側のグリッドもうっすら見える
        />
      </Sphere>

      {/* 大気層（外側の光の輪） */}
      <Sphere args={[1.08, 64, 64]}>
        <meshPhongMaterial
          color="#87ceeb"
          transparent={true}
          opacity={0.05}
          side={THREE.BackSide}
        />
      </Sphere>
    </group>
  );
}

interface SpaceSceneProps {
  state: State;
  isLaunched: boolean;
}

export function SpaceScene({ state, isLaunched }: SpaceSceneProps) {
  return (
    <div style={{ width: '100%', height: '100vh', background: '#000' }}>
      <Canvas camera={{ position: [5, 5, 5], fov: 45 }}>
        <color attach="background" args={['#000005']} />
        <ambientLight intensity={0.2} />
        <pointLight position={[10, 10, 10]} intensity={1.5} />
        
        <Stars radius={100} depth={50} count={5000} factor={4} saturation={0} fade speed={1} />
        
        <Earth />
        
        {isLaunched && <Satellite state={state} />}
        
        <OrbitControls 
          makeDefault 
          enableDamping 
          dampingFactor={0.05} 
          minDistance={1.1} 
          maxDistance={20} 
        />
      </Canvas>
    </div>
  );
}
