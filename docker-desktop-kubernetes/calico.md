- calicoとは
ネットワークポリシー（awsのsg）を利用してkubernetes内ネットワークの通信制御をするために必要なcontainer network interface。

- 専用のnamespaceを作成
```
kubectl create namespace tigera-operator

```

- helmからインストール
```
helm install calico projectcalico/tigera-operator --version v3.25.0 --namespace tigera-operator

```


