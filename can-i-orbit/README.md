# Can I Orbit?

衛星を地球の軌道に投入できるかを楽しむ 3D 物理シミュレーターです。

## 概要
万有引力の法則に基づいた物理計算（4次ルンゲ＝クッタ法）を行い、指定した初速と角度で衛星がどのような軌道を描くかをシミュレーションします。

- **第1宇宙速度 (~7.9 km/s)**: 地球を周回する安定した円軌道に乗ります。
- **第2宇宙速度 (~11.2 km/s)**: 地球の重力を振り切り、宇宙の彼方へ脱出します。

## 技術スタック
- **Frontend**: React (TypeScript)
- **3D Engine**: Three.js (React Three Fiber)
- **Physics**: 手実装の軌道力学エンジン (RK4)

## 起動方法

```bash
cd can-i-orbit
npm install
npm run dev
```

起動後、ブラウザで [http://localhost:5173](http://localhost:5173) を開いてください。

## 操作方法
- **Launch Velocity**: 打ち上げ時の初速 (km/s)
- **Launch Angle**: 打ち上げ時の水平からの角度 (degrees)
- **Mouse**: 
  - 左ドラッグ: カメラ回転
  - 右ドラッグ: 平行移動
  - スクロール: ズーム
