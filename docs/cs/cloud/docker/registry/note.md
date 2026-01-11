# Mirrors

## 官方镜像

hub.docker.com

国内访问官方镜像有时候网络不稳定，这时候可以使用第三方的镜像加速器

阿里云

七牛云

## 其他镜像

docker.io
gcr.io
quay.io

使用翻墙软件

使用azure的镜像加速(目前仅限Azure China IP)

可用的注册中心

docker.io
dockerhub.azk8s.cn
hub-mirror.c.163.com

quay.io

quay.mirrors.ustc.edu.cn
quay-mirror.qiniu.com
quay.azk8s.cn

gcr.io
gcr.azk8s.cn
gcr.mirrors.ustc.edu.cn/

k8s.gcr.io
gcr.azk8s.cn/google-containers

registry.aliyuncs.com

## Note

修改配置文件  /etc/docker/daemon.json

registry-mirrors 使用特定的镜像中心
insecure-registries 使用http访问

```sh
# 修改配置
{
  "debug": true,
  "experimental": false,
  "registry-mirrors": [
    "http://192.168.199.111:5000",
    "https://b3j0mwsf.mirror.aliyuncs.com"
  ],
  "insecure-registries": [
    "http://192.168.199.111:5000"
  ]
}

# 重新启动
systemctl daemon-reload
systemctl  restart docker
```

## proxy

### daemon

```sh
mkdir -p /etc/systemd/system/docker.service.d
vi /etc/systemd/system/docker.service.d/http-proxy.conf

[Service]
# HTTP
Environment="HTTP_PROXY=http://192.168.199.111:1087"
# HTTPS
# Environment="HTTPS_PROXY=http://192.168.199.111:1087"
```

[proxy](https://docs.docker.com/engine/admin/systemd/#http-proxy)

## 搭建registry

[registry](https://github.com/docker/distribution)

```bash
docker pull registry
docker -itd -p 5000:5000 registry
```

### HELM

```bash
helm init --stable-repo-url http://mirror.azure.cn/kubernetes/charts/

helm repo remove stable
helm repo add stable http://mirror.azure.cn/kubernetes/charts/
helm repo add incubator http://mirror.azure.cn/kubernetes/charts-incubator/
helm repo update
helm repo list
```
