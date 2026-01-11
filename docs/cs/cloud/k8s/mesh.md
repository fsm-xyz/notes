
## service mesh

linkerd, istio

### 阶段1

proxy

早期通过简单的proxy进行实现服务治理

控制面

### 阶段2

sidecar

通过把proxy和业务进程绑定到一起，形成一个sidecar模式

控制面

### 阶段3

在原来数据面的基础上，增加控制面

istio

### istio实现

分为数据面和控制面

可以基于k8s, vm， consul等，istio使用适配进行服务转换

### 控制面

+ pilot组件实现服务的注册，发现等
+ citadel组件实现服务间的证书分发
+ mixer组件实现遥感数据收集和策略执行
+ galley负责配置校验

Istio 1.5 重建了控制平面，回归单体架构，废弃mixer, 部分逻辑由proxy实现，多个组件服务合并到一个istiod服务

### 数据面

默认使用envoy做proxy，数据来自pilot的数据

### 功能

+ 服务注册，发现，负载均衡
+ 限流
+ 熔断
+ 跨域
+ 灰度发布
+ 重试
+ 重定向
+ 分布式追踪

### bookinfo的展示

+ grafana
+ kiali
+ prometheus
+ jaeger

各种功能的演示
