import { useState, useEffect, useCallback } from 'react';
import * as satellite from 'satellite.js';
import axios from 'axios';

export interface DebrisObject {
  id: string;
  name: string;
  satrec: satellite.SatRec;
  position: [number, number, number];
  isCollected: boolean;
}

export const useDebris = () => {
  const [debris, setDebris] = useState<DebrisObject[]>([]);
  const [score, setScore] = useState(0);

  const fetchDebris = async () => {
    try {
      // 例として Cosmos 2251 のデブリを取得
      const response = await axios.get('/api/celestrak/GROUP=cosmos-2251-debris&FORMAT=tle');
      const lines = response.data.trim().split('\n');
      const newDebris: DebrisObject[] = [];

      for (let i = 0; i < lines.length; i += 3) {
        const name = lines[i].trim();
        const tle1 = lines[i + 1];
        const tle2 = lines[i + 2];

        if (tle1 && tle2) {
          const satrec = satellite.twoline2satrec(tle1, tle2);
          newDebris.push({
            id: name + i,
            name,
            satrec,
            position: [0, 0, 0],
            isCollected: false,
          });
        }
      }
      setDebris(newDebris);
    } catch (error) {
      console.error('Failed to fetch debris:', error);
    }
  };

  useEffect(() => {
    fetchDebris();
  }, []);

  const updatePositions = useCallback(() => {
    const now = new Date();
    setDebris((prev) =>
      prev.map((d) => {
        if (d.isCollected) return d;

        const positionAndVelocity = satellite.propagate(d.satrec, now);
        if (!positionAndVelocity || !positionAndVelocity.position) return d;
        const positionEci = positionAndVelocity.position;

        if (typeof positionEci !== 'boolean') {
          // ECI to Three.js coordinates (Scaling: 1 unit = 100km is common, but let's use 1 unit = 1000km for visibility)
          // Earth radius is ~6371km. If Earth is scale 1, then x/6371.
          const scale = 1 / 6371;
          return {
            ...d,
            position: [
              positionEci.x * scale,
              positionEci.z * scale,
              -positionEci.y * scale,
            ] as [number, number, number],
          };
        }
        return d;
      })
    );
  }, []);

  const collectDebris = (id: string) => {
    setDebris((prev) =>
      prev.map((d) => {
        if (d.id === id && !d.isCollected) {
          setScore((s) => s + 100);
          return { ...d, isCollected: true };
        }
        return d;
      })
    );
  };

  return { debris, score, updatePositions, collectDebris };
};
