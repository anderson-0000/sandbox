import { useRef } from 'react';
import { useStore } from '../../store/useStore';
import * as THREE from 'three';

const P_WAVE_SPEED = 7000; 
const S_WAVE_SPEED = 4000; 
const GRAVITY = 9.8;
const WATER_DEPTH = 4000;
// 津波の速度を視覚的な分かりやすさのために 2.5 倍にブースト
const TSUNAMI_SPEED = Math.sqrt(GRAVITY * WATER_DEPTH) * 2.5; 

export const Sea = () => {
  const { epicenter, currentTime, magnitude } = useStore();
  const seaRef = useRef<THREE.Group>(null);

  const timeSinceQuake = epicenter ? currentTime - epicenter.time : 0;
  
  const pWaveRadius = timeSinceQuake > 0 ? timeSinceQuake * P_WAVE_SPEED : 0;
  const sWaveRadius = timeSinceQuake > 0 ? timeSinceQuake * S_WAVE_SPEED : 0;
  const tsunamiRadius = timeSinceQuake > 0 ? timeSinceQuake * TSUNAMI_SPEED : 0;

  // 厚みを画面上のピクセル密度に合わせて強制的に太くする
  // 最小でも 50,000m (50km) の厚みを持たせる
  const pWaveThickness = Math.max(20000, pWaveRadius * 0.05);
  const sWaveThickness = Math.max(20000, sWaveRadius * 0.05);
  const tsunamiThickness = Math.max(50000, tsunamiRadius * 0.2 * (magnitude / 7.0));

  return (
    <group ref={seaRef}>
      {/* 視界の邪魔にならない、非常に暗いグリッド */}
      <gridHelper args={[4000000, 40, 0x1e293b, 0x0f172a]} position={[0, -100, 0]} />
      
      {epicenter && (
        <group>
          {/* P Wave (Blue Ring) */}
          {pWaveRadius > 0 && pWaveRadius < 4000000 && (
            <mesh rotation={[-Math.PI / 2, 0, 0]} position={[epicenter.x, 100, epicenter.z]}>
              <ringGeometry args={[Math.max(0.1, pWaveRadius - pWaveThickness), pWaveRadius + 5000, 64]} />
              <meshBasicMaterial color="#3b82f6" transparent opacity={0.8} toneMapped={false} />
            </mesh>
          )}
          
          {/* S Wave (Red Ring) */}
          {sWaveRadius > 0 && sWaveRadius < 4000000 && (
            <mesh rotation={[-Math.PI / 2, 0, 0]} position={[epicenter.x, 200, epicenter.z]}>
              <ringGeometry args={[Math.max(0.1, sWaveRadius - sWaveThickness), sWaveRadius + 5000, 64]} />
              <meshBasicMaterial color="#ef4444" transparent opacity={0.9} toneMapped={false} />
            </mesh>
          )}

          {/* Tsunami (Cyan/White - Solid Disk and Glow) */}
          {tsunamiRadius > 0 && tsunamiRadius < 4000000 && (
            <group position={[epicenter.x, 0, epicenter.z]}>
              {/* 本体: 厚みのある不透明な層 */}
              <mesh rotation={[-Math.PI / 2, 0, 0]} position={[0, 400, 0]}>
                <ringGeometry args={[Math.max(0.1, tsunamiRadius - tsunamiThickness), tsunamiRadius, 128]} />
                <meshBasicMaterial color="#06b6d4" transparent opacity={0.7} toneMapped={false} />
              </mesh>
              {/* 波頭: 最も明るいエッジ */}
              <mesh rotation={[-Math.PI / 2, 0, 0]} position={[0, 500, 0]}>
                <ringGeometry args={[Math.max(0.1, tsunamiRadius - 10000), tsunamiRadius + 5000, 128]} />
                <meshBasicMaterial color="#e0f2fe" transparent opacity={1.0} toneMapped={false} />
              </mesh>
            </group>
          )}

          {/* 震源地の点（赤い球体） */}
          <mesh position={[epicenter.x, 5000, epicenter.z]}>
            <sphereGeometry args={[10000, 32, 32]} />
            <meshBasicMaterial color="#ff0000" toneMapped={false} />
          </mesh>

          {/* 震央マーカー（垂直に伸びる光の柱） */}
          <mesh position={[epicenter.x, 250000, epicenter.z]}>
            <cylinderGeometry args={[2000, 2000, 500000, 16]} />
            <meshBasicMaterial color="#ff0000" transparent opacity={0.3} toneMapped={false} />
          </mesh>
        </group>
      )}
    </group>
  );
};
