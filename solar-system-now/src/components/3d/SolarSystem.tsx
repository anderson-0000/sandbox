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
  const advanceTime = useStore((state) => state.advanceTime);
  const showOrbits = useStore((state) => state.showOrbits);
  const showSunOrbit = useStore((state) => state.showSunOrbit);

  // Time loop
  useFrame((_, delta) => {
    advanceTime(delta);
  });

  return (
    <>
      <Sun />
      {showSunOrbit && <SunOrbitLine />}
      {Object.values(PLANETS_DATA).map((planet) => (
        <React.Fragment key={planet.id}>
          {showOrbits && <OrbitLine data={planet} />}
          <Planet data={planet} />
        </React.Fragment>
      ))}
    </>
  );
};

export default SolarSystem;