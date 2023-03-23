# kustomize
標準で入っていて、namespaceの環境差分を吸収してくれるツール

workspaceを使っているterraformみたいな感じで、variable.tfにまとめて環境ごとの値を設定する感じ。

これあればhelmのvalueを環境ごとに作って管理する必要はない？？

- ディレクトリ構成

.
  `base`
    各種yaml
  `overlays`
     `develop`
        kustomization.yaml
     `staging`
        kustomization.yaml
     `production`
        kustomization.yaml

- 作成されるリソース一覧を表示
`kustomize build overlays/develop`

- リソースをクラスターにリリース
`kubectl apply -k overlays/develop`

- リソースの削除
`kubectl delete -k overlays/develop`
