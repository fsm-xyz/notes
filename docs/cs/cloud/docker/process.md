# 进程查看

宿主机是可以看到容器内部开启的进程的

## 查看 pid

```sh
ps -ef | grep api-server

docker ps
docker top container

PID 进程 id, PPID 父 ID
```

## 查找容器 ID

```sh
docker inspect -f "{{.Id}}" \$(docker ps -q) |grep PID

docker inspect -f "{{.Id}} {{.State.Pid}} {{.Config.Hostname}}" \$(docker ps -q) |grep PID

for i in `docker ps |grep Up|awk '{print $1}'`;do echo \ &&docker top $i &&echo ID=$i; done |grep -A 10 `<PID>`
```

## 排查

```sh
# 设置启动命令 使容器启动
tail -f /dev/null
```
