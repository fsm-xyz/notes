# IO

## IO模型

* 阻塞IO
* 非阻塞IO
* 多路复用(事件驱动)
* 信号驱动IO
* 异步IO

### 核心操作

recvfrom

## TCP/UDP

### UDP

报文格式

-------32bit--------
源端口号    目的端口号
长度        校验和
应用数据(报文)

### TCP

三次握手

server  half open (端半打开), DDOS 攻击

client          server

SYN_SENT
                SYN_RCVD
ESTABLISHED
                ESTABLISHED

四次断开

client          server

FIN1
                CLOSE WAIT
FIN2

                LAST ACK
TIME WAIT
                CLOSED

### 报文格式

-------32bit--------
源端口号    目的端口号
序号
确认号
首部长度    接收窗口
校验和      紧急数据指针
选项
数据

### 可靠性传输

停等
滑动窗口协议(回退N步)
选择重传

### 拥塞控制

拥塞窗口 cwnd

### 拥塞控制算法

1. 慢启动
2. 拥塞避免
3. 快速恢复

### 流量控制服务

### 关键

TIME WAIT

MSL maximum segment lifetime

MSS maximum segment size

MTU maximum transmission unit

RTT
