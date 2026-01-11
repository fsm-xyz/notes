# History

## 公司

随着发展公司名字由dotcloud到docker
核心依赖从LXC变为golang编写的containerd
原来的库拆分为cli和engine

在 Docker 1.8 之前，Docker 守护进程启动的命令为
docker -d

Docker 1.8 开始，启动命令变成了
docker daemon

Docker 1.11 开始，守护进程启动命令变成了
dockerd

2017, 项目由dokcer变迁为moby, 官方成立了一个moby的开源组织维护
任何组织和个人都可以基于moby构建自己的容器产品
docker官方在此基础上构建自己的产品

## moby

[moby](https://github.com/moby/moby)

## docker官方的仓库

[engine](https://github.com/docker/engine)
[cli](https://github.com/docker/cli)
[compose](https://github.com/docker/compose)
[swarm](https://github.com/docker/swarm)
[docker-ce](https://github.com/docker/docker-ce)
...

## containered

[containerd](https://github.com/containerd/containerd)

## Products

官方基于engine构建不同的产品，主要是Docker CE和Docker EE

Docker CE

+ Linux
+ Docker Desktop(Mac, Windows)
+ Docker ToolBox(Mac, Windows)

Docker EE
Docker Enterprise Platform
Docker Hub

## 关系

![img](https://pbs.twimg.com/media/C-KGrQ1XsAEkq1P.jpg:large)

## Runtime

LXC
libcontainer
runc
containerd
cri-o

## referrence

<https://medium.com/devopslinks/an-overall-view-on-docker-ecosystem-containers-moby-swarm-linuxkit-containerd-kubernetes-5e4972a6a1e8>
<https://www.mirantis.com/blog/ok-i-give-up-is-docker-now-moby-and-what-is-linuxkit/>
<http://alexander.holbreich.org/docker-moby/>
