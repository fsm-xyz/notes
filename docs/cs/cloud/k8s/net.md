# 网络模式

## docker

### docker Daemmon网络

创建docker0网桥，设置iptables规则

### docker 容器网络模式

+ bridge模式
+ host模式
+ other容器模式
+ none
+ overlay
+ ipvlan
+ macvlan

### 暴露访问

+ port mapping
+ 主机网络

## k8s

### Pod网络(CNI支持)

+ Overlay
+ BGP
+ Underlay

Flannel(Bridge)
Calico(BGP)
Antrea(OpneSwitch)

## 暴露访问


+ hostNetwork
+ hostPort

+ LB
+ NodePort
