# docker desctop kubernetes

- 別のPCのdocker desctop kubernetesにkubectlで接続する

```
if ! ps aux | grep "6443:kubernetes.docker.internal:6443" | grep -v grep > /dev/null
then
  ssh -fN -L 6443:kubernetes.docker.internal:6443 remotehost
fi
ssh remotehost cat .kube/config > .kube/config
```
