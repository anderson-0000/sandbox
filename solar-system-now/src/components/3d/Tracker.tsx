import React from 'react';
import { useFrame, useThree } from '@react-three/fiber';
import { useStore } from '../../hooks/useStore';
import * as THREE from 'three';

/**
 * Tracker component to "escape" 3D state to the parent React/Zustand state.
 * This follows the guideline in GEMINI.md to prevent crashes and maintain clean separation.
 */
const Tracker: React.FC = () => {
  const { scene } = useThree();
  const setCameraTarget = useStore((state) => state.setCameraTarget);
  const setSunPosition = useStore((state) => state.setSunPosition);
  const selectedObjectName = useStore((state) => state.selectedObjectName);
  const viewMode = useStore((state) => state.viewMode);

  useFrame(() => {
    // 1. Track Sun position for global reference if needed
    const sun = scene.getObjectByName('Sun');
    if (sun) {
      const sunPos = new THREE.Vector3();
      sun.getWorldPosition(sunPos);
      setSunPosition(sunPos);
    }

    // 2. Track Selected Object for Camera Target
    let targetPos = new THREE.Vector3(0, 0, 0);
    if (viewMode === 'solar_system' || viewMode === 'earth' || viewMode === 'galactic') {
      const targetObj = scene.getObjectByName(selectedObjectName || 'Sun');
      if (targetObj) {
        targetObj.getWorldPosition(targetPos);
      }
    }

    // Sanity check
    if (!isNaN(targetPos.x) && !isNaN(targetPos.y) && !isNaN(targetPos.z)) {
      setCameraTarget(targetPos);
    }
  });

  return null;
};

export default Tracker;
