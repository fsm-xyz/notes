## SDK客户端架构

1. 优化keepalive心跳(客户端和服务端都进行心跳检测，如果有消息发送可以取消本次的心跳检测，减少发包)
2. 使用SurealDB替换SQLlite实现实时数据同步(Live Queries, Subscribable APIs)

接入层，协议层，业务层，存储层