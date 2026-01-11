# Kafka

## 概念

生产者: 批量发送, 单条发送确认  (批量条数，定时，异步发送，同步发送)

消费者：group消费，单次消费，批量消费 消费确认

broker

topic

offset

订阅

分区

副本


## 技术

1. 顺序IO
2. PageCache页缓存加速进行读写，异步刷盘，操作系统
3. 批量压缩处理(Snappy、LZ4、Zstandard)
4. 零拷贝，sendfile
5. 分区并行处理
6. 存储设计, log分段设置，由.log(数据)和.index(偏移量)组成

7. 生产者ack机制, 0 不需要确认, 1leader确认, all 需要副本确认
8. 组消费， 消费确认 手动提交（commitSync/commitAsync） vs 自动提交

## 高可用

+ 至少2个副本，就是三台机器组成集群
+ 幂等设计，消费者offset存储到目标系统

## zookeeper

旧版本使用zoookeeper保存broker和topic 消费记录，leaeder选举，监控broker存活信息

## 面试高频问题

1. Kafka 为什么快？
    顺序I/O + 页缓存 + 零拷贝 + 批量压缩 + 分区并行。

2. 如何保证消息不丢失？
    Producer：acks=all + 重试。
    Broker：min.insync.replicas>=2 + 禁用 Unclean Leader 选举。
    Consumer：关闭自动提交，处理完业务后手动提交偏移量。

3. 如何保证消息顺序？
    单分区内消息有序（同一 Key 的消息路由到同一分区）。

4. Rebalance 的触发条件？如何避免频繁 Rebalance？
    条件：消费者增减、订阅 Topic 变化、分区数变化。
    优化：调大 session.timeout.ms 和 heartbeat.interval.ms。

5. Kafka 与 RabbitMQ 的区别？
    维度	Kafka	RabbitMQ
    设计目标	高吞吐、日志流	复杂路由、低延迟
    消息模型	发布-订阅（持久化存储）	Queue/Exchange（内存/磁盘）
    顺序保证	分区内有序	单个队列有序
    协议	自定义二进制协议	AMQP

6. 选举过程
    zk watch
    KRaft 模式

CAP

AP, CP

## RAFT

每个follower会和leader保持心跳，如果心跳超时，会在随机时间内发起投票，把自己标记为Candicate, 自己的term任期加+1, 向其他的follower发起请求，给自己上票，重置自己的选举超时时间

选举消息: term + candidateID + lastIndex + lastTerm  

其他的follower收到请求，如果在当前选举周期没有投过，进行任期判断，如果大于自己的任期就同意，进行index和term比较，数据最新的最多的，同意

赢得选举：超过半数(N/2+1)
输掉选举: 如果在等待时间，别的比它更新，自动放弃
选举超时: 没有任何一个 Candidate 获得超过半数的选票, 每个 Candidate 会等待自己的选举超时计时器再次到期，然后递增任期号，发起新一轮的选举

## 问题

支持数据流很好，业务流不理想，支持的功能简单

RocketMQ
1. 异步刷盘，容易宕机丢数据
2. 延迟队列
3. 不支持消息过滤机制，业务筛选特定标签数据
4. 不支持事务消息，存在数据一致性问题(半事务消息)
5. 消息有序，未消费的数据进入死信队列
6. mmap