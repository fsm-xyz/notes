# http

## 概念

请求行 请求头  请求体

状态行 响应头  响应体  

http1,http2,http3

1.0是短连接, 1.1加入keep-alive和pipeline, 2.0多路复用(基于流的数据帧，单连接并行处理)解决pipeline的队头阻塞, 3.0基于QUIC解决TCP的队头阻塞

host
connection: close用来表示需要关闭连接, keep-alive表示需要保持连接
content-type
content-length


缓存: Cache-Control， Expires, ETag, Last-Modified

Cookie和Session

https加密流程

TCP三次握手

TLS握手（非对称加密交换密钥）

对称加密传输数据

http Upgrade和alpn和NPN


### http3

1 RTT（首次），0-RTT（恢复）

quic 应用层自己解决数据重传

连接迁移 connection id

0 Rtt

content-encoding

压缩

对称加密，非对称加密

短连接，长连接

## 常见问题


跨域

http1的

keepalive

pipeline流水线

队头阻塞(基于TCP的重传和http1.0的pipeline队头阻塞)

http 下载怎么知道文件结束

1xx: 信息性（如101协议切换）

2xx: 成功（200 OK、201 Created）

3xx: 重定向（301永久移动、302临时移动、304 Not Modified）

4xx: 客户端错误（400 Bad Request、401 Unauthorized、403 Forbidden、404 Not Found）

5xx: 服务端错误（500 Internal Server Error、502 Bad Gateway 503 服务不可用）


2. HTTP/2核心优化
二进制分帧（Binary Framing）

头部压缩（HPACK算法）

服务器推送（Server Push）

多路复用（Multiplexing）

3. HTTP/3革新
基于QUIC协议（UDP实现可靠传输）

解决TCP队头阻塞问题

0-RTT快速建立连接
