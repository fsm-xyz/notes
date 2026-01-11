# Docker

## 核心概念

+ cgroups(CPU 内存)
+ namespaces(网络，进程，Mount)
+ 文件系统(Overlay, COW)

+ 镜像
+ 容器

共享宿主机内核, 通过cgroups和namespaces进行资源隔离和资源分配

## 镜像

镜像分层存储（只读层 + 读写层），写时复制（CoW）节省空间，文件发生了变化才会重新生成新的存储层

### 构建

#### 多级构建

+ 中间步骤进行多层RUN，尽量利用缓存
+ 最终镜像尽量合并RUN指令，只复制最终产物
+ 稳定不动的放前面
+ 多利用缓存，加速
+ 并行构建
+ 最小化基础镜像(alpine, distroless, scratch, slim, chainguard
+ .dockerignore 忽略不必要的文件到构件环境种


## 网络

+ bridge(port端口映射,默认)
+ host
+ container共享
+ none
+ macvlan 为容器分配独立 MAC 地址，直接绑定物理网卡
+ ipvlan
+ Overlay      基于 VXLAN 封装跨主机流量，需配合 Swarm/K8s 集群
+ Calico	BGP 路由宣告容器 IP

### 技术

+ 网桥
+ veth pair
+ NAT
