# sandbox

検証用コードだけでなく、普段の業務で使う小さなツールや、実務で扱う技術の最小再現コードを管理するリポジトリです。

## プロジェクト構成

### インフラ・クラウド
- **[architecture/](./architecture/)**: AWS/Azure の構成図（Draw.io, Mermaid）
- **[azure-kubernetes/](./azure-kubernetes/)**, **[azure-terraform/](./azure-terraform/)**: Azure 関連の検証コード
- **[kubernetes-doc/](./kubernetes-doc/)**, **[docker-desktop-kubernetes/](./docker-desktop-kubernetes/)**: Kubernetes のドキュメントと設定
- **[argo-workflows/](./argo-workflows/)**: Argo Workflows のテンプレートとマニフェスト
- **[cdk8s/](./cdk8s/)**: cdk8s (Go) による Kubernetes マニフェスト生成
- **[kustomize/](./kustomize/)**: Kustomize の検証

### アプリケーション・ツール
- **[can-i-orbit/](./can-i-orbit/)**: 衛星投入 3D シミュレーター (React / Three.js)
  - **起動方法**: `cd can-i-orbit && npm install && npm run dev`
- **[orbit-visualizer/](./orbit-visualizer/)**: 衛星軌道の 3D 可視化ツール (React / Three.js / R3F)
  - **起動方法**: `cd orbit-visualizer && npm install && npm run dev`
- **[wealth-orbit/](./wealth-orbit/)**: 資産目標到達 3D シミュレーター (React / R3F)
  - **起動方法**: `cd wealth-orbit && npm install && npm run dev`
- **[space-debris-cleaner/](./space-debris-cleaner/)**: 宇宙ゴミ回収 3D ゲーム (React / R3F / satellite.js)
  - **起動方法**: `cd space-debris-cleaner && npm install && npm run dev`
- **[solar-system-now/](./solar-system-now/)**: 太陽系の現在の惑星位置シミュレーター
- **[zoom-to-atom/](./zoom-to-atom/)**: 宇宙のスケールから原子までズームする科学的可視化ツール
- **[blood-pressure-app/](./blood-pressure-app/)**: 血圧管理アプリ (JS/HTML/CSS)
- **[investment-web-app/](./investment-web-app/)**: 投資シミュレーション Web アプリ (Python)
- **[job-draft/](./job-draft/)**: 求人情報やユーザーランキングの取得ツール
- **[open-weather-map/](./open-weather-map/)**: OpenWeatherMap API 連携
- **[slack/](./slack/)**, **[switch-bot/](./switch-bot/)**: 各種 API/デバイス連携ツール
- **[print-kids/](./print-kids/)**: 学習プリント用ダウンロードスクリプト

### プログラミング言語別検証
- **[go/](./go/)**: Go 言語の基本機能やチャット実装
- **[python/](./python/)**: Python の基本機能、数学的検証
- **[rust/](./rust/)**: Rust の基本
- **[R/](./R/)**: R 言語によるシミュレーションと図形描画
- **[shell/](./shell/)**: シェルスクリプトによるツール

### 学習・その他
- **[atcoder/](./atcoder/)**: AtCoder の回答コード (Python)
- **[English/](./English/)**: 英語学習用メモ
- **[github-actions/](./github-actions/)**, **[github-apps/](./github-apps/)**: GitHub 連携・自動化
- **[vimrc/](./vimrc/)**, **[zshrc/](./zshrc/)**: 設定ファイル (Dotfiles)
