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
