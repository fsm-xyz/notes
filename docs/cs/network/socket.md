# Socket

socket是在传输层和应用层之间的一个软件抽象层
传输层根据套接字标识进行多路复用和多路分解, 映射到对应的进程中

## UDP

udp 服务端通过一个bind把一个socket和端口号绑定，直接处理接收到的数据
一个udp套接字用一个二元组表示(目的ip, 目的port)，返回的时候使用源ip和port

## TCP

服务端通过bind进行端口绑定，该socket进行监听，
TCP是基于连接的，有新的tcp建立连接请求，新建一个套接字进行处理，标识一个tcp连接
四元组标识(ip, port, ip, port)，多个连接(多个套接字)共同使用用一个端口号

## MTU

链路层帧能承载的最大数据量

## MSS

## RTT

## 查看socke情况

```sh
while true;
do
        date;
        netstat -n | awk '/^tcp/ {++state[$NF]} END {for(key in state) print key,"\t",state[key]}'i;
        sleep 2;
done;

TIME=5;
while true;
do
    netstat -ant |grep 1433| awk '/ESTABLISHED|TIME_WAIT|LISTEN|CLOSE_WAIT/ {count[$6]++} END {for(s in count) {printf("%12s : %6d\n", s, count[s]); }}'; 
    echo -------------------; 
    sleep $TIME; 
done


ss -s

ss -tpan

sysctl -a | grep tw

```
tcpdump -nn -i eth0 port 80

arp -n
arp -d xxx
```