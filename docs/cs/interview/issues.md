# issues

## golang无法自动识别容器资源限制

    手动设置CPU和内存
    使用user的automaxproces库

## 网络异常  nf_conntrack: table full, dropping packet
    设置对应的连接时间

## 容器就绪，存活探针

k8s服务不定时报503
    ruby的修改keepalive timeout参数，解决
    golang默认不过期，是由于k8s退出时，同时发送sigterm和更改endpoints列表，异步执行，导致流量到了关闭的pod上，增加lifecycle的prestop即可

代理启动比业务容器慢，导致服务重启次数过多
    values.global.proxy.holdApplicationUntilProxyStarts

非mesh环境，k8s grpc访问不均衡
    使用mesh代理，实现访问
    使用外部的服务发现机制

证书的存放问题

监控获取

链路传递

日志打印

envoy和nginx的路由规则问题 (rest)

DNS解析异常(修改options数量，设置超时，设置autopath @kubernetes)

有状态使用svc即可，不需要pod.svc访问，避免异常

grpc的自定义数据传递

grpc低版本的链式传递

业务容器和代理容器同时设置探针，导致Unhealthy变多

    https://istio.io/latest/docs/ops/configuration/mesh/app-health-check/

istio默认rewrite proxy的readiness到业务容器，导致超时
    设置禁用即可rewriteAppHTTPProbe: false

打印body遇到报错
    content:[libprotobuf ERROR external/com_google_protobuf/src/google/protobuf/wire_format_lite.cc:578] String field 'google.protobuf.Value.string_value' contains invalid UTF-8 data when serializing a protocol buffer. Use the 'bytes' type if you intend to send raw bytes.
    https://github.com/envoyproxy/envoy/issues/9822
    https://github.com/envoyproxy/envoy/pull/13131

context的超时控制(业务还是代理设置)

打通metrics，log，trace

开关打印，出错了打印

### 业务监听podip，服务无法访问，0.0.0.0则可以，

历史版本直接把所有流量转发给lo设备，导致监听podip的服务无法接受到访问

127.0.0.1 -> lo
podip -> eth0

1.10修改转发规则，实现正常的逻辑

https://istio.io/latest/blog/2021/upcoming-networking-changes/
