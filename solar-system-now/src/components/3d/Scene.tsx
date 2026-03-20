import React, { Suspense } from 'react';
import { Canvas } from '@react-three/fiber';
import { Stars, PerspectiveCamera } from '@react-three/drei';
import * as THREE from 'three';
import SolarSystem from './SolarSystem';
import CameraController from './CameraController';
import SpaceDust from './SpaceDust';

const Scene: React.FC = () => {
  return (
    <div className="w-full h-full relative">
      <Canvas 
        shadows 
        gl={{ 
          antialias: true,
          toneMapping: THREE.ACESFilmicToneMapping,
          toneMappingExposure: 0.8
        }}
      >
        <Suspense fallback={null}>
          <PerspectiveCamera 
            makeDefault 
            position={[5000, 3000, 8000]} 
            fov={45} 
            far={2000000} 
            near={0.01} // Extreme close-ups allowed
          />
          <CameraController />
          
          <ambientLight intensity={0.05} />
          <Stars radius={150000} depth={50000} count={50000} factor={40} saturation={0.5} fade speed={0.5} />
          
          <SpaceDust />
          <SolarSystem />
        </Suspense>
      </Canvas>
    </div>
  );
};

export default Scene;
