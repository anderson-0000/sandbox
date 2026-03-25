import React, { useRef, useEffect } from 'react';
import * as THREE from 'three';
import { useThree } from '@react-three/fiber';
import { CameraControls } from '@react-three/drei';
import { useStore } from '../../hooks/useStore';

const CameraController: React.FC = () => {
  const cameraControlsRef = useRef<CameraControls | null>(null);
  const { scene } = useThree();

  const selectedObjectName = useStore((state) => state.selectedObjectName);
  const viewMode = useStore((state) => state.viewMode);
  const zoomDistance = useStore((state) => state.zoomDistance);
  const setZoomDistance = useStore((state) => state.setZoomDistance);
  const sunPosition = useStore((state) => state.sunPosition);
  
  const isInternalUpdate = useRef(false);

  // Constants for specific centers
  const SAGITTARIUS_A = new THREE.Vector3(0, 0, 0);
  const LOCAL_GROUP_CM = new THREE.Vector3(10000000, 5000000, 20000000);
  const GREAT_ATTRACTOR = new THREE.Vector3(100000000, 50000000, -100000000);

  // Synchronize camera to store only during user interaction (pinch/scroll)
  useEffect(() => {
    const controls = cameraControlsRef.current;
    if (!controls) return;

    const onControl = () => {
      if (!isInternalUpdate.current) {
        setZoomDistance(controls.distance);
      }
    };

    controls.addEventListener('control', onControl);
    return () => controls.removeEventListener('control', onControl);
  }, [setZoomDistance]);

  // Synchronize store to camera (Slider -> Camera)
  useEffect(() => {
    if (!cameraControlsRef.current) return;
    const controls = cameraControlsRef.current;
    if (Math.abs(controls.distance - zoomDistance) > (zoomDistance * 0.01)) {
      isInternalUpdate.current = true;
      controls.dollyTo(zoomDistance, false);
      setTimeout(() => { isInternalUpdate.current = false; }, 50);
    }
  }, [zoomDistance]);

  // Handle Focus (selectedObjectName) and Mode Changes
  useEffect(() => {
    if (!cameraControlsRef.current) return;
    const controls = cameraControlsRef.current;
    
    let target = new THREE.Vector3();
    let dist = zoomDistance;

    const targetObj = selectedObjectName ? scene.getObjectByName(selectedObjectName) : (selectedObjectName === 'Sun' ? scene.getObjectByName('Sun') : null);

    if (targetObj && (viewMode === 'solar_system' || viewMode === 'earth')) {
        targetObj.updateWorldMatrix(true, false);
        targetObj.getWorldPosition(target);
        
        // Use radius for appropriate zoom distance
        const radius = targetObj.userData?.radius || 1;
        dist = radius * 4; // Closer zoom
        if (selectedObjectName === 'Sun') dist = radius * 2.5;
    } else {
        switch(viewMode) {
            case 'earth':
                const earth = scene.getObjectByName('Earth');
                if (earth) {
                    earth.updateWorldMatrix(true, false);
                    earth.getWorldPosition(target);
                    const radius = earth.userData?.radius || 1;
                    dist = radius * 4;
                } else {
                    dist = 20;
                }
                break;
            case 'solar_system':
                target.copy(sunPosition);
                dist = 5000;
                break;
            case 'orion_arm':
                target.copy(SAGITTARIUS_A);
                dist = 500000;
                break;
            case 'milky_way':
                target.copy(SAGITTARIUS_A);
                dist = 5000000;
                break;
            case 'local_group':
                target.copy(LOCAL_GROUP_CM);
                dist = 30000000;
                break;
            case 'virgo_supercluster':
                target.copy(GREAT_ATTRACTOR);
                dist = 150000000;
                break;
            case 'galactic':
                target.copy(sunPosition);
                dist = 5000;
                break;
        }
    }

    isInternalUpdate.current = true;
    controls.setLookAt(target.x, target.y + dist, target.z + dist, target.x, target.y, target.z, true);
    setZoomDistance(dist);
    setTimeout(() => { isInternalUpdate.current = false; }, 100);
  }, [viewMode, selectedObjectName, sunPosition, scene, setZoomDistance, zoomDistance]);

  return (
    <CameraControls
      ref={cameraControlsRef}
      minDistance={0.001}
      maxDistance={2000000000}
      dollyToCursor={true}
      dollySpeed={3.0}
      truckSpeed={0}
      mouseButtons={{
        left: 1,
        middle: 8,
        right: 0,
        wheel: 8,
      }}
    />
  );
};

export default CameraController;
