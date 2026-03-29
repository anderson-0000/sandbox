import React, { useRef, useEffect, useState } from 'react';
import { Stage, Layer, Line, Transformer, Circle, Text, Group } from 'react-konva';
import { useProjectStore } from '../../store/useProjectStore';
import Grid from './Grid';
import Room from '../objects/Room';
import Vehicle from '../objects/Vehicle';
import Projector from '../objects/Projector';
import ArchitecturalElement from '../objects/ArchitecturalElement';
import { mmToPx, formatArea } from '../../utils/unitConverter';
import { snapToGrid } from '../../utils/snapping';
import { calculatePolygonAreaMm2, mm2ToM2 } from '../../utils/areaCalc';

const ProjectStage: React.FC = () => {
  const containerRef = useRef<HTMLDivElement>(null);
  const stageRef = useRef<any>(null);
  const transformerRef = useRef<any>(null);
  const [size, setSize] = useState({ width: 0, height: 0 });
  
  const { 
    objects, projectSettings, plot, setSelectedId, selectedId,
    stageScale, setStageScale, stagePos, setStagePos, updatePlotPoint
  } = useProjectStore();

  useEffect(() => {
    const updateSize = () => {
      if (containerRef.current) {
        setSize({
          width: containerRef.current.offsetWidth,
          height: containerRef.current.offsetHeight,
        });
      }
    };
    updateSize();
    window.addEventListener('resize', updateSize);
    return () => window.removeEventListener('resize', updateSize);
  }, []);

  useEffect(() => {
    if (transformerRef.current && selectedId) {
      const selectedNode = stageRef.current.findOne(`#${selectedId}`);
      if (selectedNode) {
        transformerRef.current.nodes([selectedNode]);
        transformerRef.current.getLayer().batchDraw();
      } else {
        transformerRef.current.nodes([]);
      }
    } else if (transformerRef.current) {
      transformerRef.current.nodes([]);
    }
  }, [selectedId, objects]);

  const handleWheel = (e: any) => {
    e.evt.preventDefault();
    const scaleBy = 1.1;
    const stage = stageRef.current;
    const oldScale = stage.scaleX();
    const mousePointTo = {
      x: stage.getPointerPosition().x / oldScale - stage.x() / oldScale,
      y: stage.getPointerPosition().y / oldScale - stage.y() / oldScale,
    };
    const newScale = e.evt.deltaY < 0 ? oldScale * scaleBy : oldScale / scaleBy;
    setStageScale(newScale);
    setStagePos({
      x: (stage.getPointerPosition().x / newScale - mousePointTo.x) * newScale,
      y: (stage.getPointerPosition().y / newScale - mousePointTo.y) * newScale,
    });
  };

  const plotPoints = plot.points.flatMap(p => [mmToPx(p.x), mmToPx(p.y)]);
  const plotAreaM2 = mm2ToM2(calculatePolygonAreaMm2(plot.points));

  // Calculate plot center for area label
  const plotCenter = plot.points.reduce(
    (acc, p) => ({ x: acc.x + p.x / plot.points.length, y: acc.y + p.y / plot.points.length }),
    { x: 0, y: 0 }
  );

  return (
    <div ref={containerRef} className="flex-1 bg-slate-100 overflow-hidden relative cursor-crosshair">
      <Stage
        ref={stageRef}
        width={size.width}
        height={size.height}
        scaleX={stageScale}
        scaleY={stageScale}
        x={stagePos.x}
        y={stagePos.y}
        draggable
        onWheel={handleWheel}
        onDragEnd={(e) => {
          if (e.target === stageRef.current) setStagePos({ x: e.target.x(), y: e.target.y() });
        }}
        onClick={(e) => { if (e.target === e.target.getStage()) setSelectedId(null); }}
      >
        <Layer>
          <Grid gridSizeMm={projectSettings.gridSizeMm} width={50000} height={50000} />

          {/* Site Plot */}
          <Line points={plotPoints} closed stroke="#94a3b8" strokeWidth={2} fill="#ffffff" />
          
          {/* Plot Labels (Dimensions) */}
          {plot.points.map((p, i) => {
            const nextP = plot.points[(i + 1) % plot.points.length];
            const distMm = Math.sqrt(Math.pow(nextP.x - p.x, 2) + Math.pow(nextP.y - p.y, 2));
            const midX = mmToPx((p.x + nextP.x) / 2);
            const midY = mmToPx((p.y + nextP.y) / 2);
            return (
              <Text
                key={`dim-${i}`}
                x={midX}
                y={midY}
                text={`${(distMm / 1000).toFixed(2)}m`}
                fontSize={10 / stageScale}
                fill="#64748b"
                align="center"
                offsetX={20 / stageScale}
              />
            );
          })}

          {/* Plot Area Label */}
          <Group x={mmToPx(plotCenter.x)} y={mmToPx(plotCenter.y)}>
            <Text
              text={formatArea(plotAreaM2)}
              fontSize={14 / stageScale}
              fontStyle="bold"
              fill="#94a3b8"
              align="center"
              offsetX={50 / stageScale}
            />
            <Text
              text="敷地面積"
              y={-15 / stageScale}
              fontSize={10 / stageScale}
              fill="#cbd5e1"
              align="center"
              offsetX={20 / stageScale}
            />
          </Group>

          {/* Plot Edit Handles */}
          {plot.points.map((p, i) => (
            <Circle
              key={`plot-pt-${i}`}
              x={mmToPx(p.x)}
              y={mmToPx(p.y)}
              radius={6 / stageScale}
              fill="#94a3b8"
              stroke="#fff"
              strokeWidth={2 / stageScale}
              draggable
              onDragMove={(e) => {
                const x = Math.round(e.target.x() / 0.1);
                const y = Math.round(e.target.y() / 0.1);
                updatePlotPoint(i, { x: snapToGrid(x, projectSettings.gridSizeMm), y: snapToGrid(y, projectSettings.gridSizeMm) });
              }}
            />
          ))}

          {/* Objects */}
          {objects.map((obj) => {
            if (obj.type === 'vehicle') return <Vehicle key={obj.id} object={obj} />;
            if (obj.type === 'equipment') return <Projector key={obj.id} object={obj} />;
            if (['stairs', 'kitchen', 'bath', 'entrance', 'door', 'window'].includes(obj.type)) {
              return <ArchitecturalElement key={obj.id} object={obj} />;
            }
            return <Room key={obj.id} object={obj} />;
          })}

          <Transformer
            ref={transformerRef}
            boundBoxFunc={(oldBox, newBox) => {
              const minSize = 5; // Allow smaller but not zero
              if (newBox.width < minSize || newBox.height < minSize) return oldBox;
              return newBox;
            }}
            rotateEnabled={true}
            anchorSize={8 / stageScale}
            padding={5 / stageScale}
          />
        </Layer>
      </Stage>

      <div className="absolute bottom-6 left-6 flex flex-col gap-2 z-50">
        <button onClick={() => setStageScale(stageScale * 1.2)} className="w-10 h-10 bg-white shadow-md border border-slate-200 rounded-full flex items-center justify-center font-bold text-slate-600 hover:bg-slate-50">+</button>
        <button onClick={() => setStageScale(stageScale / 1.2)} className="w-10 h-10 bg-white shadow-md border border-slate-200 rounded-full flex items-center justify-center font-bold text-slate-600 hover:bg-slate-50">-</button>
      </div>
    </div>
  );
};

export default ProjectStage;
