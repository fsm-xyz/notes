# 网络

## debian网路配置

```sh
vim /etc/network/interfaces

#开机自动激活eth0接口
auto eth0

#配置eth0接口为静态设置IP地址
iface eth0 inet static
address 192.168.60.110
netmask 255.255.255.0  
gateway 192.168.60.2
```

## debian自动补全

```sh
apt-get install bash-completion
```
