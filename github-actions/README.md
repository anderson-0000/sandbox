# GitHub Actions

GitHub Actions のワークフロー置き場です。  
Terraform を中心に、インフラまわりの CI の最小構成をまとめています。

## 含まれるもの
- **terraform-ci.yml**  
  PR 時に実行される Terraform の CI。  
  - fmt  
  - validate  
  - tflint（--init）  
  - trivy  
  - plan（PR に結果をコメント）

## 目的
- Terraform CI の最小構成の保存  
- 実務で使う CI のベースとして利用
