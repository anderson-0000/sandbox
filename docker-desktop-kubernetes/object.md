# Kubernetesオブジェクトは.ymlで指定する
https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/kubernetes-objects/

## apiVersion
  どのバージョンのkubernetes apiを利用してオブジェクトを作成するか。`v1` `apps/v1` `batch/v1` など。 

## kind
  どの種類のオブジェクトを作成するか。この内容によってapiVersionがきまる。`Deployment` `Service` `StatefullSet` `CronJob` `Job` など。

  以下コマンドで使える一覧が参考になる
  ```
  kubectl api-resources
  ```

## metadata
  オブジェクトを一意に特定するための情報。`name:` `namespace:` など。

## spec
  オブジェクトの望ましい状態を定義する。フォーマットは、Kubernetesオブジェクトごとに異なり、オブジェクトごとに特有な入れ子のフィールドを持っています。

  https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/ でspecで使えるキーを探すと良い。
