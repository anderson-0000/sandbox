export const SCALE = 0.1; // 1mm = 0.1px (1px = 10mm)

export const mmToPx = (mm: number): number => mm * SCALE;
export const pxToMm = (px: number): number => px / SCALE;

export const m2ToTsubo = (m2: number): number => m2 * 0.3025;
export const m2ToJo = (m2: number): number => m2 / 1.62;

export const formatArea = (m2: number): string => {
  const tsubo = m2ToTsubo(m2);
  const jo = m2ToJo(m2);
  return `${m2.toFixed(2)} ㎡ (${jo.toFixed(1)} 畳 / ${tsubo.toFixed(2)} 坪)`;
};
