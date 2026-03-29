import React from 'react';
import { Rect, Text, Group, Line } from 'react-konva';
import type { ProjectObject } from '../../types';
import { mmToPx } from '../../utils/unitConverter';
import { useProjectStore } from '../../store/useProjectStore';

interface ProjectorProps {
  object: ProjectObject;
}

const Projector: React.FC<ProjectorProps> = ({ object }) => {
  const { setSelectedId, selectedId, updateObject } = useProjectStore();
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
    const newWidthMm = Math.max(100, Math.round((node.width() * scaleX) / 0.1));
    const newHeightMm = Math.max(100, Math.round((node.height() * scaleY) / 0.1));
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
  const throwDistance = mmToPx(3000);
  const screenWidth = mmToPx(2200);

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
        fill="#334155"
        stroke={isSelected ? '#3b82f6' : '#1e293b'}
        strokeWidth={isSelected ? 2 : 1}
        cornerRadius={4}
      />
      {isSelected && (
        <Group listening={false}>
          <Line
            points={[width / 2, 0, width / 2 - screenWidth / 2, -throwDistance, width / 2 + screenWidth / 2, -throwDistance]}
            closed
            fill="rgba(59, 130, 246, 0.1)"
            stroke="#3b82f6"
            strokeWidth={1}
            dash={[5, 5]}
          />
        </Group>
      )}
      <Text
        text={object.name}
        y={height + 5}
        width={width}
        align="center"
        fontSize={10}
        fill="#475569"
      />
    </Group>
  );
};

export default Projector;
