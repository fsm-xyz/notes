# debug

## strace

查看程序跟内核的交互

```sh
strace -ff -o log ./redis-server
```

## nc

查看命令交互细节
```sh
nc 127.0.0.1 6379
```

## exec

创建一个连接
```sh
ls /proc/$$/fd
# 创建套接字
exec 8<> /dev/tcp/localhost/6379
echo -e 'set a 1' 1>& 8

cat 0<& 8

# 销毁
exec 8<& - 
## benchmark

```sh
redis-benchmark -t set -c 50
```
