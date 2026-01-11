# 配置文件

```sh
kubelet env
/var/lib/kubelet/kubeadm-flags.env

kubelet config
/var/lib/kubelet/config.yaml

certs
/etc/kubernetes/pki
etcd/ca
etcd/server
etcd/healthcheck-client
etcd/peer
apiserver-etcd-client
front-proxy-ca
front-proxy-client

kubeconfig
/etc/kubernetes
admin.conf
kubelet.conf
controller-manager.conf
scheduler.conf
/etc/kubernetes/manifests

/var/run/dockershim.sock

mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

## pod network

kubectl apply -f [podnetwork].yaml
<https://kubernetes.io/docs/concepts/cluster-administration/addons/>

## join node

```sh
kubeadm join 192.168.227.151:6443 --token 3uwe71.xg71nzyoyer6ta1s --discovery-token-ca-cert-hash sha256:ab1d3c3f37f9f0a5dd48671f1616db34b95bfd9b2e4e606c452fd7a3a8e68057
```
