import React, { Suspense } from 'react';
import { Canvas } from '@react-three/fiber';
import { Stars, PerspectiveCamera } from '@react-three/drei';
import SolarSystem from './SolarSystem';
import CameraController from './CameraController';

const Scene: React.FC = () => {
  return (
    <div className="w-full h-full relative">
      <Canvas shadows>
        <Suspense fallback={null}>
          <PerspectiveCamera makeDefault position={[5000, 3000, 8000]} fov={45} far={200000} />
          <CameraController />
          
          <ambientLight intensity={0.2} />
          <Stars radius={50000} depth={10000} count={30000} factor={40} saturation={0.5} fade speed={0.5} />
          
          <SolarSystem />
        </Suspense>
      </Canvas>
    </div>
  );
};

export default Scene;
