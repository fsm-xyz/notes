# storage

本地存储, 网络存储

## 分类

+ Pod Volumes
+ Persistent Volumes

### Pod Volumes

+ 本地存储(emptydir/hostpath)
+ 网络(secret/configmap)

### Persistent Volumes

+ Static Volume Provisioning
+ Dynamic Volume Provisioning

#### local pv

使用本地文件

```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: local-storage
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
```
