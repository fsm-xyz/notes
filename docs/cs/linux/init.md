# init

## Linux 启动过程

* SysV init
* systemd

### SysV init启动

```sh
开机 -> BIOS自检 ->查找引导设备(MBR, EFI)并加载 -> 执行引导程序(grub, grub2) -> 加载内核, 执行/sbin/init程序, 读取/etc/inittab(运行级别), 获取/etc/fstab(分区信息)挂载, 启动/etc/init.d下的服务

* init进程 是linux的第一个进程, PID 1, 其他进程都是它的子进程
* 每次只启动一个服务, init负责后台管理启动的服务
```

#### 关机

相反过程, 停止服务, 写在文件系统

### systemd启动

开机 -> BIOS自检 ->查找引导设备(MBR, EFI)并加载 -> 执行引导程序(grub) -> 加载内核, 执行systemd程序, 读取/etc/inittab(运行级别), 获取/etc/fstab(分区信息)挂载, 并发启动/etc/systemd/system/下的服务

#### 主要优点

* 流程简化
* 启动速度快
* 并发启动服务
* 用cgroup
* 性能分析工具

#### 关机流程

执行/usr/lib/systemd下的对应服务

```sh
systemd-halt.service
systemd-poweroff.service
systemd-reboot,service
```

## 计算机的几种状态

```sh
logout              退出当前登录
suspend             挂起    硬盘、显示器断电, CPU, Memory仍然工作，数据保存内存
hybrid-sleep        睡眠    存储到内存和磁盘
hibernate           休眠    数据保存于硬盘中，CPU, Memory也停止工作
restart/reboot      重启
shutdown            关机
```

## cmd

```sh
shutdown -h now     立刻进行关机
shutdown -r now     现在重新启动计算机

halt                立刻进行关机
poweroff            立刻进行关机
reboot              现在重新启动计算机
```

shutdown -h now和shutdown -r now必须是root用户或者具有root权限的用户才能使用

而halt和reboot是Linux系统中的任何用户都可使用
