# 容器启动顺序

## 背景

在进行微服务改造的过程中，发现容器的重启次数比较多，查看k8s日志和业务内的日志，并没什么问题

当时的猜测是服务自身没有正常启动，但是直接基于k8s的是符合目标的

继续猜测如下

+ 业务容器启动快于proxy，健康检查导致重启
+ 没有正确的配置，服务依赖顺序


在前几天看版本更新的时候，看到新特性，支持proxy先于业务容器，于是就没怎么考虑第一个可能情况

在常规的业务排查和k8s排查之后，还是没确认原因

在分析k8s的event的时候，发现有时候会有一个健康检查的报错，但是有的没有，直接backoff

于是又把重心放到第一种情况，查看更新日志，同时也找到一个类似的情况，于是仔细观察更新日志，发现新特性默认不开启的

这时候发现还是需要仔细认真的品更新日志，难受住

## 分析

默认容器是按照顺序启动的，存在业务容器比proxy容器启动的快，系统进行健康检查，服务一直重启

历史的解决方案是让业务晚几秒启动，让proxy尽快启动

1.7版本新增`values.global.proxy.holdApplicationUntilProxyStarts`, 支持proxy在业务之前启动

k8s 1.18自有的Sidecar container， 目前istio使用的自有方案

## 原理

再注入容器时候会添加如下规则

```sh
lifecycle:
    postStart:
    exec:
        command:
        - pilot-agent
        - wait
```

## 如何操作

```sh
# 设置holdApplicationUntilProxyStarts为true
kubectl edit cm istio-sidecar-injector -n istio-system

# 查看生成的注入文件，需要上述规则
istioctl kube-inject -f services/ktv-server-svc/ktv-server-svc.yaml| grep lifecycle
```
