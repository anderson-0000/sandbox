import React from 'react';
import { Line, Group } from 'react-konva';
import { mmToPx } from '../../utils/unitConverter';

interface GridProps {
  gridSizeMm: number;
  width: number;
  height: number;
}

const Grid: React.FC<GridProps> = ({ gridSizeMm, width, height }) => {
  const gridSizePx = mmToPx(gridSizeMm);
  const lines = [];

  // Vertical lines
  for (let i = 0; i <= width / gridSizePx; i++) {
    lines.push(
      <Line
        key={`v-${i}`}
        points={[i * gridSizePx, 0, i * gridSizePx, height]}
        stroke="#e2e8f0"
        strokeWidth={1}
      />
    );
  }

  // Horizontal lines
  for (let j = 0; j <= height / gridSizePx; j++) {
    lines.push(
      <Line
        key={`h-${j}`}
        points={[0, j * gridSizePx, width, j * gridSizePx]}
        stroke="#e2e8f0"
        strokeWidth={1}
      />
    );
  }

  return <Group listening={false}>{lines}</Group>;
};

export default Grid;
