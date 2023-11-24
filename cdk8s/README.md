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

# pod アクセス
- nginx
http://test-site.remotehost/

- ubuntu
kubectl exec -it $(kubectl get pods -o=name --field-selector=status.phase=Running -n ubuntu) /bin/bash -n ubuntu

- arugo workflow
http://argo-workflows.remotehost/login


- argocd 
kubectl port-forward argocd-c8733361-server-54544d68c8-h9v7x 8080:argocd.remotehost:443 -n argocd

http://argocd.remotehost

user:admin
passowrd: `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d `

- argocd cli login
kubectl port-forward argocd-c8733361-server-54544d68c8-h9v7x 8080:grpc.argocd.remotehost:80 -n argocd
argocd login localhost:8080
