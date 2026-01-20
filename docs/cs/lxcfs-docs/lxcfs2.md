# lxcfs

lxcfs是用户态方案，在宿主机上运行

引入这个解决容器内部的资源视图是宿主机器的，正确挂载容器配额到内部的/proc

各种库调用的是宿主机的/proc目录下的信息，k8s底层容器是cfs，导致程序误判资源，启动的线程等不合理

## 问题

cpu限流

## 实现

### 启动lxcfs

```sh
sudo mkdir -p /var/lib/lxcfs
sudo lxcfs /var/lib/lxcfs --enable-cfs
```

### docker挂载

```sh
docker run -it --cpus=2 -m 256m \
      -v /var/lib/lxcfs/proc/cpuinfo:/proc/cpuinfo:ro \
      -v /var/lib/lxcfs/proc/diskstats:/proc/diskstats:ro \
      -v /var/lib/lxcfs/proc/meminfo:/proc/meminfo:ro \
      -v /var/lib/lxcfs/proc/stat:/proc/stat:ro \
      -v /var/lib/lxcfs/proc/swaps:/proc/swaps:ro \
      -v /var/lib/lxcfs/proc/uptime:/proc/uptime:ro \
      -v /var/lib/lxcfs/proc/slabinfo:/proc/slabinfo:ro \
      -v /var/lib/lxcfs/sys/devices/system/cpu:/sys/devices/system/cpu:ro \
      ubuntu:18.04 /bin/bash

# 容器内部执行top等，看到就是正确的cpu和内存
```

### K8s环境


+ DaemonSet
+ 挂在宿主机目录(手动)
+ ServiceAccout(自动化)

[阿里云国际版lxcfs](https://www.alibabacloud.com/blog/kubernetes-demystified-using-lxcfs-to-improve-container-resource-visibility_594109?spm=a2c65.11461478.0.0.47cf5355LtkBmb)

## 替代方案

腾讯的CgroupsFS内核态方案

[文档连接](https://github.com/OpenCloudOS/TencentOS-kernel-0?tab=readme-ov-file#%E5%AE%B9%E5%99%A8%E9%9A%94%E7%A6%BB%E5%A2%9E%E5%BC%BA)