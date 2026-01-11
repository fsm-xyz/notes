# 网络丢包

## 现象

2台内网机器之间丢包严重

## 可能原因

+ 本身网络之间丢包
+ 通讯双方机器CPU满导致的丢包

## 结果

由于对方机器CPU满，导致丢包

### CPU导致的丢包

CPU处理网卡数据负载不均衡

软中断处理网卡数据


### 开启IRQ balance

+ `https://blog.csdn.net/whrszzc/article/details/50533866`
+ `https://yq.aliyun.com/articles/611355`
