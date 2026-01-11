# 网络工具

## 网络工具命令

```sh
ss
netstat
plipconfig
hostname
arp
ifconfig
ipmaddr
iptunnel
mii-tool
nameif
rarp
route
und
slattach
```

### ip

```sh
ip addr
ip addr show
ip addr add 192.168.0.1/24 dev eth0 # 设置eth0网卡IP地址192.168.0.1
ip addr del 192.168.0.1/24 dev eth0 # 删除eth0网卡IP地址
ip link set ens33 down
ip link set ens33 up
ip link show                             # 显示网络接口信息
ip link set eth0 up                  # 开启网卡
ip link set eth0 down             # 关闭网卡
ip link set eth0 promisc on   # 开启网卡的混合模式
ip link set eth0 promisc offi  # 关闭网卡的混个模式
ip link set eth0 txqueuelen 1200    # 设置网卡队列长度
ip link set eth0 mtu 1400      # 设置网卡最大传输单元

ip route show 或 ip route list  或   route -n  # 查看路由(网关)信息
ip route add 192.168.4.0/24  via  192.168.0.254 dev eth0 # 设置192.168.4.0网段的网关为192.168.0.254,数据走eth0接口
ip route add default via  192.168.0.254  dev eth0    # 设置默认网关为192.168.0.254
ip route del 192.168.4.0/24    # 删除192.168.4.0网段的网关
ip route del default    # 删除默认路由
```

### dns

```sh
nslook google.com       # 查看dns
```

## 路由

[参考](https://www.cnblogs.com/hf8051/p/4530906.html)
