# golang-protobuf

## 功能

代码生成库
提供解析库

## When the .proto file specifies syntax="proto3", there are some differences

+ Non-repeated fields of non-message type are values instead of pointers.
+ Enum types do not get an Enum method.

GRPC
RPC

## 参数

-I 基于某个路径寻找依赖, 默认当前目录
--go_out
--inkerpc_out
--java_out

对应protoc-gen-go的cha'j

paths
plugins 加载对应的rpc插件

```shell
protoc -I . --go_out=plugins=grpc:. *.proto
```
