- pod →EC2インスタンスメタデータへのアクセス制御
Podのセキュリティグループかネットワークポリシーで制御
制限するとkubernetes node のインスタンスプロファイルが読めなくなるので、IRSAの設定を行う。
podのセキュリティグループを適用するとネットワークポリシーが適用されない
  - IRSA(IAM Roles for Service Accounts)
    k8sのサービスアカウントとawsのiam roleを紐づける。EKSの機能        
      - サービスアカウントはpodが使うkubernetesが管理してるアカウント
      - ユーザーアカウントは人間が使うkubernetsが管理してる　アカウント

- podのsgとネットワークポリシーの違い
ネットワークポリシーは、ポッド間の通信のアクセスポリシーを定義するKubernetesの仕様です。
ネットワークポリシーを使用して、トラフィックを送受信するための順序付けされたルールセットを定義し、ラベルに一致するポッドに適用でる。

Podのsgは、aws ec2 sg をkubernetes podと統合したもの。
ネットワークポリシーはkubernetes内でポッド間の通信を制御する。
sgはAWS内でポッドとawsリソースの通信を制御するために使用されます。

- コンテナの実行ユーザーを設定
  - コンテナは基本一般ユーザーで実行させた方がセキュリティ上良い
  - 設定方法
    - DockerfiledでUSERを指定する
    - kubernetes マニュフェストで spec.containers.securityContext.runAsUserとspec.containers.securityContext.runAsNonRootを指定して設定を上書きする

- spoec.securityContext.seccompProfile.typeでノードOS(linux)のシステムコールを制限できる
- PSSとPSAで設定したポリシーに違反したpodのデプロイを制限できる
  - だた細かい設定はできないので、細かい設定がしたいならkyvernoを使う必要がある。
- wafでl7の攻撃対策
- リリースしたらネームスペースごとに権限管理した方がいい
- rbacとaws iamとaws-authを連携して権限管理
  - rbac managerとrbac lookupを使うと管理が少し楽になるらしい
  - podごとに権限分ける


- kubernetesのセキュリティベストプラクティス
  - CIS Kubernetes Benchmarks
    https://www.cisecurity.org/benchmark/kubernetes
    - kube-benchは実行中のクラスターが↑に準拠しているかチェックしてレポートしてくれるツール
- 脆弱性チェック
  - trivy
- Dockerファイルの設定不備
  - hadolint
- kubernetesマニフェストとhelm chartの設定不備
  - Polaris
- イメージのマルウェア検知
  - オープンソースだとない？
  - やるならお金払ってprisma cloud など使う必要がある
- ノード側のコンテナネイティブのセキュリティツール
  - Falco

- kubernetes のGUI化
  - Lens Desktop
    - kubernetes clusterの設定をguiで可視化できる
    - prometeusメトリクスの可視化
    - podのログの表示(pod単体ごとのログしか見れない)
    - podへのログイン
    - helm chartの￥検索とclusterへのデプロイ
    - kubernetes dashboardとの違いはデプロイが不要で、ローカルのkubeconfigを読み込んでclusterにアクセスできる。

- stern
  - 同一プレフィックスのpod logや同一ラベルのログを一緒に見れる
- kube-ps1
  - 今どのクラスターに入っているか分かりやすく表示する
- kube-ns
  - kubectlで使われるデフォルトネームスペースを変更できる

