# Istio

从一体化代码到基础框架，再到中间件处理，实现架构的抽象

## 功能

+ 可观测性
+ 流量路由、负载均衡、安全加密
+ 限流，熔断，灰度，AB测试
+ 网关
+ 故障恢复：超时（timeout）、重试（retries）、熔断（CircuitBreaker）

## 组件

架构从多个微服务合并到一个单一服务多模块

poloit(istiod)

### 实现

+ Gateway
+ VirtualService
+ DestinationRule
+ mTls
+ Hbone
+ redirect
+ tproxy
+ epef
+ xds

## 流程

### Envoy 作用

流量代理：拦截进出 Pod 的流量，实现负载均衡（轮询/随机）和故障注入

安全网关：终止 TLS 连接，执行 JWT 验证
案例：某服务超时 → 调整 Envoy 的 timeout 参数优化响应

### Pilot 实现服务发现

监听 K8s API Server 获取 Service/Endpoint 变更。

转换规则为 Envoy 配置（XDS 协议）

通过 gRPC 流将配置推送给 Sidecar

### 金丝雀发布配置

定义 DestinationRule 声明版本子集（subsets: v1, v2）。

VirtualService 按权重分流（如 v1:90%, v2:10%）。

监控 v2 错误率 → 逐步调高权重68。

### 指标数据

指标收集：Envoy 直接上报 Prometheus

多集群

ingressgateway


ambient

ztunnel     DaemonSet 

iptables/CNI

waypoint

HTTP/gRPC 路由、流量拆分、熔断限流、JWT 认证等 L7 策略

按 ServiceAccount 或 Namespace 粒度伸缩。
