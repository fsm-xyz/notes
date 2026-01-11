# keepalive

## tcp的keepallive

tcp的keepalive默认是关闭的，主要是探测tcp连接是否存活

当开启keepalive时，系统会在一定时间没有数据交互后，进行探测

`tcp_keepalive_time`空闲时长, 每次发送心跳的周期, 默认2小时
`tcp_keepalive_intvl`每次探测的间隔时间, 默认75s
`tcp_keepalive_probes`探测的总次数, 默认9次

## http的keep-alive

http的keep-alive主要是为了连接复用

```sh
Connection: Keep-Alive
Keep-Alive: timeout=5, max=1000
```

Connection 表示是否复用
Keep-Alive 空闲连接需要保持打开状态的最小时长（以秒为单位），在连接关闭之前，在此连接可以发送的请求的最大值，HTTP管道连接则可以用它来限制管道的使用。
