import Scene from './components/3d/Scene';
import TimeControls from './components/ui/TimeControls';
import InfoPanel from './components/ui/InfoPanel';
import PlanetSelector from './components/ui/PlanetSelector';

function App() {
  return (
    <div className="w-screen h-screen relative bg-black overflow-hidden text-white">
      {/* 3D Scene - Rendered first, serves as background */}
      <div className="absolute inset-0 z-0">
        <Scene />
      </div>

      {/* UI Layers - Using fixed positioning for components to avoid blocking the whole screen */}
      <div className="fixed inset-0 z-50 pointer-events-none">
        <div className="relative w-full h-full">
          <PlanetSelector />
          <InfoPanel />
          <TimeControls />
        </div>
      </div>
    </div>
  );
}

export default App;
