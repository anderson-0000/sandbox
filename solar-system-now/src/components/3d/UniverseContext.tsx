import React, { useMemo } from 'react';
import * as THREE from 'three';
import { Html } from '@react-three/drei';
import { useStore } from '../../hooks/useStore';

const UniverseContext: React.FC = () => {
  const viewMode = useStore((state) => state.viewMode);

  // Positions relative to sunPosition (current) or absolute?
  // Let's define Sagittarius A* as absolute (0,0,0) for simplicity in Galactic view,
  // but sunPosition moves.
  // Actually, Sagittarius A* is at some fixed point. 
  // Let's say Sagittarius A* is at (0,0,0).
  // Then sunPosition is at its current galactic position.
  
  const points = useMemo(() => ({
    sagittariusA: new THREE.Vector3(0, 0, 0),
    localGroupCM: new THREE.Vector3(10000000, 5000000, 20000000), // Approximate
    andromeda: new THREE.Vector3(20000000, 10000000, 40000000),
    greatAttractor: new THREE.Vector3(100000000, 50000000, -100000000),
  }), []);

  return (
    <group>
      {/* Sagittarius A* (Galactic Center) */}
      <mesh position={points.sagittariusA}>
        <sphereGeometry args={[500000, 32, 32]} />
        <meshBasicMaterial color="#ffccaa" transparent opacity={0.6} />
        <Html distanceFactor={5000000}>
          <div className="text-[10px] text-orange-200 whitespace-nowrap bg-black/50 px-2 py-1 rounded">いて座A* (銀河中心)</div>
        </Html>
      </mesh>

      {/* Local Group Common Center of Mass */}
      <mesh position={points.localGroupCM}>
        <sphereGeometry args={[200000, 32, 32]} />
        <meshBasicMaterial color="#aaccff" transparent opacity={0.4} />
        <Html distanceFactor={10000000}>
          <div className="text-[10px] text-blue-200 whitespace-nowrap bg-black/50 px-2 py-1 rounded">局所銀河群 共通重心</div>
        </Html>
      </mesh>

      {/* Andromeda */}
      <mesh position={points.andromeda}>
        <sphereGeometry args={[800000, 32, 32]} />
        <meshBasicMaterial color="#ccaaff" transparent opacity={0.3} />
        <Html distanceFactor={20000000}>
          <div className="text-[10px] text-purple-200 whitespace-nowrap bg-black/50 px-2 py-1 rounded">アンドロメダ銀河 (M31)</div>
        </Html>
      </mesh>

      {/* Great Attractor */}
      <mesh position={points.greatAttractor}>
        <sphereGeometry args={[5000000, 32, 32]} />
        <meshBasicMaterial color="#ffffff" transparent opacity={0.1} />
        <Html distanceFactor={100000000}>
          <div className="text-[10px] text-white whitespace-nowrap bg-black/50 px-2 py-1 rounded font-bold">グレート・アトラクター (巨大引力源)</div>
        </Html>
      </mesh>

      {/* Galactic Plane reference (around Sag A*) */}
      {(viewMode === 'milky_way' || viewMode === 'orion_arm') && (
        <gridHelper args={[100000000, 50, 0x444444, 0x222222]} rotation={[Math.PI / 2, 0, 0]} position={[0, 0, 0]} />
      )}
    </group>
  );
};

export default UniverseContext;