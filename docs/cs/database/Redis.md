# Redis

单线程模型

## RDB

快照


手动 SAVE（阻塞）/ BGSAVE（后台）

## AOF

刷盘策略：

配置	数据安全	性能
appendfsync always	最高（每条刷盘）	最低
appendfsync everysec	适中（每秒刷盘）	中等（默认）
appendfsync no	最低（OS决定）	最高
重写机制：BGREWRITEAOF 压缩命令（消除冗余）

## 主从

Master+Slave

Sentinel

## 集群

proxy机制
分片机制：16384 个槽（Slot），每个节点负责部分槽数据
路由：客户端计算 CRC16(key) % 16384 定位槽

Gossip


缓存穿透        查询不存在的数据            空值缓存,布隆过滤器拦截

缓存雪崩        大量 key 同时过期           随机过期时间        集群部署分散风险

缓存击穿        热点 key 过期瞬间高并发查询     永不过期 + 后台更新

## 问答

Q：为什么 Redis 单线程还快？

✅ 答：

内存操作（纳秒级）

单线程避免上下文切换/锁竞争

IO 多路复用（epoll）处理高并发连接

6.0 后网络 IO 多线程提升吞吐（命令执行仍单线程）

client cache

autopipeline
