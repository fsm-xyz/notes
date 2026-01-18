# Module proxy protocol

[Module proxy protocol](https://golang.org/cmd/go/#hdr-Module_proxy_protocol)

任何实现如下GET接口的Web服务器都可以当做module proxy
请求没有任何参数, 所以file:///URL 都是可以的

```sh
# 获取已知的module版本
GET $GOPROXY/<module>/@v/list

# 获取module某个版本信息
GET $GOPROXY/<module>/@v/<version>.info

# 获取module某个版本的mod信息
GET $GOPROXY/<module>/@v/<version>.mod

# 获取module某个版本文件
GET $GOPROXY/<module>/@v/<version>.mod

# 获取最新的(如果list没有合适的版本)
GET $GOPROXY/<module>/@latest
```

流程
list -> [latest] -> /info -> mod -> zip

## 大小写

为了避免大小写，大写字母会转化， A-> !a

## 代理

+ [goproxyio](https://goproxy.io)
+ [goproxycn](https://goproxy.cn)
