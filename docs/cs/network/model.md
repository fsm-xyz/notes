# OSI模型和tcp/ip模型

## OSI模型

应用层          HTTP FTP TFTP SMTP SNMP DNS TELNET HTTPS POP3 DHCP
表示层          JPEG、ASCll、DECOIC、加密格式
会话层          对应主机进程，指本地主机与远程主机正在进行的会话
传输层          定义传输数据的协议端口号，以及流控和差错校验。协议有：TCP UDP，数据包一旦离开网卡即进入网络传输层
网络层          协议有：ICMP IGMP IP（IPV4 IPV6） ARP RARP
数据链路层      建立逻辑连接、进行硬件地址寻址、差错校验 [2]  等功能。（由底层网络定义协议）将比特组合成字节进而组合成帧，用MAC地址访问介质，错误发现但不能纠正。
物理层          中继器、集线器、还有我们通常说的双绞线也工作在物理层，ISO2110，IEEE802，IEEE802.2

## TCP/IP模型

应用层
传输层
网络层
主机层

## TCP/UDP协议

TCP(Transmission Control Protocol)和UDP(User Datagram Protocol)协议属于传输层协议。

1. 其中TCP提供IP环境下的数据可靠传输，它提供的服务包括数据流传送、可靠性、有效流控、全双工操作和多路复用。通过面向连接、端到端和可靠的数据包发送。通俗说，它是事先为所发送的数据开辟出连接好的通道，然后再进行数据发送；而UDP则不为IP提供可靠性、流控或差错恢复功能。
2. 一般来说，TCP对应的是可靠性要求高的应用，而UDP对应的则是可靠性要求低、传输经济的应用。
3. TCP支持的应用协议主要有：Telnet、FTP、SMTP等；UDP支持的应用层协议主要有：NFS（网络文件系统）、SNMP（简单网络管理协议）、DNS（主域名称系统）、TFTP（通用文件传输协议）等。

## TCP三次握手,四次断开

CLOSE-WAIT
TIME-WAIT
四次挥手原因：
可靠地实现TCP全双工连接的终止

出现大量TIME-WAIT
原因：
使用短连接，完成一次请求后会主动断开连接，就会造成大量time_wait状态，TIMEWAIT状态持续几分钟

危害：
服务端的话, 有可能会资源(socket, 端口)耗尽

解决方案：

1. nginx开启了长连接keepalive
2. go里面可以设置DefaultMaxIdleConnsPerHost复用的最大个数
