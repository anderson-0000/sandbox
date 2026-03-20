import Scene from './components/3d/Scene';
import TimeControls from './components/ui/TimeControls';
import InfoPanel from './components/ui/InfoPanel';
import PlanetSelector from './components/ui/PlanetSelector';

function App() {
  return (
    <div className="w-screen h-screen relative bg-black overflow-hidden text-white">
      {/* 3D Scene (Background) */}
      <div className="absolute inset-0 z-0">
        <Scene />
      </div>

      {/* UI Components - They have their own absolute/fixed positioning and high z-index */}
      <PlanetSelector />
      <InfoPanel />
      <TimeControls />
    </div>
  );
}

export default App;
