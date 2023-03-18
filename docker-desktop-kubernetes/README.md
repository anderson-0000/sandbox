# Note on docker desktop kubernetes 
- 接続先確認
```
kubectl config get-contexts
kubectl config view
```

- 接続先変更
```
kubectl config use-context docker-desktop
```

- ノード確認
```
kubectl get nodes
kubectl describe nodes
```

- ネームスペース確認
```
kubectl get namespace
```

```
default:他にNamespaceを持っていないオブジェクトのためのデフォルトNamespace
kube-system:Kubernetesシステムによって作成されたオブジェクトのためのNamespace
kube-public:このNamespaceは自動的に作成され、全てのユーザーから読み取り可能です。(認証されていないユーザーも含みます。) このNamespaceは、リソースをクラスター全体を通じてパブリックに表示・読み取り可能にするため、ほとんどクラスターによって使用される用途で予約されます。]
kube-node-lease クラスターのスケールに応じたノードハートビートのパフォーマンスを向上させる各ノードに関連したLeaseオブジェクトのためのNamespace。
参照：https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/namespaces/
```


- ポッド確認
```
kubectl get pods --all-namespaces
kubectl get pods -n <名前空間>
kubectl describe pods <ポッド名> -n <名前空間>
```

- サービスの確認
```
kubectl get service --all-namespaces
```

- ポッド作成
```
kubectl apply -f <ポッドの定義ファイル>
```

- デプロイメントの確認
```
kubectl get deployment --all-namespaces
```

- kubernetes ダッシュボード
```
https://github.com/kubernetes/dashboard
```

  - インストール
  ```
  kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
  ```

  - 初期設定（サービスアカウントとClusterRoleBindingを作成）
  ```
  vi dashbord-adminuser.yaml
  https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md
  kubectl apply -f dashbord-adminuser.yaml
  ```

  - ベアラー トークンの取得
  ```
  kubectl -n kubernetes-dashboard create token admin-user | pbcopy
  ```

  - kubernetes apiサーバーへのプロキシを開始
  ```
  kubectl proxy
  ```

  - ブラウザからアクセス

  http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login

- Prometheus
  - 専用の名前空間を作成

  ```
  kubectl create namespace prometheus
  ```

  - インストール
  ```
  helm install prometheus prometheus-community/prometheus --namespace prometheus --set prometheus-node-exporter.hostRootFsMount.enabled=false
  ```

  - 更新のリリース
  ```
  helm upgrade prometheus prometheus-community/prometheus --namespace prometheus --set prometheus-node-exporter.hostRootFsMount.enabled=false
  ```

  - ポートフォワードする
  ```
  kubectl port-forward prometheus-server-5fc7649bd-frnsk 8080:9090
  ```

    - ブラウザからアクセス
    http://localhost:8080/

  - kubernetes apiサーバーへのプロキシを開始
  ```
  kubectl proxy
  ```

    - ブラウザからアクセス

    http://localhost:8001/api/v1/namespaces/prometheus/services/http:prometheus-server:80/proxy/
  
- grafana
  - 専用の名前空間を作成
  ```
    kubectl create namespace grafana
  ```

  - インストール
  ```
  helm install my-release grafana/grafana --namespace grafana
  ```

  - kubernetes apiサーバーへのプロキシを開始
  ```
  kubectl proxy
  ```

    - ブラウザからアクセス

    http://localhost:8001/api/v1/namespaces/grafana/services/http:my-release-grafana:80/proxy/

  - adminユーザーのパスワードを取得
  ```
  kubectl get secret --namespace grafana my-release-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
  ```
  
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

- インストールできるチャート一覧を表示
```
helm search repo
```

- チャートのダウンロード
```
helm pull <チャートURL | リポジトリ名/チャート名>
tar -xzvf <ファイル名>
```

- チャートをインストール
```
helm install <名前> <リポジトリ名> --namespace <名前空間>
```

- チャートの新バージョンをリリース
```
helm upgrade  <名前> <リポジトリ名> --namespace <名前空間>
```

- リリース一覧を表示
```
helm list -n <名前空間>
```

- リリースのステータスを確認
```
helm status <リリース名>
```

- リリースをクラスタから削除
```
helm delete <リリース名>
```


