import { useCallback, useRef } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { Stars, PerspectiveCamera, OrbitControls } from '@react-three/drei';
import * as THREE from 'three';
import { useDebris } from './hooks/useDebris';
import { Earth } from './components/Earth';
import { Ship } from './components/Ship';
import './App.css';

// デブリ描画と追跡のマネージャー
const DebrisAndControls = ({ debris, updatePositions, onCollect, shipPosRef }: any) => {
  const controlsRef = useRef<any>(null);
  const lineRef = useRef<THREE.Line>(null);

  useFrame(() => {
    updatePositions();
    
    // カメラの注視点（target）を宇宙船に合わせる
    if (controlsRef.current && shipPosRef.current) {
      controlsRef.current.target.lerp(shipPosRef.current, 0.1);
      controlsRef.current.update();
    }

    // 補助線の更新
    if (lineRef.current) {
      const positions = new Float32Array([0, 0, 0, ...shipPosRef.current.toArray()]);
      lineRef.current.geometry.setAttribute('position', new THREE.BufferAttribute(positions, 3));
      lineRef.current.geometry.attributes.position.needsUpdate = true;
    }

    // 衝突判定
    debris.forEach((d: any) => {
      if (!d.isCollected) {
        const dPos = new THREE.Vector3(...d.position);
        if (shipPosRef.current.distanceTo(dPos) < 0.4) {
          onCollect(d.id);
        }
      }
    });
  });

  return (
    <>
      <OrbitControls ref={controlsRef} enablePan={false} makeDefault minDistance={2} maxDistance={100} />
      
      {/* デブリ描画 */}
      {debris.map((d: any) => (
        !d.isCollected && (
          <mesh key={d.id} position={d.position}>
            <sphereGeometry args={[0.08, 8, 8]} />
            <meshBasicMaterial color="#ffcc00" />
          </mesh>
        )
      ))}

      {/* 補助線: 宇宙船から地球の中心へ (高さ把握用) */}
      <line ref={lineRef}>
        <bufferGeometry />
        <lineBasicMaterial color="#00ffcc" transparent opacity={0.3} />
      </line>
    </>
  );
};

function App() {
  const { debris, score, updatePositions, collectDebris } = useDebris();
  const shipPosRef = useRef(new THREE.Vector3(0, 0, 5));

  const handleShipUpdate = useCallback((pos: THREE.Vector3) => {
    shipPosRef.current.copy(pos);
  }, []);

  return (
    <div className="game-container">
      <div className="ui">
        <h1>スペース・デブリ・アクション</h1>
        <p>宇宙船を操作してデブリを回収してください</p>
        <div className="controls-hint">
          <b>W/A/S/D</b>: 前後左右, <b>R/F</b>: 上昇/下降<br />
          <b>ドラッグ</b>: 視点回転, <b>ホイール</b>: ズーム
        </div>
        <div className="score">スコア: {score}</div>
        <p>残りデブリ: {debris.filter(d => !d.isCollected).length}</p>
      </div>

      <div className="canvas-container">
        <Canvas>
          <PerspectiveCamera makeDefault position={[10, 10, 10]} fov={60} />
          <ambientLight intensity={0.5} />
          <pointLight position={[10, 10, 10]} intensity={1.5} />
          <Stars radius={200} depth={100} count={10000} factor={6} saturation={0} fade speed={1} />
          
          <Earth />
          
          <DebrisAndControls 
            debris={debris} 
            updatePositions={updatePositions} 
            onCollect={collectDebris}
            shipPosRef={shipPosRef}
          />
          
          <Ship onUpdate={handleShipUpdate} />
        </Canvas>
      </div>
    </div>
  );
}

export default App;
