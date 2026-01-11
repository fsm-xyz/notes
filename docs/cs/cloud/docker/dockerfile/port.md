# port

## EXPOSE端口

[官网描述](https://docs.docker.com/engine/reference/builder/#expose)
容器构建者告诉使用者, 这个容器暴露的端口号有哪些, 但是并不公开，需要用户-p或者-P

1. 命令行--expose
2. 在Dockerfile中EXPOSE

## publish端口

+ -p
格式:
hostIP:hostPort:containerPort | hostIP::containerPort | hostPort:containerPort | containerPort
最后一种会在主机上随机分配端口

+ -P
将expose的端口随机对应到主机端口上
