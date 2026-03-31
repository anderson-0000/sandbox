import { Canvas } from '@react-three/fiber';
import { OrbitControls, PerspectiveCamera } from '@react-three/drei';
import { useStore } from './store/useStore';
import { Sea } from './components/canvas/Sea';
import { Tracker } from './components/Tracker';
import { ControlPanel } from './components/ui/ControlPanel';

/**
 * クリックを受け取るための透明な地面コンポーネント
 */
const ClickSurface = () => {
  const setEpicenter = useStore((state) => state.setEpicenter);

  return (
    <mesh 
      rotation={[-Math.PI / 2, 0, 0]} 
      position={[0, -10, 0]}
      onClick={(e) => {
        e.stopPropagation();
        setEpicenter(e.point.x, e.point.z);
      }}
    >
      <planeGeometry args={[4000000, 4000000]} />
      <meshBasicMaterial transparent opacity={0} />
    </mesh>
  );
};

function App() {
  return (
    <div className="w-screen h-screen bg-[#020617] overflow-hidden font-sans">
      <ControlPanel />
      
      <div className="absolute bottom-6 right-6 z-10 text-right pointer-events-none">
        <h2 className="text-white font-black text-4xl italic tracking-tighter opacity-20">DISASTER PULSE</h2>
        <p className="text-slate-500 text-[10px] font-bold uppercase tracking-[0.3em] mt-1">Experimental Simulation Platform</p>
      </div>

      <Canvas
        shadows
        gl={{ antialias: true, alpha: false }}
        style={{ background: '#020617' }}
      >
        <color attach="background" args={['#020617']} />
        <PerspectiveCamera makeDefault position={[800000, 600000, 800000]} far={10000000} />
        <OrbitControls 
          maxPolarAngle={Math.PI / 2.1} 
          minDistance={1000} 
          maxDistance={5000000}
        />
        
        {/* 全体を明るくしすぎない程度の光 */}
        <ambientLight intensity={0.5} />
        <pointLight position={[1000000, 1000000, 1000000]} intensity={1} />

        <Sea />
        <ClickSurface />
        <Tracker />
      </Canvas>
    </div>
  );
}

export default App;
