# 磁盘

## 磁盘信息

sfdisk -l
free -m
free -h

## swap

### 禁用和优先级

cat /proc/swaps

### 临时禁用

```bash
swapoff -a      # 全部禁用
swapon -a       # 全部打开
swapon --show   # 开启信息
swapon -s
swapoff /swapfile   #禁用指定的swap设备
```

### 永久禁用

文件    /etc/fstab

```bash
vi /etc/fstab
# 永久禁用
# 注释掉swap的那一行
```

#### 优先级

文件    /etc/sysctl.conf

```bash
vim /etc/sysctl.conf
vm.swappiness=0             # 0-100,默认的值是60,越大表示尽量使用swap
sysctl -p                   # 生效
```
