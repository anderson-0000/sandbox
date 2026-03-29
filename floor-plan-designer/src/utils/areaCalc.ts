import type { ProjectObject, Position } from '../types';

/**
 * Calculates the area of a polygon using the Shoelace formula (in mm2)
 */
export const calculatePolygonAreaMm2 = (points: Position[]): number => {
  let area = 0;
  for (let i = 0; i < points.length; i++) {
    const p1 = points[i];
    const p2 = points[(i + 1) % points.length];
    area += (p1.x * p2.y) - (p2.x * p1.y);
  }
  return Math.abs(area) / 2;
};

/**
 * Converts mm2 to m2
 */
export const mm2ToM2 = (mm2: number): number => mm2 / 1000000;

/**
 * Calculates total floor area from objects (in m2)
 */
export const calculateFloorAreaM2 = (objects: ProjectObject[]): number => {
  const roomAreasMm2 = objects
    .filter(obj => obj.type === 'room')
    .reduce((sum, obj) => sum + (obj.dimensions.widthMm * obj.dimensions.heightMm), 0);
  return mm2ToM2(roomAreasMm2);
};

/**
 * Calculates building area (in m2)
 * For simplicity, we assume building area is the sum of ground floor objects and garages.
 */
export const calculateBuildingAreaM2 = (objects: ProjectObject[]): number => {
  const buildingObjectsMm2 = objects
    .filter(obj => (obj.type === 'room' && obj.properties.floorLevel === 1) || obj.type === 'garage')
    .reduce((sum, obj) => sum + (obj.dimensions.widthMm * obj.dimensions.heightMm), 0);
  return mm2ToM2(buildingObjectsMm2);
};
