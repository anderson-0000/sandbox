import React, { Suspense } from 'react';
import { Canvas } from '@react-three/fiber';
import { Stars, PerspectiveCamera } from '@react-three/drei';
import SolarSystem from './SolarSystem';
import CameraController from './CameraController';
import SpaceDust from './SpaceDust';
import UniverseContext from './UniverseContext';
import Tracker from './Tracker';

const Scene: React.FC = () => {
  return (
    <div className="w-full h-full relative">
      <Canvas 
        shadows={false} // Disable shadows for debugging
        gl={{ 
          antialias: true,
          alpha: true
        }}
      >
        <Suspense fallback={null}>
          <PerspectiveCamera 
            makeDefault 
            position={[0, 1000, 2000]} // Start closer to center
            fov={45} 
            far={1000000} // Reduce far plane to a more reasonable range
            near={1} 
          />
          <Tracker />
          <CameraController />
          
          <ambientLight intensity={1.5} /> // Much stronger light
          <Stars radius={100000} depth={50000} count={5000} factor={4} saturation={0} fade speed={1} />
          
          <SpaceDust />
          <UniverseContext />
          <SolarSystem />
        </Suspense>
      </Canvas>
    </div>
  );
};

export default Scene;
