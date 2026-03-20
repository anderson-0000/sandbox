import React, { useRef, useEffect } from 'react';
import * as THREE from 'three';
import { useFrame, useThree } from '@react-three/fiber';
import { CameraControls } from '@react-three/drei';
import { useStore } from '../../hooks/useStore';
import { SUN_GALACTIC_VELOCITY } from '../../engine/kepler';

const CameraController: React.FC = () => {
  const cameraControlsRef = useRef<CameraControls | null>(null);
  const { scene } = useThree();

  const selectedObjectName = useStore((state) => state.selectedObjectName);
  const surfaceTargetName = useStore((state) => state.surfaceTargetName);
  const viewMode = useStore((state) => state.viewMode);
  const setCameraTarget = useStore((state) => state.setCameraTarget);
  const zoomDistance = useStore((state) => state.zoomDistance);
  const setZoomDistance = useStore((state) => state.setZoomDistance);
  const sunPosition = useStore((state) => state.sunPosition);
  
  const lastTargetPosition = useRef<THREE.Vector3>(new THREE.Vector3());
  const prevSelectedObjectName = useRef(selectedObjectName);
  const initialized = useRef(false);
  const isInternalUpdate = useRef(false);

  useFrame(() => {
    if (!cameraControlsRef.current) return;
    const controls = cameraControlsRef.current;

    // Sync distance to store
    if (initialized.current && viewMode !== 'galactic') {
      const actualDistance = controls.distance;
      if (Math.abs(actualDistance - zoomDistance) > 0.01 && !isInternalUpdate.current) {
        setZoomDistance(actualDistance);
      }
    }

    const selectedObject = selectedObjectName ? scene.getObjectByName(selectedObjectName) : (selectedObjectName === 'Sun' ? scene.getObjectByName('Sun') : null);
    const targetObject = selectedObject || scene.getObjectByName('Sun');
    if (!targetObject) return;

    const currentTargetPosition = new THREE.Vector3();
    targetObject.getWorldPosition(currentTargetPosition);

    // Initial positioning
    if (prevSelectedObjectName.current !== selectedObjectName || !initialized.current) {
      lastTargetPosition.current.copy(currentTargetPosition);
      prevSelectedObjectName.current = selectedObjectName;
      initialized.current = true;
      
      if (selectedObjectName && viewMode === 'orbit') {
        const radius = (targetObject.userData && targetObject.userData.radius) ? targetObject.userData.radius : 1;
        controls.fitToBox(targetObject, true, { paddingTop: radius * 3, paddingBottom: radius * 3, paddingLeft: radius * 3, paddingRight: radius * 3 });
      }
      return;
    }

    if (viewMode === 'orbit') {
      if (selectedObjectName) {
        const displacement = currentTargetPosition.clone().sub(lastTargetPosition.current);
        if (displacement.lengthSq() > 0) {
          const currentCamPos = new THREE.Vector3();
          const currentTarget = new THREE.Vector3();
          controls.getPosition(currentCamPos);
          controls.getTarget(currentTarget);
          controls.setLookAt(
            currentCamPos.x + displacement.x, currentCamPos.y + displacement.y, currentCamPos.z + displacement.z,
            currentTarget.x + displacement.x, currentTarget.y + displacement.y, currentTarget.z + displacement.z,
            false
          );
        }
        setCameraTarget(currentTargetPosition);
      }
    } else if (viewMode === 'surface') {
      const radius = (targetObject.userData && targetObject.userData.radius) ? targetObject.userData.radius : 1;
      const surfacePosition = currentTargetPosition.clone().add(new THREE.Vector3(0, radius * 1.2, 0));
      let lookAtTarget = new THREE.Vector3();
      if (surfaceTargetName) {
        const tObj = scene.getObjectByName(surfaceTargetName);
        if (tObj) tObj.getWorldPosition(lookAtTarget);
        else lookAtTarget.copy(surfacePosition).add(new THREE.Vector3(0, radius * 10, 0));
      } else {
        lookAtTarget.copy(surfacePosition).add(new THREE.Vector3(0, radius * 10, 0));
      }
      controls.setLookAt(surfacePosition.x, surfacePosition.y, surfacePosition.z, lookAtTarget.x, lookAtTarget.y, lookAtTarget.z, false);
    } else if (viewMode === 'galactic') {
      // FRONT VIEW: Place camera ahead of the sun in the direction of velocity
      const velocityDir = SUN_GALACTIC_VELOCITY.clone().normalize();
      const aheadPosition = sunPosition.clone().add(velocityDir.multiplyScalar(zoomDistance));
      
      // Look back at the sun
      controls.setLookAt(
        aheadPosition.x, aheadPosition.y, aheadPosition.z,
        sunPosition.x, sunPosition.y, sunPosition.z,
        false
      );
    }

    lastTargetPosition.current.copy(currentTargetPosition);
  });

  useEffect(() => {
    if (!cameraControlsRef.current) return;
    const controls = cameraControlsRef.current;
    if (viewMode !== 'galactic' && Math.abs(controls.distance - zoomDistance) > 0.01) {
      isInternalUpdate.current = true;
      controls.dollyTo(zoomDistance, false);
      requestAnimationFrame(() => { isInternalUpdate.current = false; });
    }
  }, [zoomDistance, viewMode]);

  useEffect(() => {
    if (!cameraControlsRef.current) return;
    const controls = cameraControlsRef.current;
    if (viewMode === 'galactic') {
        // Initial setup for galactic view
        controls.setLookAt(0, 0, 5000, 0, 0, 0, true);
    } else if (!selectedObjectName && viewMode === 'orbit') {
      controls.setLookAt(0, 1500, 4000, 0, 0, 0, true);
    }
  }, [selectedObjectName, viewMode, scene]);

  return (
    <CameraControls
      ref={cameraControlsRef}
      minDistance={0.001}
      maxDistance={2000000}
      dollyToCursor={true}
      dollySpeed={3.0}
      truckSpeed={0}
      mouseButtons={{
        left: viewMode === 'galactic' ? 0 : 1, // Disable rotation in galactic view to maintain focus
        middle: 8,
        right: 0,
        wheel: 8,
      }}
    />
  );
};

export default CameraController;