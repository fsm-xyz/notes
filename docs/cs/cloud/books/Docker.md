# Docker源码分析

容器管理引擎，资源管理

容器是一个抽象的概念，本质是一个进程，由父进程fork产生

## 核心linux的技术

namespaces
cgroups
namespaces主要负责命名空间的隔离，cgroups负责资源限制

## 模块

DockerClient
DockerDaemon
Server
Engine
Job
Dirver
LXC
libcontainer
DockerRegistry
Graph
image
volumes

## Daemon网络

桥接(bridge, veth pair, iptables)

## 容器网络

桥接
由于宿主机的ip和veth pair的不在同一网段，还不足以使宿主机以外的网络发现容器，为使外部可以访问容器服务，使用NAT技术进行网络报文转发，端口映射

主机

其他容器
kubernetes的Pod实现，容器间共享原理，实现组的概念

none

## 镜像

linux系统启动时，内核回挂载一个只读的rootfs，当检查完整性后，决定是否将会其切换成读写模式或者最后挂载另一种文件系统忽略rootfs

rootfs代表容器启动时，内部进程可见的文件系统或者Docker容器的根目录，利用联合挂载的技术，在rootfs上面挂载一个读写文件系统

文件系统，挂载点

COW和whiteout

image只是包含一小部分文件的集合，rootfs由很多个image组成

父镜像和根镜像

layer的概念包含可以读写的文件这一层

/var/lib/docker

mnt diff layers

## build

context
多级构建，合并RUN命令
每次都会在一个容器环境中运行
容器link

## dockerinit

容器内运行的第一个进程是dockerinit进程，类似linux的init进程，负责初始换系统，是所有其他进程的祖先进程

## Swarm

单机Docker在分布式的环境下，比较有限，Swarm提供了Docker集群能力

## Machine

docker-machine, 类似vgrant

对接各种IaaS，快速创建虚拟的Docker节点和集群

## Compose

原生管理容器不方便，提供配置文件进行管理，还可以进行多容器的管理
