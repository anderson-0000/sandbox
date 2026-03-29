export type ModuleType = 'shaku' | 'meter';

export interface Position {
  x: number;
  y: number;
}

export interface Dimensions {
  widthMm: number;
  heightMm: number;
}

export type ObjectType = 'room' | 'garage' | 'vehicle' | 'equipment' | 'stairs' | 'entrance' | 'door' | 'window' | 'bath' | 'kitchen';

export interface ProjectObject {
  id: string;
  type: ObjectType;
  name: string;
  dimensions: Dimensions;
  position: Position;
  rotation: number;
  properties: {
    color?: string;
    floorLevel?: number;
    [key: string]: any;
  };
}

export interface Plot {
  points: Position[];
}

export interface ProjectSettings {
  module: ModuleType;
  gridSizeMm: number;
  wallThicknessMm: number;
}

export interface ProjectData {
  version: string;
  projectSettings: ProjectSettings;
  plot: Plot;
  objects: ProjectObject[];
}
