# kafka

## 组件

+ broker
+ consumer, consumer group
+ partition
+ 多副本(replication)
+ topic
+ producer

## 模式

生产者push到broker, 消费者pull数据

## 消息队列相关(rabbitmq, activemq, rocketmq, kafka)

## 生产方式

+ 同步发送
+ 异步发送

acks：0, 1, n, -1

## 消费方式

+ 同步消费
+ 异步消费

## 作用

+ 业务解耦
+ 流量削峰

## 数据高可靠性

+ 多副本
+ 同步发送
+ 同步消费


## 技术

+ 批量处理
+ 顺序读写
+ zero copy
+ page cache

## 消息丢失，重复

+ 生产端消息丢失
+ 消费端消息丢失
+ 消费端消息重复
