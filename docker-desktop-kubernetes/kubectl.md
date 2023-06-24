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

- ログ確認
```
kubectl logs <ポッド名>
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

- apply前確認
```
kubectl diff -f <ポッドの定義ファイル>
```

- ポッド作成
```
kubectl apply -f <ポッドの定義ファイル>
```

- ポッド削除
```
kubectl delete -f <ポッドの定義ファイル>
```

- 利用可能なAPIリソースの一覧を表示
```
kubectl api-resources
```

- podのオートスケール確認
```
kubectl get hpa -w
```

- 現在のコンテキスト(接続先)を調べる
```
kubectl config current-context
```

- コンテキストの一覧
```
kubectl config get-contexts
```

- コンテキスト切り替え
```
kubectl config use-context ${コンテキスト}
```

- 全ネームスペースの全サービスを表示
```
kubectl get all -A
```
