import { useState, useEffect, useCallback } from 'react';
// @ts-ignore
import * as satellite from 'satellite.js';
import axios from 'axios';

export interface DebrisObject {
  id: string;
  name: string;
  satrec: any;
  position: [number, number, number];
  isCollected: boolean;
}

export const useDebris = () => {
  const [debris, setDebris] = useState<DebrisObject[]>([]);
  const [score, setScore] = useState(0);

  const fetchDebris = async () => {
    try {
      // 確実にキャッシュからデータを取得
      const response = await axios.get('/api/celestrak/GROUP=active&FORMAT=tle');
      // 改行コード \r\n と \n の両方に対応
      const lines = response.data.replace(/\r/g, '').split('\n').filter((l: string) => l.trim() !== '');
      const newDebris: DebrisObject[] = [];

      // 表示するデブリの数を 500個に増やす
      for (let i = 0; i < lines.length && newDebris.length < 500; i += 3) {
        const name = lines[i].trim();
        const tle1 = lines[i + 1];
        const tle2 = lines[i + 2];

        if (tle1 && tle2 && tle1.startsWith('1') && tle2.startsWith('2')) {
          try {
            const satrec = satellite.twoline2satrec(tle1, tle2);
            newDebris.push({
              id: name + i,
              name,
              satrec,
              position: [0, 0, 0],
              isCollected: false,
            });
          } catch (e) {
            // スキップ
          }
        }
      }
      console.log(`Successfully loaded ${newDebris.length} debris objects from cache.`);
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
          // 地球の半径 6371km を基準にスケール (1 unit = 6371km)
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

  const collectDebris = useCallback((id: string) => {
    setDebris((prev) =>
      prev.map((d) => {
        if (d.id === id && !d.isCollected) {
          setScore((s) => s + 100);
          return { ...d, isCollected: true };
        }
        return d;
      })
    );
  }, []);

  return { debris, score, updatePositions, collectDebris };
};
