# プロジェクトを作成
$ mkdir ${プロジェクト名}
$ cd ${プロジェクト名}
$ cdk8s init ${プロジェクト名}

# kubernetesのAPIオブジェクトをcdk8sで使えるようにする
## 特定のバージョンのすべてのKubernetes APIオブジェクト構造がプロジェクト内のimportsディレクトリにファイルが生成される。
$ cdk8s import

# オブジェクトの定義
$ vi main.go
// define resources here
の行にコードを追加

# kubernetes マニュフェストを作成
$ cdk8s synth

# kubernetesクラスターにデプロイ
$ kubectl apply -f dist/${プロジェクト名}.k8s.yaml

# helmの変更可能なパラメータのYAMLファイルを出力する
helm inspect values argo/argo-cd > default_value.yaml
