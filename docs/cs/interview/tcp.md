# TCP

## 概念

三次握手

SYN_SENT            SYN=1, seq=x
                                                        SYN=1, ACK=x+1, seq=y   SYN_RCVD
ESTABLISHED         ACK=1, ack=y+1

防止历史连接初始化
确保双方收发能力正常

四次挥手

FIN_WAIT_1      FIN, seq=u
                                            CLOSE_WAIT
                            ack=u+1
FIN_WAIT_2

                            FIN=1, seq=w    LAST_ACK
TIME_WAIT
                ACK=1, ack=w+1              CLOSED

全双工

SYN（建连）、ACK（确认）、FIN（断开）、RST（重置）、PSH（推送数据）、URG（紧急指针）

## 可靠性

超时重传
快速重传
滑动窗口

TCP或周期检查连接是否存活, keepalive (内核中可以设置)

TCP是面向连接的, 可靠的传输

UDP是面向数据包的传输

最大传输速度

## 流量控制

Window Size（窗口大小）

MTU（最大传输单元）：数据链路层限制（如以太网 1500 字节）。

MSS（最大段大小）：TCP 层限制（= MTU - IP头 - TCP头，通常 1460 字节）。

2MSL

拥塞控制.滑动窗口

### 拥塞控制

慢启动

拥塞避免

快重传/快恢复


TCP 如何保证可靠性？
→ 校验和 + 序列号/ACK + 重传 + 流量控制 + 拥塞控制。

TIME_WAIT 太多怎么办？
→ 优化方案：

调整内核参数（如 net.ipv4.tcp_tw_reuse=1 复用 TIME_WAIT 连接）。

程序层设置 SO_REUSEADDR 选项。

粘包/拆包问题如何处理？
→ 根本原因：TCP 是字节流协议，无消息边界。
→ 解决方案：

固定长度消息（效率低）。

分隔符（如 \n）。

消息头声明长度（如 HTTP 的 Content-Length）。

SYN Flood 攻击原理与防御？

原理：伪造大量 SYN 请求耗尽服务端资源（半连接队列满）。

防御：

SYN Cookie（不存储连接状态）。

限制 SYN 速率（iptables）。

增大半连接队列。

为什么 UDP 有时比 TCP 快？

无连接建立/断开开销。

无拥塞控制（可定制传输策略）。

无重传机制（适合容忍丢包的场景）。

net.ipv4.tcp_tw_reuse
net.ipv4.tcp_tw_recyle