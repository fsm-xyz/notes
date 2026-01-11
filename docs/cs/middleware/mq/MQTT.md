# MQTT

QoS

QoS 0（最多一次）	最低	发完即忘，可能丢失	环境传感器（容忍丢失）
QoS 1（至少一次）	中	确认重传，可能重复	设备状态上报（需保底到达）
QoS 2（恰好一次）	最高	四次握手保证唯一性	支付指令、关键控制（强可靠）


Clean Session = 1：连接断开后 Broker 丢弃会话信息（无状态）。

Clean Session = 0：Broker **保存订阅列表和未确认消息**（QoS>0），重连后恢复。

遗嘱消息（Last Will）

保留消息（Retained Message）
Broker 为 Topic 存储最新一条消息，新订阅者立即收到该消息（如设备最后一次状态）。


EMQX