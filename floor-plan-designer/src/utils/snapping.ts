export const snapToPrecision = (value: number, precision: number): number => {
  return Math.round(value / precision) * precision;
};

export const snapToGrid = (value: number, gridSize: number): number => {
  return snapToPrecision(value, gridSize);
};

export const snapPosition = (pos: { x: number, y: number }, gridSizeMm: number) => {
  return {
    x: snapToPrecision(pos.x, gridSizeMm),
    y: snapToPrecision(pos.y, gridSizeMm),
  };
};
