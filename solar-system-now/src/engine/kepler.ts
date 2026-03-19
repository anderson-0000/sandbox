import * as THREE from 'three';

/**
 * Orbital elements at Epoch J2000 (roughly Year 2000 Jan 1.5)
 * Reference: NASA JPL Approximate Positions of the Planets
 * 
 * Rates are per century. For this prototype, we use the constants at J2000.
 */

export const SUN_RADIUS = 109.1; // Earth = 1.0

export interface OrbitalElements {
  id: string; // Internal ID (e.g. 'Earth')
  name: string; // Display name (e.g. '地球')
  a: number;   // AU
  e: number;
  i: number;   // deg
  L: number;   // deg
  lp: number;  // deg (longitude of perihelion)
  lan: number; // deg (longitude of ascending node)
  color: string;
  radius: number; // Relative to Earth (1.0)
}

export const PLANETS_DATA: Record<string, OrbitalElements> = {
  Mercury: { id: 'Mercury', name: '水星', a: 0.38709893, e: 0.20563069, i: 7.00487, L: 252.25084, lp: 77.45645, lan: 48.33167, color: '#A5A5A5', radius: 0.383 },
  Venus: { id: 'Venus', name: '金星', a: 0.72333199, e: 0.00677323, i: 3.39471, L: 181.97973, lp: 131.53298, lan: 76.68069, color: '#E3BB76', radius: 0.949 },
  Earth: { id: 'Earth', name: '地球', a: 1.00000011, e: 0.01671022, i: 0.00005, L: 100.46435, lp: 102.94719, lan: -11.26064, color: '#2271B3', radius: 1.0 },
  Mars: { id: 'Mars', name: '火星', a: 1.52366231, e: 0.09341233, i: 1.85061, L: 355.45332, lp: 336.04084, lan: 49.57854, color: '#E27B58', radius: 0.532 },
  Jupiter: { id: 'Jupiter', name: '木星', a: 5.20336301, e: 0.04839266, i: 1.30530, L: 34.40438, lp: 14.75385, lan: 100.55615, color: '#D39C7E', radius: 10.97 },
  Saturn: { id: 'Saturn', name: '土星', a: 9.53707032, e: 0.05415060, i: 2.48446, L: 49.94432, lp: 92.43194, lan: 113.71504, color: '#C5AB6E', radius: 9.14 },
  Uranus: { id: 'Uranus', name: '天王星', a: 19.19126393, e: 0.04716771, i: 0.76986, L: 313.23218, lp: 170.96424, lan: 74.22988, color: '#BBE1E4', radius: 3.98 },
  Neptune: { id: 'Neptune', name: '海王星', a: 30.06896348, e: 0.00858587, i: 1.76917, L: 304.88003, lp: 44.97135, lan: 131.72169, color: '#6081FF', radius: 3.86 },
};

export interface MoonOrbitalElements extends OrbitalElements {
  parent: string;
}

export const MOONS_DATA: Record<string, MoonOrbitalElements[]> = {
  Earth: [
    { id: 'Moon', name: '月', parent: 'Earth', a: 0.00257, e: 0.0549, i: 5.14, L: 135.27, lp: 83.35, lan: 125.08, color: '#D2D2D2', radius: 0.272 }
  ],
  Jupiter: [
    { id: 'Io', name: 'イオ', parent: 'Jupiter', a: 0.0028, e: 0.0041, i: 0.04, L: 342.3, lp: 100, lan: 0, color: '#F9F147', radius: 0.286 },
    { id: 'Europa', name: 'エウロパ', parent: 'Jupiter', a: 0.0045, e: 0.009, i: 0.47, L: 171.7, lp: 100, lan: 0, color: '#B7B19C', radius: 0.245 },
    { id: 'Ganymede', name: 'ガニメデ', parent: 'Jupiter', a: 0.0071, e: 0.0013, i: 0.2, L: 317.7, lp: 100, lan: 0, color: '#958878', radius: 0.413 },
    { id: 'Callisto', name: 'カリスト', parent: 'Jupiter', a: 0.0125, e: 0.0074, i: 0.19, L: 292.5, lp: 100, lan: 0, color: '#665C52', radius: 0.378 }
  ],
  Saturn: [
    { id: 'Enceladus', name: 'エンケラドゥス', parent: 'Saturn', a: 0.0016, e: 0.0047, i: 0.0, L: 200, lp: 100, lan: 0, color: '#FFFFFF', radius: 0.040 },
    { id: 'Titan', name: 'タイタン', parent: 'Saturn', a: 0.0082, e: 0.028, i: 0.33, L: 186, lp: 100, lan: 0, color: '#E4A422', radius: 0.404 }
  ],
  Neptune: [
    { id: 'Triton', name: 'トリトン', parent: 'Neptune', a: 0.0023, e: 0.000016, i: 156.8, L: 0, lp: 0, lan: 0, color: '#F4F4F4', radius: 0.212 }
  ]
};

