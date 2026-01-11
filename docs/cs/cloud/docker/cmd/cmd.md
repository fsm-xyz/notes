# 通用命令

## 查找images

```bash
docker search redis
```

## 镜像

```bash
docker pull redis
docker push name
docker images
docker rmi
docker commit
docker tag
```

## 查看状态

```bash
# 显示正在运行的
docker ps
# 显示所有的
docker ps -a
docker stats [container]
```

## 运行

```bash
docker run -d --name redis --net=host redis:4.0
```

## 停止

```bash
docker rm container
docker rm container -f
docker start container
docker stop container
docker kill container
```

## 进入容器

```bash
docker exec -it moon bash
docker exec -it -v /bin:/bin container bash
```

## 日志

```bash
docker logs
```

## 端口

```bash
docker port
```

## 帮助

```bash
docker help start
docker start --help
```

-t:在新容器内指定一个伪终端或终端。
-i:允许你对容器内的标准输入 (STDIN) 进行交互
-d:守护态运行
-v:宿主机目录映射到容器

--name 容器名称
-p 端口映射
--rm    运行结束后删除该容器
--link  连接到某个容器, 后面的redis注入到hosts文件, 指向前面的容器
--net
--network

## Redis

```sh
#!/bin/bash
# 自动化脚本练习
name=nginx
port=6379
for i in {1..6}
do
    cmd="docker run -d -p $port$i:$port --name $name$i $name"
    # cmd="docker start $name$i"
    echo $cmd
    $cmd

done
docker ps

docker ps -a | awk '{print $1}'| while read line; do docker rm -f $line;done
```


## CMD执行文件动态获取

```sh
ARG cmd=default
ENV cmd ${cmd}

CMD ["/bin/sh", "-c", "./$cmd"]
```

## 删除镜像
```sh

docker image prune

docker image ls | grep -v TAG | awk '{print $1":"$2}' | xargs docker rmi
```

## 删除容器

```sh
docker container prune
````