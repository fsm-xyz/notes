
接入层:
    全局负载均衡: DNS + CDN
    API网关：Kong, APISIX, Spring Cloud Gateway
        认证鉴权、限流熔断、日志监控
        集群网关的高并发 nginx, envoy， harporxy

                高可用 lvs和 keepalived+ VIP

                多集群

                多可用区（AZ）+ 异地多活
业务服务层：
    微服务架构 + RPC， 服务注册-发现， 配置中心
    服务网格：Istio（流量管理、安全策略

数据层
    主从(读写分离)
    集群 (分库分表)

    分布式数据库    数据分片+副本集

    大数据分析	HBase + Spark, ClickHouse	列式存储+分布式计算

    消息队列	Kafka, Pulsar, RabbitMQ集群	分区复制+持久化


可观察性: metrics. logs. traces

高可用: 健康检查, 故障隔离，水平伸缩

重试, 超时控制, 熔断

限流 熔断  

弹性伸缩

分布式事务：

    Seata（AT/TCC模式）

服务级调度	Istio VirtualService + DestinationRule	按Header将VIP用户导流到独立泳道
全局调度	DNS分地域解析 + CDN边缘计算	用户就近接入上海/深圳机房
集群调度	Kubernetes Ingress + Service	根据Node资源利用率调度Pod

主动降级	配置中心推送降级开关	大促期间关闭非核心服务
熔断器	Sentinel熔断规则（慢调用比例）	支付服务RT>1s且错误率>50%
流量整形	Redis令牌桶限流	秒杀场景限制QPS=1000
资源隔离	Kubernetes Namespace + Quota	隔离测试环境与生产环境资源

网络隔离、故障注入


在线服务(延迟敏感) + 离线任务(吞吐优先)混部