// Aliases for MOONS_DATA to support Japanese names if used
MOONS_DATA['地球'] = MOONS_DATA['Earth'];
MOONS_DATA['木星'] = MOONS_DATA['Jupiter'];
MOONS_DATA['土星'] = MOONS_DATA['Saturn'];
MOONS_DATA['海王星'] = MOONS_DATA['Neptune'];

const DEG2RAD = Math.PI / 180;

export function getJulianDate(date: Date): number {
  return date.getTime() / 86400000 + 2440587.5;
}

export function getCenturiesSinceJ2000(jd: number): number {
  return (jd - 2451545.0) / 36525;
}

function solveKepler(M: number, e: number): number {
  let E = M;
  const tolerance = 1e-6;
  for (let i = 0; i < 10; i++) {
    const deltaE = (M - (E - e * Math.sin(E))) / (1 - e * Math.cos(E));
    E += deltaE;
    if (Math.abs(deltaE) < tolerance) break;
  }
  return E;
}

export function getPlanetPosition(planet: OrbitalElements, date: Date, isMoon: boolean = false): THREE.Vector3 {
  const jd = getJulianDate(date);
  const T = getCenturiesSinceJ2000(jd);

  let orbitalPeriod;
  if (isMoon) {
    orbitalPeriod = Math.pow(planet.a, 1.5) * 580; 
  } else {
    orbitalPeriod = Math.pow(planet.a, 1.5);
  }

  const meanMotion = (360 * 36525) / (365.25 * orbitalPeriod); // deg per century
  
  let L = planet.L + meanMotion * T;
  L = L % 360;
  if (L < 0) L += 360;

  const M = (L - planet.lp) * DEG2RAD;
  const e = planet.e;
  const a = planet.a;
  const i = planet.i * DEG2RAD;
  const lan = planet.lan * DEG2RAD;
  const lp = planet.lp * DEG2RAD;
  const w = lp - lan;

  const E = solveKepler(M, e);

  const x_plane = a * (Math.cos(E) - e);
  const y_plane = a * Math.sqrt(1 - e * e) * Math.sin(E);

  const cos_lan = Math.cos(lan);
  const sin_lan = Math.sin(lan);
  const cos_w = Math.cos(w);
  const sin_w = Math.sin(w);
  const cos_i = Math.cos(i);
  const sin_i = Math.sin(i);

  const x = x_plane * (cos_lan * cos_w - sin_lan * sin_w * cos_i) - y_plane * (cos_lan * sin_w + sin_lan * cos_w * cos_i);
  const y = x_plane * (sin_lan * cos_w + cos_lan * sin_w * cos_i) + y_plane * (cos_lan * cos_w * cos_i - sin_lan * sin_w);
  const z = x_plane * (sin_w * sin_i) + y_plane * (cos_w * sin_i);

  // 1 AU = 1000 units
  const SCALE = 1000;
  // Scale moon distance to be visible relative to planet size
  const DISTANCE_SCALE = isMoon ? SCALE * 15 : SCALE;
  
  return new THREE.Vector3(x * DISTANCE_SCALE, z * DISTANCE_SCALE, y * DISTANCE_SCALE);
}
