# docker desctop kubernetes

- 別のPCのdocker desctop kubernetesにkubectlで接続する

kubernetesが動いてるPCで`cat ~/.kube/config`し、中身をローカルPCの`~/.kube/config`に貼り付け
`ssh -L 6443:kubernetes.docker.internal:6443 -i ~/.ssh/id_rsa remotehost` でポートフォワードする
