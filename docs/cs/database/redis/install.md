# Redis

## 源码安装

```bash
# 下载
wget http://download.redis.io/releases/redis-4.0.10.tar.gz
tar xzf redis-4.0.10.tar.gz
cd redis-4.0.10
# 安装依赖
yum install -y jemalloc
yum install -y tcl
# 编译
make
make clean
# 验证
src/redis-server
src/redis-cli

```

## 开机启动服务

1. yum安装会自动创建*.service
2. 手动源码安装需要自己添加/lib/systemd/system/*.service
3. 文件格式
[Unit]
Description=redis - high performance kv db
After=network-online.target remote-fs.target nss-lookup.target
Wants=network-online.target

[Service]
Type=forking
ExecStart=/bin/redis-server /etc/redis/redis.conf
ExecStop=/bin/redis-cli shutdown

[Install]
WantedBy=multi-user.target

1. 添加
systemctl enable redis
systemctl start redis
systemctl status redis

## 远程访问配置

redis.conf
    > +logfile ""
    > +bind
    > +protected-mode
    > +requirepass

