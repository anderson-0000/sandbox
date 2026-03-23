import React from 'react';
import * as THREE from 'three';
import { useFrame } from '@react-three/fiber';
import { useStore } from '../../hooks/useStore';
import { PLANETS_DATA } from '../../engine/kepler';
import Sun from './Sun';
import Planet from './Planet';
import OrbitLine from './OrbitLine';
import SunOrbitLine from './SunOrbitLine';

const SolarSystem: React.FC = () => {
  const groupRef = React.useRef<THREE.Group>(null);
  const advanceTime = useStore((state) => state.advanceTime);
  const sunPosition = useStore((state) => state.sunPosition);
  const showOrbits = useStore((state) => state.showOrbits);
  const showSunOrbit = useStore((state) => state.showSunOrbit);

  // Time loop
  useFrame((_, delta) => {
    advanceTime(delta);
    if (groupRef.current) {
      groupRef.current.position.copy(sunPosition);
    }
  });

  return (
    <>
      {showSunOrbit && <SunOrbitLine />}
      <group ref={groupRef}>
        <Sun />
        {Object.values(PLANETS_DATA).map((planet) => (
          <React.Fragment key={planet.id}>
            {showOrbits && <OrbitLine data={planet} />}
            <Planet data={planet} />
          </React.Fragment>
        ))}
      </group>
    </>
  );
};

export default SolarSystem;