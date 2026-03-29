import React from 'react';
import { Rect, Text, Group, Line } from 'react-konva';
import type { ProjectObject } from '../../types';
import { mmToPx } from '../../utils/unitConverter';
import { useProjectStore } from '../../store/useProjectStore';

interface VehicleProps {
  object: ProjectObject;
}

const Vehicle: React.FC<VehicleProps> = ({ object }) => {
  const { updateObject, setSelectedId, selectedId } = useProjectStore();
  const isSelected = selectedId === object.id;

  const handleDragEnd = (e: any) => {
    const node = e.target;
    updateObject(object.id, {
      position: { x: Math.round(node.x() / 0.1), y: Math.round(node.y() / 0.1) }
    });
  };

  const handleTransformEnd = (e: any) => {
    const node = e.target;
    const scaleX = node.scaleX();
    const scaleY = node.scaleY();
    node.scaleX(1);
    node.scaleY(1);
    const newWidthMm = Math.max(500, Math.round((node.width() * scaleX) / 0.1));
    const newHeightMm = Math.max(500, Math.round((node.height() * scaleY) / 0.1));
    updateObject(object.id, {
      dimensions: { 
        widthMm: newWidthMm, 
        heightMm: newHeightMm 
      },
      rotation: Math.round(node.rotation()),
    });
  };

  const width = mmToPx(object.dimensions.widthMm);
  const height = mmToPx(object.dimensions.heightMm);
  const doorWidth = mmToPx(800);

  return (
    <Group
      id={object.id}
      x={mmToPx(object.position.x)}
      y={mmToPx(object.position.y)}
      width={width}
      height={height}
      draggable
      onDragEnd={handleDragEnd}
      onTransformEnd={handleTransformEnd}
      onClick={() => setSelectedId(object.id)}
      onTap={() => setSelectedId(object.id)}
      rotation={object.rotation}
    >
      <Rect
        width={width}
        height={height}
        fill="#cbd5e1"
        stroke={isSelected ? '#3b82f6' : '#94a3b8'}
        strokeWidth={isSelected ? 2 : 1}
        cornerRadius={width * 0.15}
      />
      <Line
        points={[-doorWidth, height * 0.3, 0, height * 0.3]}
        stroke="#ef4444"
        strokeWidth={1}
        dash={[4, 4]}
        opacity={0.5}
      />
      <Line
        points={[width, height * 0.3, width + doorWidth, height * 0.3]}
        stroke="#ef4444"
        strokeWidth={1}
        dash={[4, 4]}
        opacity={0.5}
      />
      <Text
        text={object.name}
        width={width}
        height={height}
        align="center"
        verticalAlign="middle"
        fontSize={10}
        fill="#475569"
      />
    </Group>
  );
};

export default Vehicle;
