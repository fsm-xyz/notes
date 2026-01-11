# 编译docker

## 自定义方式

利用go11的module功能可以方便解决依赖问题

以下是在v0.3.0基础上自定义编译

```go
git clone git@github.com:moby/moby
cd moby
git checkout v0.3.0
go mod init github.com/dotcloud/docker
go mod tidy
cd docker
go build
```

由于docker公司改名, 根据不同的文件内容选择不同的uri

```shell
go mod init github.com/dotcloud/docker
go mod init github.com/docker/docker
```

## 启动

旧版本docker一个执行文件, 根据-d参数判断是daemon还是cli
新版本dokcerd, docker

```shell
./docker -d
```

## 官方Makefile

暂时只看了下Makefile, 修改GOPATH等, 未实践编译

## 报错

mac机器上由于实现不同所以docker0网桥不显示
[参考](https://docs.docker.com/docker-for-mac/networking/#there-is-no-docker0-bridge-on-macos)

## build动态传入参数

`docker build -t yellow:2.0 --build-arg envType=dev .`

```sh
ARG envType=default_value
ENV envType ${envType}
```
