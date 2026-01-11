# protobuf

protobuf(ProtocolBuffer)是谷歌开源的一个二进制序列化协议，性能高，支持多语言

## 基本原理

对收到的数据以byte进行解析, 采用小端字节序
比如程序里面int是4字节，使用varint编码可以压缩为1-5字节，负数使用`Zigzag`进行编码

T-V
T-L-V
  
+ type 字段类型
  + 0  Varint 变长(1-10字节) int32, int64, uint32, uint64, bool, enum, sint32, sint64
  + 1  64bit  8字节 fixed64, sfixed64, double
  + 2  Length-delimi 变长  string, bytes, embeded messages, packed...
  + 3  Start group 已弃用
  + 4  End group  已弃用
  + 5  32bit  4字节 fixed32, sfixed32, float

例子

```proto
message Test3 {
  optional Test1 c = 3;
}
```

使c = 150
十六进制: 1a 03 08 96 01
二进制: 0001 1010 0000 0011 0000 1000 1001 0110 0000 0001

0标志位  0011字段编号 010字段类型 0000 0011字段长度

标志位，0表示这是最后一个字节，1表示下一个字节也是

[参考](https://developers.google.cn/protocol-buffers/docs/encoding)
