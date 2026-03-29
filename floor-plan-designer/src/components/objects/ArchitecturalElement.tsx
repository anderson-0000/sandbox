import React from 'react';
import { Rect, Text, Group, Line, Arc, Circle as KonvaCircle } from 'react-konva';
import type { ProjectObject } from '../../types';
import { mmToPx } from '../../utils/unitConverter';
import { snapToPrecision } from '../../utils/snapping';
import { useProjectStore } from '../../store/useProjectStore';

interface ArchitecturalElementProps {
  object: ProjectObject;
}

const ArchitecturalElement: React.FC<ArchitecturalElementProps> = ({ object }) => {
  const { updateObject, setSelectedId, selectedId } = useProjectStore();
  const isSelected = selectedId === object.id;

  const SNAP_UNIT = 10; 

  const handleDragEnd = (e: any) => {
    const node = e.target;
    const snappedX = snapToPrecision(node.x() / 0.1, SNAP_UNIT);
    const snappedY = snapToPrecision(node.y() / 0.1, SNAP_UNIT);
    updateObject(object.id, { position: { x: snappedX, y: snappedY } });
    node.position({ x: mmToPx(snappedX), y: mmToPx(snappedY) });
  };

  const handleTransformEnd = (e: any) => {
    const node = e.target;
    const scaleX = node.scaleX();
    const scaleY = node.scaleY();
    node.scaleX(1);
    node.scaleY(1);
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

  const renderInternal = () => {
    switch (object.type) {
      case 'stairs':
        const steps = 8;
        const stepLines = [];
        for (let i = 1; i < steps; i++) {
          stepLines.push(<Line key={i} points={[0, (height / steps) * i, width, (height / steps) * i]} stroke="#cbd5e1" strokeWidth={1} />);
        }
        return <>{stepLines}</>;
      case 'kitchen':
        return (
          <>
            <Rect x={width * 0.05} y={height * 0.1} width={width * 0.9} height={height * 0.8} stroke="#cbd5e1" strokeWidth={1} />
            <Rect x={width * 0.1} y={height * 0.25} width={width * 0.3} height={height * 0.5} stroke="#cbd5e1" strokeWidth={1} cornerRadius={2} />
            <KonvaCircle x={width * 0.7} y={height * 0.5} radius={height * 0.2} stroke="#cbd5e1" strokeWidth={1} />
            <KonvaCircle x={width * 0.85} y={height * 0.5} radius={height * 0.15} stroke="#cbd5e1" strokeWidth={1} />
          </>
        );
      case 'bath':
        const tubWidth = width * 0.5;
        return (
          <>
            {/* Tub Area */}
            <Rect x={width * 0.05} y={height * 0.05} width={tubWidth} height={height * 0.9} cornerRadius={width * 0.05} stroke="#94a3b8" strokeWidth={1} />
            <Rect x={width * 0.1} y={height * 0.1} width={tubWidth * 0.8} height={height * 0.8} cornerRadius={width * 0.1} fill="rgba(224, 242, 254, 0.5)" stroke="#cbd5e1" strokeWidth={1} />
            
            {/* Washing Area (洗い場) */}
            <Rect x={tubWidth + width * 0.05} y={height * 0.05} width={width - tubWidth - width * 0.1} height={height * 0.9} stroke="#cbd5e1" strokeWidth={0.5} dash={[2, 2]} />
            <KonvaCircle x={tubWidth + (width - tubWidth) / 2} y={height * 0.2} radius={mmToPx(50)} fill="#cbd5e1" /> {/* Drain */}
            <Rect x={tubWidth + (width - tubWidth) / 2 - width * 0.1} y={height * 0.7} width={width * 0.2} height={height * 0.1} fill="#e2e8f0" cornerRadius={2} /> {/* Counter */}
          </>
        );
      case 'entrance':
        return <Line points={[0, 0, width, height]} stroke="#cbd5e1" strokeWidth={1} dash={[2, 2]} />;
      case 'door':
        return <Arc x={0} y={height} innerRadius={0} outerRadius={width} angle={90} fill="rgba(180, 83, 9, 0.1)" stroke="#b45309" strokeWidth={1} rotation={-90} />;
      case 'window':
        return (
          <>
            <Rect y={height * 0.4} width={width} height={height * 0.2} stroke="#38bdf8" strokeWidth={1} />
            <Line points={[0, height / 2, width, height / 2]} stroke="#38bdf8" strokeWidth={1} />
          </>
        );
      default:
        return null;
    }
  };

  const areaM2 = (object.dimensions.widthMm * object.dimensions.heightMm) / 1000000;
  const jo = areaM2 / 1.62;

  return (
    <Group id={object.id} x={mmToPx(object.position.x)} y={mmToPx(object.position.y)} width={width} height={height} draggable onDragEnd={handleDragEnd} onTransformEnd={handleTransformEnd} onClick={() => setSelectedId(object.id)} onTap={() => setSelectedId(object.id)} rotation={object.rotation}>
      <Rect width={width} height={height} fill={object.properties.color || '#fff'} stroke={isSelected ? '#3b82f6' : '#cbd5e1'} strokeWidth={isSelected ? 2 : 1} />
      {renderInternal()}
      <Text text={`${object.name} (${jo.toFixed(1)}畳)`} width={width} height={height} align="center" verticalAlign="bottom" fontSize={8} fill="#64748b" padding={2} />
    </Group>
  );
};

export default ArchitecturalElement;
