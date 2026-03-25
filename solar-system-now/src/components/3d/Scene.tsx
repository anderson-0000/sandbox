import React, { Suspense } from 'react';
import { Canvas } from '@react-three/fiber';
import { Stars, PerspectiveCamera } from '@react-three/drei';
import SolarSystem from './SolarSystem';
import CameraController from './CameraController';
import SpaceDust from './SpaceDust';
import UniverseContext from './UniverseContext';
import Tracker from './Tracker';
import GalaxyStars from './GalaxyStars';

const Scene: React.FC = () => {
  return (
    <div className="w-full h-full relative">
      <Canvas 
        shadows={false} 
        gl={{ 
          antialias: true,
          alpha: true
        }}
      >
        <Suspense fallback={null}>
          <PerspectiveCamera 
            makeDefault 
            position={[0, 1000, 2000]} 
            fov={45} 
            far={10000000000} // Increased far plane for Galactic view
            near={1} 
          />
          <Tracker />
          <CameraController />
          
          <ambientLight intensity={1.5} />
          
          {/* Universal background stars */}
          <Stars 
            radius={2000000000} 
            depth={1000000000} 
            count={20000} 
            factor={10} 
            saturation={0} 
            fade 
            speed={1} 
          />
          
          {/* Detailed Milky Way distribution */}
          <GalaxyStars />
          
          <SpaceDust />
          <UniverseContext />
          <SolarSystem />
        </Suspense>
      </Canvas>
    </div>
  );
};

export default Scene;
