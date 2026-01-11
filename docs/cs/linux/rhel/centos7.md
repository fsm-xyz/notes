# 配置ip

## 网络相关配置文件

```text
/etc/resolv.conf                        DNS
/etc/hosts                              主机跟ip的映射
/etc/sysconfig/network                  网络接口, 路由信息, 网关
/etc/sysconfig/network-scripts/ifcfg-*  网络接口的配置信息, 每一个网络接口对应一个文件
```

## DHCP

BOOTPROTO=dhcp

## static

```sh
BOOTPROTO=static
IPADDR=192.168.227.165
GATEWAY=192.168.227.2
NETMASK=255.255.255.0
DNS1=8.8.8.8
DNS2=192.168.227.2
```

## 重启网络

```sh
systemctl restart network
```
