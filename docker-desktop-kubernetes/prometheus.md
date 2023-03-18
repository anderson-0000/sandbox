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
  
