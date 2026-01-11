# 网络设备接口名字

## 为什么现在的eth0变成了ens33

简单的来讲, eth0 的名字是内核取的, 而这个名字是受驱动程序的先后顺序决定的. 多个网卡的话, 每次重启 ethx 后面的 数字可能会随机变化.
这也是改名的根本原因.

[参考](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/)
