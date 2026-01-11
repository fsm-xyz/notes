# 本地开发使用consul

## consul

consul使用client代理本机器上所有的服务发现和注册请求

### 端口号

| 端口号 | 协议 | 功能 |
|:--|:--|:--|
8300 | TCP     | agent server 使用的，用于处理其他agent发来的请求
8301 | TCP/UDP | agent使用此端口处理LAN中的gossip
8302 | TCP/UDP | agent server使用此端口处理WAN中的与其他server的gossip
8400 | TCP     | agent用于处理从CLI来的RPC请求
8500 | TCP     | agent用于处理HTTP API
8600 | TCP/UDP | agent用于处理 DNS 查询

### 误区

* 一开始我以为consul client是代理全部流量(业务流量和服务发现等)，类似Envoy

## ZK

zk可以本地直接连接到服务注册中心

## 框架

框架可以提供一个参数，是否把本地服务注册到registry
业务方通过设置参数进行本地服务启动

## 本地环境搭建

### 下载

[下载](https://www.consul.io/downloads.html)
[安装] 把二进制文件放到$PATH下面

### 启动

```sh
#!/bin/sh

consul agent -config-file=./conf/config.json -node=yourhostname
```
