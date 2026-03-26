import { Canvas } from '@react-three/fiber';
import { OrbitControls, Stars } from '@react-three/drei';
import { useDebris } from './hooks/useDebris';
import { Earth, DebrisField } from './components/Earth';
import './App.css';

function App() {
  const { debris, score, updatePositions, collectDebris } = useDebris();

  return (
    <div className="game-container">
      <div className="ui">
        <h1>Space Debris Cleaner</h1>
        <p>地球の周りのゴミをクリックして回収しよう</p>
        <div className="score">Score: {score}</div>
        <p>残りデブリ: {debris.filter(d => !d.isCollected).length}</p>
      </div>

      <div className="canvas-container">
        <Canvas camera={{ position: [0, 0, 3], fov: 45 }}>
          <ambientLight intensity={0.5} />
          <pointLight position={[10, 10, 10]} intensity={1.5} />
          <Stars radius={100} depth={50} count={5000} factor={4} saturation={0} fade speed={1} />
          
          <Earth />
          <DebrisField 
            debris={debris} 
            onCollect={collectDebris} 
            updatePositions={updatePositions} 
          />
          
          <OrbitControls enablePan={false} minDistance={1.2} maxDistance={10} />
        </Canvas>
      </div>
    </div>
  );
}

export default App;
