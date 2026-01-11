# redis

+ 主从集群
+ 分片集群

## Redis cluster

Redis Cluster是一种服务器Sharding技术，3.0版本开始正式提供。
Redis Cluster中，Sharding采用slot(槽)的概念，一共分成16384个槽，这有点儿类pre sharding思路。对于每个进入Redis的键值对，根据key进行散列，分配到这16384个slot中的某一个中。使用的hash算法也比较简单，就是CRC16后16384取模。

## Redis Sharding集群

Redis Sharding可以说是Redis Cluster出来之前，业界普遍使用的多Redis实例集群方法。其主要思想是采用哈希算法将Redis数据的key进行散列，通过hash函数，特定的key会映射到特定的Redis节点上。这样，客户端就知道该向哪个Redis节点操作数据。

## 分片方案

哈希分片(一致性hash)

### 客户端分片

使用一个基础库进行分片，在同一个进程内执行

优点: 性能高
缺点: 暴露细节太多, 一旦分片规则发生变化, 得所有的业务依赖都进行更改

### proxy分片

分片逻辑在这个proxy服务端实现

优点: 屏蔽技术细节, 方便变更分片规则, 跨语言
缺点: 多了网络IO和服务器消耗

### 服务端分片

redis本身支持的方案

## 中间件

+ codis
+ twemproxy

## 数据迁移

全量rehash
