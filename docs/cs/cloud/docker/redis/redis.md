# Redis

## 普通

```bash
docker run --name moon  -d redis
```

## 创建开启持久化的容器

```bash
docker run --name moon -d redis redis-server --appendonly yes
```

## 自定义的配置

### 新建Dockerfile

```bash
FROM redis
COPY redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
```

### 命令行挂载启动

```bash
docker run -v /myredis/conf/redis.conf:/usr/local/etc/redis/redis.conf --name myredis redis redis-server /usr/local/etc/redis/redis.conf
```

## 通过容器连接到对应容器的redis

```bash
docker run --rm --link moon:redis -it redis /bin/bash
docker run --rm --link moon:redis -it redis redis-cli -h redis -p 6379
```

## 直接宿主机连接(使用-p 指定的端口号)

```bash
redis-cli
```