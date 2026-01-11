# Debian

升级

## 查看当前系统

```sh
uname -r
cat /etc/issue
cat /etc/os-release
cat /etc/debian_version
lsb_release -a
```
## 版本升级

修改源为新版

```sh
# /etc/apt/sources.list
# 替换为新版的地址

apt update && apt full-upgrade
apt dist-upgrade -y

reboot

lsb_release -a

```

## 内核升级

```sh
# 修改`/etc/apt/sources.list`
# 增加或者取消注释`backports`这一行
echo "deb http://deb.debian.org/debian bullseye-backports main" | tee -a /etc/apt/sources.list
deb http://mirrors.tencentyun.com/debian bullseye-backports main contrib non-free

apt update
apt upgrade -y
apt dist-upgrade -y

# apt -t bullseye-backports upgrade
# apt -t bullseye-backports install xxx

apt search linux-image

apt install linux-image-6.1.0-0.deb11.6-amd64-unsigned
apt install linux-headers-6.1.0-0.deb11.6-amd64

# 写入boot，重启

update-grub
reboot
uname -r

# 卸载旧内核
dpkg --list | grep linux-image
dpkg --list | grep linux-headers

apt purge linux-image-4.19.0-11-amd64
apt purge linux-headers-4.19.0-11-amd64

apt autoclean && apt --purge autoremove

```
