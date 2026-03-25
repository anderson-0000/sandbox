import React, { useMemo } from 'react';
import * as THREE from 'three';
import { Html } from '@react-three/drei';
import { useStore } from '../../hooks/useStore';

const UniverseContext: React.FC = () => {
  const viewMode = useStore((state) => state.viewMode);
  const showLabels = useStore((state) => state.showLabels);

  // Positions in arbitrary galactic coordinates
  const points = useMemo(() => ({
    sagittariusA: new THREE.Vector3(0, 0, 0),
    localGroupCM: new THREE.Vector3(10000000, 5000000, 20000000), 
    andromeda: new THREE.Vector3(25000000, 12000000, 45000000),
    greatAttractor: new THREE.Vector3(150000000, 80000000, -120000000),
  }), []);

  // Display conditions for large scale objects
  const showGalacticCenter = showLabels && ['milky_way', 'local_group', 'virgo_supercluster'].includes(viewMode);
  const showLocalGroup = showLabels && ['local_group', 'virgo_supercluster'].includes(viewMode);
  const showSupercluster = showLabels && ['virgo_supercluster'].includes(viewMode);
  
  // Orion Arm mode should only show the Sun's vicinity
  const isGalacticPlus = ['milky_way', 'local_group', 'virgo_supercluster', 'galactic'].includes(viewMode);

  return (
    <group>
      {/* Sagittarius A* (Galactic Center) - Only label, no mesh as requested */}
      {showGalacticCenter && (
        <group position={points.sagittariusA}>
          <Html distanceFactor={50000000} center>
            <div className="text-[12px] text-orange-200 whitespace-nowrap bg-black/70 px-2 py-1 rounded border border-orange-500/30">いて座A* (銀河中心)</div>
          </Html>
        </group>
      )}

      {/* Local Group Common Center of Mass */}
      {showLocalGroup && (
        <>
          <group position={points.localGroupCM}>
            <Html distanceFactor={100000000} center>
              <div className="text-[12px] text-blue-200 whitespace-nowrap bg-black/70 px-2 py-1 rounded border border-blue-500/30">局所銀河群 共通重心</div>
            </Html>
          </group>

          {/* Andromeda */}
          <group position={points.andromeda}>
            <Html distanceFactor={150000000} center>
              <div className="text-[12px] text-purple-200 whitespace-nowrap bg-black/70 px-2 py-1 rounded border border-purple-500/30">アンドロメダ銀河 (M31)</div>
            </Html>
          </group>
        </>
      )}

      {/* Great Attractor */}
      {showSupercluster && (
        <group position={points.greatAttractor}>
          <Html distanceFactor={500000000} center>
            <div className="text-[14px] text-white whitespace-nowrap bg-black/80 px-3 py-1 rounded border border-white/50 font-bold">グレート・アトラクター (巨大引力源)</div>
          </Html>
        </group>
      )}

      {/* Galactic Plane reference */}
      {isGalacticPlus && (
        <gridHelper 
          args={[500000000, 100, 0x444444, 0x222222]} 
          rotation={[Math.PI / 2, 0, 0]} 
          position={[0, 0, 0]} 
        />
      )}
    </group>
  );
};

export default UniverseContext;
