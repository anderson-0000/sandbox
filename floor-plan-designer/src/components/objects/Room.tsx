import React from 'react';
import { Rect, Text, Group } from 'react-konva';
import type { ProjectObject } from '../../types';
import { mmToPx } from '../../utils/unitConverter';
import { snapToPrecision } from '../../utils/snapping';
import { useProjectStore } from '../../store/useProjectStore';

interface RoomProps {
  object: ProjectObject;
}

const Room: React.FC<RoomProps> = ({ object }) => {
  const { projectSettings, updateObject, setSelectedId, selectedId, objects } = useProjectStore();
  const gridSizeMm = projectSettings.gridSizeMm;
  const wallThicknessMm = projectSettings.wallThicknessMm;
  const isSelected = selectedId === object.id;

  const SNAP_THRESHOLD = 200; // 200mm within other objects to snap

  // For rooms, snap to half-module (e.g., 455mm) for more flexibility
  const SNAP_UNIT = gridSizeMm / 2;

  const findSnapPosition = (x: number, y: number, width: number, height: number) => {
    let snappedX = x;
    let snappedY = y;

    for (const other of objects) {
      if (other.id === object.id) continue;
      const ox = other.position.x;
      const oy = other.position.y;
      const ow = other.dimensions.widthMm;
      const oh = other.dimensions.heightMm;

      if (Math.abs(x - (ox + ow)) < SNAP_THRESHOLD) snappedX = ox + ow;
      else if (Math.abs((x + width) - ox) < SNAP_THRESHOLD) snappedX = ox - width;
      else if (Math.abs(x - ox) < SNAP_THRESHOLD) snappedX = ox;

      if (Math.abs(y - (oy + oh)) < SNAP_THRESHOLD) snappedY = oy + oh;
      else if (Math.abs((y + height) - oy) < SNAP_THRESHOLD) snappedY = oy - height;
      else if (Math.abs(y - oy) < SNAP_THRESHOLD) snappedY = oy;
    }
    return { x: snappedX, y: snappedY };
  };

  const handleDragMove = (e: any) => {
    const node = e.target;
    const rawX = node.x() / 0.1;
    const rawY = node.y() / 0.1;
    const { x, y } = findSnapPosition(rawX, rawY, object.dimensions.widthMm, object.dimensions.heightMm);
    const finalX = snapToPrecision(x, SNAP_UNIT);
    const finalY = snapToPrecision(y, SNAP_UNIT);
    node.position({ x: mmToPx(finalX), y: mmToPx(finalY) });
  };

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

    // Snap to SNAP_UNIT, but allow at least 100mm
    const newWidthMm = Math.max(100, snapToPrecision((node.width() * scaleX) / 0.1, SNAP_UNIT));
    const newHeightMm = Math.max(100, snapToPrecision((node.height() * scaleY) / 0.1, SNAP_UNIT));
    const newXMm = snapToPrecision(node.x() / 0.1, SNAP_UNIT);
    const newYMm = snapToPrecision(node.y() / 0.1, SNAP_UNIT);

    updateObject(object.id, {
      position: { x: newXMm, y: newYMm },
      dimensions: { widthMm: newWidthMm, heightMm: newHeightMm },
      rotation: Math.round(node.rotation()),
    });
  };

  const width = mmToPx(object.dimensions.widthMm);
  const height = mmToPx(object.dimensions.heightMm);
  const wallPx = mmToPx(wallThicknessMm);

  // Check neighbors to avoid double walls, but keep one wall (based on ID priority)
  const shouldHideWall = (side: 'top' | 'bottom' | 'left' | 'right') => {
    const x = object.position.x;
    const y = object.position.y;
    const w = object.dimensions.widthMm;
    const h = object.dimensions.heightMm;

    return objects.some(other => {
      if (other.id === object.id || other.type !== 'room') return false;
      
      // ID priority: higher ID hides its wall if adjacent to a lower ID
      if (object.id < other.id) return false;

      const ox = other.position.x;
      const oy = other.position.y;
      const ow = other.dimensions.widthMm;
      const oh = other.dimensions.heightMm;

      switch(side) {
        case 'left': return Math.abs(x - (ox + ow)) < 10 && Math.max(y, oy) < Math.min(y + h, oy + oh);
        case 'right': return Math.abs((x + w) - ox) < 10 && Math.max(y, oy) < Math.min(y + h, oy + oh);
        case 'top': return Math.abs(y - (oy + oh)) < 10 && Math.max(x, ox) < Math.min(x + w, ox + ow);
        case 'bottom': return Math.abs((y + h) - oy) < 10 && Math.max(x, ox) < Math.min(x + w, ox + ow);
      }
      return false;
    });
  };

  return (
    <Group
      id={object.id}
      x={mmToPx(object.position.x)}
      y={mmToPx(object.position.y)}
      width={width}
      height={height}
      draggable
      onDragMove={handleDragMove}
      onDragEnd={handleDragEnd}
      onTransformEnd={handleTransformEnd}
      onClick={() => setSelectedId(object.id)}
      onTap={() => setSelectedId(object.id)}
      rotation={object.rotation}
    >
      {/* Dynamic Wall Rendering */}
      <Rect
        x={shouldHideWall('left') ? 0 : -wallPx}
        y={shouldHideWall('top') ? 0 : -wallPx}
        width={width + (shouldHideWall('left') ? 0 : wallPx) + (shouldHideWall('right') ? 0 : wallPx)}
        height={height + (shouldHideWall('top') ? 0 : wallPx) + (shouldHideWall('bottom') ? 0 : wallPx)}
        fill="#cbd5e1"
        stroke="#94a3b8"
        strokeWidth={0.5}
      />
      <Rect
        width={width}
        height={height}
        fill={object.properties.color || '#f0f8ff'}
        stroke={isSelected ? '#3b82f6' : '#cbd5e1'}
        strokeWidth={isSelected ? 2 : 1}
      />
      <Text text={`${object.name} (${(((object.dimensions.widthMm * object.dimensions.heightMm) / 1000000) / 1.62).toFixed(1)}畳)`} width={width} height={height} align="center" verticalAlign="middle" fontSize={11} fill="#475569" />
      <Text text={`${(object.dimensions.widthMm / 1000).toFixed(2)}m x ${(object.dimensions.heightMm / 1000).toFixed(2)}m (${((object.dimensions.widthMm * object.dimensions.heightMm) / 1000000).toFixed(2)}㎡)`} width={width} y={height - 14} align="center" fontSize={8} fill="#94a3b8" />
    </Group>
  );
};

export default Room;
