import * as satellite from 'satellite.js';
import { Vector3 } from 'three';

export function getSatellitePosition(line1: string, line2: string, date: Date): Vector3 | null {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const satrec = (satellite as any).twoline2satrec(line1, line2);
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const positionAndVelocity = (satellite as any).propagate(satrec, date);
  const positionEci = positionAndVelocity.position;

  if (!positionEci || typeof positionEci !== 'object') {
    return null;
  }

  // Handle potential failures in propagation
  const x = positionEci.x;
  const y = positionEci.y;
  const z = positionEci.z;
  
  if (typeof x !== 'number' || typeof y !== 'number' || typeof z !== 'number') {
      return null;
  }

  // Mapping: ECI to Three.js (Y-up)
  // ECI X -> Three X
  // ECI Z -> Three Y (North)
  // ECI Y -> Three -Z
  return new Vector3(x, z, -y);
}
