import React from 'react';
import { useFrame } from '@react-three/fiber';
import { useStore } from '../../hooks/useStore';
import { PLANETS_DATA } from '../../engine/kepler';
import Sun from './Sun';
import Planet from './Planet';
import OrbitLine from './OrbitLine';

const SolarSystem: React.FC = () => {
  const advanceTime = useStore((state) => state.advanceTime);

  // Time loop
  useFrame((_, delta) => {
    advanceTime(delta);
  });

  return (
    <>
      <Sun />
      {Object.values(PLANETS_DATA).map((planet) => (
        <React.Fragment key={planet.id}>
          <OrbitLine data={planet} />
          <Planet data={planet} />
        </React.Fragment>
      ))}
    </>
  );
};

export default SolarSystem;
