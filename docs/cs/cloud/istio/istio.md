# istio

用来做服务之间的治理

连接，安全，控制，观察

## 核心概念

+ VirtualService

   用来做服务之间的流量治理
   + 路由
   + 重定向
   + 重写
   + 重试
   + 故障注入(延时，错误码)
   + 镜像
   + 跨域

+ DestanationRule

    做一个服务下的endpoint的负载均衡和熔断等

+ Gateway

  描述网关收到流量如何分发，对应一个VirtualService
  类比Ingress和IngreesController
    Ingress同时描述服务入口和后端路由，注册到IngressController
    IngreesController是具体的实现(Nginx)
  Gateway描述服务的访问，而服务的路由定义在VirtualService，解藕定义和路由规则

## 核心组件

+ Pilot
+ Galley
+ Citadel

## 排查流量

```sh
kubectl logs istio-ingressgateway-xxx -c istio-proxy
```