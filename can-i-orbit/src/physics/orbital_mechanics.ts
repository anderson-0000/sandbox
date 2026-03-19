export interface Vector3D {
  x: number;
  y: number;
  z: number;
}

export type State = {
  position: Vector3D;
  velocity: Vector3D;
  time: number;
};

// 定数 (単位: m, kg, s)
export const G = 6.67430e-11;
export const EARTH_MASS = 5.972e24;
export const EARTH_RADIUS = 6.371e6; // 6371 km

// 加速度を計算: a = - (G * M / r^3) * r
function getAcceleration(pos: Vector3D): Vector3D {
  const r2 = pos.x * pos.x + pos.y * pos.y + pos.z * pos.z;
  const r = Math.sqrt(r2);
  if (r < EARTH_RADIUS) return { x: 0, y: 0, z: 0 }; // 地球内部（衝突）

  const mag = -(G * EARTH_MASS) / (r * r2);
  return {
    x: mag * pos.x,
    y: mag * pos.y,
    z: mag * pos.z,
  };
}

// 4次ルンゲ＝クッタ法 (RK4)
export function integrate(state: State, dt: number): State {
  const { position: p, velocity: v } = state;

  // k1
  const v1 = v;
  const a1 = getAcceleration(p);

  // k2
  const p2 = { x: p.x + v1.x * dt / 2, y: p.y + v1.y * dt / 2, z: p.z + v1.z * dt / 2 };
  const v2 = { x: v.x + a1.x * dt / 2, y: v.y + a1.y * dt / 2, z: v.z + a1.z * dt / 2 };
  const a2 = getAcceleration(p2);

  // k3
  const p3 = { x: p.x + v2.x * dt / 2, y: p.y + v2.y * dt / 2, z: p.z + v2.z * dt / 2 };
  const v3 = { x: v.x + a2.x * dt / 2, y: v.y + a2.y * dt / 2, z: v.z + a2.z * dt / 2 };
  const a3 = getAcceleration(p3);

  // k4
  const p4 = { x: p.x + v3.x * dt, y: p.y + v3.y * dt, z: p.z + v3.z * dt };
  const v4 = { x: v.x + a3.x * dt, y: v.y + a3.y * dt, z: v.z + a3.z * dt };
  const a4 = getAcceleration(p4);

  return {
    position: {
      x: p.x + (dt / 6) * (v1.x + 2 * v2.x + 2 * v3.x + v4.x),
      y: p.y + (dt / 6) * (v1.y + 2 * v2.y + 2 * v3.y + v4.y),
      z: p.z + (dt / 6) * (v1.z + 2 * v2.z + 2 * v3.z + v4.z),
    },
    velocity: {
      x: v.x + (dt / 6) * (a1.x + 2 * a2.x + 2 * a3.x + a4.x),
      y: v.y + (dt / 6) * (a1.y + 2 * a2.y + 2 * a3.y + a4.y),
      z: v.z + (dt / 6) * (a1.z + 2 * a2.z + 2 * a3.z + a4.z),
    },
    time: state.time + dt,
  };
}

// 衝突判定
export function checkCollision(pos: Vector3D): boolean {
  const r2 = pos.x * pos.x + pos.y * pos.y + pos.z * pos.z;
  return r2 < EARTH_RADIUS * EARTH_RADIUS;
}

// 脱出判定 (地球の第2宇宙速度 11.2km/s 程度)
export function checkEscape(pos: Vector3D, vel: Vector3D): boolean {
  const r = Math.sqrt(pos.x * pos.x + pos.y * pos.y + pos.z * pos.z);
  const v2 = vel.x * vel.x + vel.y * vel.y + vel.z * vel.z;
  const escapeVelocity2 = (2 * G * EARTH_MASS) / r;
  return v2 > escapeVelocity2 && r > EARTH_RADIUS * 10; // 十分に離れていて脱出速度を超えている場合
}
