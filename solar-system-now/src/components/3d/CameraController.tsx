import React, { useRef } from 'react';
import { useFrame } from '@react-three/fiber';
import { OrbitControls } from '@react-three/drei';
import * as THREE from 'three';
import { useStore } from '../../hooks/useStore';
import { PLANETS_DATA, MOONS_DATA, getPlanetPosition } from '../../engine/kepler';

const CameraController: React.FC = () => {
  const controlsRef = useRef<any>(null);
  const selectedPlanet = useStore((state) => state.selectedPlanet);
  const currentDate = useStore((state) => state.currentDate);
  const targetPos = useRef(new THREE.Vector3(0, 0, 0));
  const prevTargetPos = useRef(new THREE.Vector3(0, 0, 0));
  const lastSelectedPlanet = useRef<string | null>(null);
  const isFollowing = useRef(false);

  useFrame((state, delta) => {
    // Reset following state if planet changed
    if (lastSelectedPlanet.current !== selectedPlanet) {
      isFollowing.current = false;
      lastSelectedPlanet.current = selectedPlanet;
    }

    if (selectedPlanet === 'Sun') {
      targetPos.current.set(0, 0, 0);
    } else if (selectedPlanet && PLANETS_DATA[selectedPlanet]) {
      const pos = getPlanetPosition(PLANETS_DATA[selectedPlanet], currentDate);
      targetPos.current.copy(pos);
    } else if (selectedPlanet) {
      // Find moon in MOONS_DATA
      let moonData = null;
      for (const parent in MOONS_DATA) {
        const found = MOONS_DATA[parent].find(m => m.id === selectedPlanet);
        if (found) {
          moonData = found;
          break;
        }
      }

      if (moonData) {
        const parentPlanet = PLANETS_DATA[moonData.parent];
        const planetPos = getPlanetPosition(parentPlanet, currentDate);
        const moonRelPos = getPlanetPosition(moonData, currentDate, true);
        targetPos.current.copy(planetPos).add(moonRelPos);
      }
    } else {
      targetPos.current.set(0, 0, 0);
    }

    if (controlsRef.current) {
      if (selectedPlanet) {
        if (!isFollowing.current) {
          // Smoothly transition to the target
          controlsRef.current.target.lerp(targetPos.current, delta * 5);
          
          // If close enough, start exact following
          if (controlsRef.current.target.distanceTo(targetPos.current) < 1) {
            isFollowing.current = true;
          }
        } else {
          // Maintain relative distance by moving camera with the target
          const movement = targetPos.current.clone().sub(prevTargetPos.current);
          state.camera.position.add(movement);
          controlsRef.current.target.copy(targetPos.current);
        }
      } else {
        // Standard lerp for no selection
        controlsRef.current.target.lerp(targetPos.current, delta * 5);
      }
      
      controlsRef.current.update();
    }
    
    prevTargetPos.current.copy(targetPos.current);
  });

  return (
    <OrbitControls
      ref={controlsRef}
      makeDefault
      enableDamping
      dampingFactor={0.05}
      minDistance={1}
      maxDistance={50000}
    />
  );
};

export default CameraController;
