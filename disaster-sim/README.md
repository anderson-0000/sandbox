# Disaster Pulse Simulator

地震発生時の地震波（P波・S波）の伝播と、水深に基づいた津波の広がりを3Dで視覚化するリアルタイム・シミュレーター。

## 🚀 起動方法 (How to Run)

```bash
# プロジェクトディレクトリへ移動
cd disaster-sim

# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev
```

ブラウザで `http://localhost:5173` を開いてください。

## 🕹 使い方

1. **震源の設定**: キャンバス（海面）の任意の場所をクリックすると、その地点が震源（Epicenter）として設定されます。
2. **シミュレーション開始**: 左上のパネルの `Start` ボタンを押すと、時間の経過とともに地震波と津波が広がります。
3. **パラメータ調整**: マグニチュード（Magnitude）や再生速度（Playback Speed）をスライダーで変更できます。
4. **リセット**: `RotateCcw` アイコンを押すとシミュレーションがリセットされます。

## 🔍 シミュレーション内容

- **P波 (Primary Wave)**: 約 7.0 km/s で伝わる高速な縦波（青色の輪）。
- **S波 (Secondary Wave)**: 約 4.0 km/s で伝わる低速な横波（赤色の輪）。
- **津波 (Tsunami)**: 水深 4000m と仮定し、$v = \sqrt{gh} \approx 200$ m/s で広がる波面（シアン色の輪）。

## 🛠 技術スタック

- React 19 / TypeScript
- React-Three-Fiber (3D レンダリング)
- Zustand (状態管理)
- Lucide React (アイコン)
- Vanilla CSS (Tailwind 4)
