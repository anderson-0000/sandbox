# Note on helm
- インストール
```
brew install helm
```

- プラグインをインストール
```
helm plugin install https://github.com/databus23/helm-diff
```

- プラグイン一覧表示
```
helm plugin list
```

- チャートリポジトリに公式の Helm Stable チャートを追加
```
helm repo add stable https://charts.helm.sh/stable
```

- ダッシュボードリポジトリを追加
```
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
```

- プロメテウスコミュニティのリポジトリを追加
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

- グラファナのリポジトリを追加
```
helm repo add grafana https://grafana.github.io/helm-charts
```

- Calicoのリポジトリを追加
```
helm repo add projectcalico https://docs.tigera.io/calico/charts

```

- Bitnamiのリポジトリを追加
```
helm repo add bitnami https://charts.bitnami.com/bitnami
```

- リポジトリ一覧表示
```
helm repo list
```

- リポジトリを削除
```
helm repo remove <リポジトリ名>
```

- チャートリポジトリから最新のチャート情報を取得
```
helm repo update 
```

- ローカルのhelmクライアントに保存したリポジトリを検索
```
helm search repo
```

- helmチャートでカスタマイズできる項目
```
helm inspect values <リポジトリ名> > values.yaml
```
