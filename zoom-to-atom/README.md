# Zoom to Atom Simulation

「人間」から「分子」「原子」「原子核」まで、スケールを自在にズームイン・アウトできる3Dシミュレーションアプリケーションです。

## 起動方法 (How to Run)

1.  依存関係のインストール:
    ```bash
    npm install
    ```

2.  開発サーバーの起動:
    ```bash
    npm run dev
    ```

3.  ブラウザで `http://localhost:5173/` にアクセスしてください。

## 操作方法
- **ズーム:** マウスホイールまたはタッチ操作でズームイン・アウト。
- **ドラッグ:** カメラの微調整（OrbitControls）。

## 技術構成
- **Frontend:** React (Vite) + TypeScript
- **3D Engine:** React-Three-Fiber (R3F) + Three.js
- **State Management:** Zustand
- **Components:** `Tracker` パターンによる R3F 安全設計
