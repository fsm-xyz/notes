# 硬件概念

- [交换机](https://zh.wikipedia.org/wiki/%E7%B6%B2%E8%B7%AF%E4%BA%A4%E6%8F%9B%E5%99%A8)
- [集线器](https://zh.wikipedia.org/wiki/%E9%9B%86%E7%B7%9A%E5%99%A8)
- [网桥](https://zh.wikipedia.org/wiki/%E6%A9%8B%E6%8E%A5%E5%99%A8)
- [路由器](https://zh.wikipedia.org/wiki/%E8%B7%AF%E7%94%B1%E5%99%A8)

## 网卡

用来网络流量的硬件

## 虚拟网卡

用软件模拟出来的网卡

## 网段

## 网络模式

Bridge
[NAT](https://zh.wikipedia.org/wiki/%E7%BD%91%E7%BB%9C%E5%9C%B0%E5%9D%80%E8%BD%AC%E6%8D%A2)
Host Only

## ifconfig的信息

在CentOS7以及Ubuntu16.04往后的版本中，网卡设备号不再使用eth（有线）或wlan（无线）作为前缀来标识

- en      以太网
- wl      无线网卡
- lo      回环地址
- vir     虚拟接口

## VM的虚拟网络

vmnet0
vmnet1
vmnet8

## CentOS7网络配置相关文件

`/etc/resolv.conf`                                           # DNS配置文件
`/etc/hosts`                                                # 主机名到IP地址的映射 ,不该主机名基本不会动他。
`/etc/sysconfig/network`                                      # 所有的网络接口和路由信息，网关只有最后一个有效。
`/etc/sysconfig/network-script/ifcfg-<interface-name>`        # 每一个网络接口的配置信息
