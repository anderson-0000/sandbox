# Azure Terraform

このディレクトリには、Terraformを使用してAzureのインフラを構築するための設定ファイルが格納されています。

## ディレクトリ構成

- `modules/`: 再利用可能なTerraformモジュールを格納します。
  - `kubernetes/`: AKSクラスタ関連のモジュール
  - `network/`: ネットワーク関連のモジュール
  - `resource_group/`: リソースグループ関連のモジュール
- `service/`: `modules` を利用してインフラを構築するための具体的なTerraformコードを格納します。

## 使い方

```bash
cd service
terraform init
terraform plan
terraform apply
```
