# istio

## 配置修改

```sh
# 设置业务容器必须得等proxy就绪后启动
 --set meshConfig.defaultConfig.holdApplicationUntilProxyStarts: true
# 修改polit的资源占用(默认值太大，调试的时候调小点)
 --set values.pilot.resources.requests.cpu=200m
```

## envoy

### headers

#### grpc 请求失败

grpc请求缺少`:authority`字段导致被envoy proxy拒绝掉

原因: proxy 嗅探到这次出口流量为http2, 检查`:authority`字段不存在, 就返回错误，参考etcd自身实现, 就可以发现原因

排查过程: 搭建etcd服务, istio环境, 分别查看服务的istio-proxy容器的日志输出, 看是否有流量进来

```sh
# 代理输出
kubectl logs -f xxx -c istio-proxy
# 业务容器输出
kubectl logs -f xxx

# 日志输出
proxy的日志里面有个报错`http2.invalid.header.field`
```

服务里面etcd请求etcd服务端也是grpc请求，日志显示正常
业务里面通过etcd的grpc业务请求失败，业务报错`http2.invalid.header.field`

对比发现错误的请求里面缺少`:authority`的值

##### 对比

```sh
# 业务etcd历史方式
target := fmt.Sprintf("%s:///", rb.Scheme()),
# 修改后
target := fmt.Sprintf("%s:///%s", rb.Scheme(), "123"),
# etcd写法
target := fmt.Sprintf("%s://%p/%s", resolver.Schema, c, authority(c.endpoints[0]))
# grpc
fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName)
# dial option
WithAuthority()
```

https://github.com/etcd-io/etcd/blob/bf5c936ff1de422b48cc313435aa40ef6f2057ac/client/v3/client.go#L306
https://github.com/grpc/grpc-go/blob/57aaa10b8a9e575f4834f19fa63c0d2e184f372e/clientconn.go#L259
https://imroc.cc/istio/faq/headless-svc.html
https://blog.csdn.net/u013536232/article/details/108556544

### envoy不能打印request body
    使用EnvoyFilter实现

### envoy默认的上传文件大小限制 http code 413
    使用EnvoyFilter实现

## 网络

### 1.10网络

istio1.10之前存在, 自动把eth0的流量转发给lo，导致之间听eth0 或者lo 不能接收流量

详细描述: https://zhonghua.io/2019/07/11/istio-xds-podip/

https://istio.io/latest/blog/2021/statefulsets-made-easier/
https://istio.io/latest/blog/2021/upcoming-networking-changes/
