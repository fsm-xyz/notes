# Network

## 本地

* Bridge
* Host
* Container
* None

## 跨主机

* overlay
* macvlan
* third_party flannel、weave、calico
  
## 分析

host模式下，共享主机网络栈，效率损失最小，隔离性差

## Note

目前主机模式只能在主机模式下用

The host networking driver only works on Linux hosts, and is not supported on Docker Desktop for Mac, Docker Desktop for Windows, or Docker EE for Windows Server.

[host network](https://docs.docker.com/network/network-tutorial-host/)

## 资料

[参考](https://blog.csdn.net/Rapig1/article/details/102470936)
