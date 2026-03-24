import { useEffect, useRef, useMemo, useState } from 'react';
import { useFrame, useThree } from '@react-three/fiber';
import { InstancedMesh, Object3D, Color } from 'three';
import { fetchActiveSatellites, type SatelliteData } from '../services/api';
import { getSatellitePosition } from '../utils/tle';

const tempObject = new Object3D();
const tempColor = new Color();

const COLORS: Record<string, string> = {
  iss: '#ffcc00',             // Gold
  telescope: '#ff00ff',        // Magenta
  science_mission: '#ff8800', // Orange (NASA/JAXA)
  active: '#00ff00',           // Green
  starlink: '#00ffff',         // Cyan
  debris: '#ff0000',           // Red
  weather: '#ffff00',          // Yellow
  other: '#ffffff'             // White
};

interface SatelliteCloudProps {
  onSelect: (sat: SatelliteData | null, altitude: number) => void;
  selectedSat: SatelliteData | null;
  onDataLoaded: (count: number) => void;
}

export default function SatelliteCloud({ onSelect, selectedSat, onDataLoaded }: SatelliteCloudProps) {
  const meshRef = useRef<InstancedMesh>(null!);
  const [satellites, setSatellites] = useState<SatelliteData[]>([]);
  const [hovered, setHovered] = useState<number | null>(null);
  
  const selectedIndex = useMemo(() => {
    if (!selectedSat) return -1;
    return satellites.findIndex(s => s.catId === selectedSat.catId);
  }, [selectedSat, satellites]);

  useEffect(() => {
    async function loadData() {
      console.log("Requesting consolidated active satellite data...");
      try {
        const sats = await fetchActiveSatellites();
        console.log(`Loaded ${sats.length} satellites into the scene.`);
        setSatellites(sats);
        onDataLoaded(sats.length);
      } catch (e) {
        console.error("Failed to load data", e);
        onDataLoaded(0);
      }
    }
    loadData();
  }, []);

  const colorArray = useMemo(() => {
    const array = new Float32Array(satellites.length * 3);
    satellites.forEach((sat, i) => {
      const colorHex = COLORS[sat.type] || COLORS.other;
      tempColor.set(colorHex);
      tempColor.toArray(array, i * 3);
    });
    return array;
  }, [satellites]);

  const colorAttribRef = useRef<any>(null!);
  
  useEffect(() => {
    if (colorAttribRef.current) {
      colorAttribRef.current.needsUpdate = true;
    }
  }, [colorArray]);

  useFrame(() => {
    if (!meshRef.current || satellites.length === 0) return;
    
    const date = new Date();

    satellites.forEach((sat, i) => {
      const pos = getSatellitePosition(sat.line1, sat.line2, date);
      if (pos) {
        tempObject.position.copy(pos);
        
        // Dynamic scale
        let baseScale = 80; 
        if (sat.type === 'iss' || sat.type === 'telescope' || sat.type === 'science_mission') {
          baseScale = 250; 
        }
        
        let finalScale = baseScale;
        if (selectedIndex === i) {
          finalScale = baseScale * 4.0;
        } else if (hovered === i) {
          finalScale = baseScale * 2.2;
        }

        tempObject.scale.setScalar(finalScale); 
        tempObject.updateMatrix();
        meshRef.current.setMatrixAt(i, tempObject.matrix);

        if (selectedIndex === i) {
          tempColor.set('#ffffff');
          meshRef.current.setColorAt(i, tempColor);
        } else {
          const colorHex = COLORS[sat.type] || COLORS.other;
          tempColor.set(colorHex);
          meshRef.current.setColorAt(i, tempColor);
        }
      } else {
        tempObject.scale.setScalar(0);
        tempObject.updateMatrix();
        meshRef.current.setMatrixAt(i, tempObject.matrix);
      }
    });
    
    meshRef.current.instanceMatrix.needsUpdate = true;
    if (meshRef.current.instanceColor) {
      meshRef.current.instanceColor.needsUpdate = true;
    }
  });

  useEffect(() => {
    if (meshRef.current) {
      meshRef.current.geometry.computeBoundingSphere();
      if (meshRef.current.geometry.boundingSphere) {
        meshRef.current.geometry.boundingSphere.radius = 2000000;
      }
    }
  }, [satellites.length]);

  const { raycaster, mouse, camera } = useThree();

  const handlePointerClick = (e: any) => {
    e.stopPropagation();
    raycaster.setFromCamera(mouse, camera);
    raycaster.params.Mesh = { threshold: 500 }; 
    
    const intersects = raycaster.intersectObject(meshRef.current);
    
    if (intersects.length > 0) {
      intersects.sort((a, b) => a.distance - b.distance);
      const instanceId = intersects[0].instanceId;
      if (instanceId !== undefined && satellites[instanceId]) {
        const sat = satellites[instanceId];
        const pos = getSatellitePosition(sat.line1, sat.line2, new Date());
        const alt = pos ? pos.length() - 6371 : 0;
        onSelect(sat, alt);
      }
    }
  };

  if (satellites.length === 0) return null;

  return (
    <instancedMesh
      key={`sats-mesh-${satellites.length}`}
      ref={meshRef}
      args={[undefined, undefined, satellites.length]}
      onClick={handlePointerClick}
      onPointerDown={(e) => e.stopPropagation()}
      frustumCulled={false}
      onPointerOver={(e) => {
          if (e.instanceId !== undefined) {
            document.body.style.cursor = 'pointer';
            setHovered(e.instanceId);
          }
      }}
      onPointerOut={() => {
          document.body.style.cursor = 'default';
          setHovered(null);
      }}
    >
      <sphereGeometry args={[1, 12, 12]} />
      <meshBasicMaterial toneMapped={false} transparent opacity={0.9} />
      {satellites.length > 0 && (
        <instancedBufferAttribute
          ref={colorAttribRef}
          attach="instanceColor"
          args={[colorArray, 3]}
        />
      )}
    </instancedMesh>
  );
}
