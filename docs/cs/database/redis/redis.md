# Redis

## 基础数据类型

+ string
+ list
+ set
+ zset
+ hash

## 查看帮助

```sh
help
```
## 基本信息

+ 基于内存的键值数据库
+ IO多路复用(epoll)
+ 使用高效数据结构，无锁单线程操作
+ IO多线程

## 执行

在redis server中有两个循环：IO循环和定时事件。
在IO循环中，redis完成客户端连接应答、命令请求处理和命令处理结果回复等。
在定时循环中，redis完成过期key的检测等。
redis一次连接处理的过程包含几个重要的步骤：IO多路复用检测套接字状态，套接字事件分派和请求事件处理。

## Redis知识点


    https://zhuanlan.zhihu.com/p/32540678
    https://zhuanlan.zhihu.com/p/34133067
