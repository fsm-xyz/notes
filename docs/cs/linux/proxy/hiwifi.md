# proxy

## 方式

### 透明代理

1. 主路由设置透明代理
2. 旁路由设置透明代理，其他设备设置网关为旁路由，DNS为旁路由
3. 设备直接安装proxy服务，开启透明代理

### 传统代理

1. 旁路由设置代理服务，其他设备设置系统proxy地址为旁路由proxy服务地址
2. 设备直接安装proxy服务，直接设置proxy代理地址


ssh root@192.168.199.1 -p 22 -oHostKeyAlgorithms=+ssh-rsa

## FAQ

传统代理linux基于env实现（程序读取env, 有时候需要restart）
透明代理linix基于iptables

时间错误，需要设置ntp来同步时间

https://cdn.jsdelivr.net/china/gh/Dreamacro/maxmind-geoip@release/Country.mmdb

### 

```sh
# 详细信息
dmesg


root@Hiwifi:~# cat /proc/cpuinfo 

system type             : Mediatek MT7620A ver:2, eco:6
machine                 : HiWiFi Wireless R32 Board
processor               : 0
cpu model               : MIPS 24KEc V5.0
BogoMIPS                : 385.84
wait instruction        : yes
microsecond timers      : yes
tlb_entries             : 32
extra interrupt vector  : yes
hardware watchpoint     : yes, count: 4, address/irw mask: [0x0ff8, 0x0ff8, 0x0ff8, 0x0ff8]
isa                     : mips1 mips2 mips32r1 mips32r2
ASEs implemented        : mips16 dsp
shadow register sets    : 1
kscratch registers      : 0
core                    : 0
VCED exceptions         : not available
VCEI exceptions         : not available


```
cat /proc/cmdline 
 board=R32 console=ttyS1,115200 rootfstype=squashfs,jffs2

cat /proc/meminfo 
MemTotal:         126148 kB
MemFree:           62116 kB
Buffers:            4360 kB
Cached:            12724 kB
SwapCached:            0 kB
Active:            23432 kB
Inactive:           7704 kB
Active(anon):      14572 kB
Inactive(anon):      492 kB
Active(file):       8860 kB
Inactive(file):     7212 kB
Unevictable:           0 kB
Mlocked:               0 kB
SwapTotal:         62460 kB
SwapFree:          62460 kB
Dirty:                 0 kB
Writeback:             0 kB
AnonPages:         14068 kB
Mapped:             5460 kB
Shmem:              1012 kB
Slab:              10760 kB
SReclaimable:       2820 kB
SUnreclaim:         7940 kB
KernelStack:         600 kB
PageTables:          648 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:      125532 kB
Committed_AS:      50804 kB
VmallocTotal:    1048372 kB
VmallocUsed:       13064 kB
VmallocChunk:    1020052 kB


cat /proc/mtd
dev:    size   erasesize  name
mtd0: 00080000 00020000 "u-boot"
mtd1: 00080000 00020000 "debug"
mtd2: 00040000 00020000 "Factory"
mtd3: 02000000 00020000 "firmware"
mtd4: 00140000 00020000 "kernel"
mtd5: 01ec0000 00020000 "rootfs"
mtd6: 00080000 00020000 "hw_panic"
mtd7: 00080000 00020000 "bdinfo"
mtd8: 00080000 00020000 "backup"
mtd9: 01000000 00020000 "overlay"
mtd10: 02000000 00020000 "firmware_backup"
mtd11: 00200000 00020000 "oem"
mtd12: 02b40000 00020000 "opt"

