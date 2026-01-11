#!/bin/bash

# 卸载旧的docker
sudo yum remove docker \
    docker-client \
    docker-client-latest \
    docker-common \
    docker-latest \
    docker-latest-logrotate \
    docker-logrotate \
    docker-selinux \
    docker-engine-selinux \
    docker-engine

# 安装相关插件
sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2

# 增加库
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

# 安装
sudo yum install docker-ce
yum list docker-ce.x86_64  --showduplicates | sort -r
yum install -y --setopt=obsoletes=0 \
  docker-ce-18.06.1.ce-3.el7


# 开机自启动
sudo systemctl enable docker

# 启动
systemctl start docker

# 增加docker组
sudo groupadd docker

# 增加用户到docker组
sudo usermod -aG docker $USER

reboot


# 文档
k8s下
https://kubernetes.io/docs/setup/cri/#docker
原装
https://docs.docker.com/install/linux/docker-ce/centos/
