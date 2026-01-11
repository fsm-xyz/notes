# DNS

## resolv.conf

Multiple nameservers in resolv.conf are not a means to do unioning of conflicting DNS namespaces. They're expected to be purely redundant with non-conflicting (i.e. if one doesn't know about something another does, it has to ignore the query or ServFail, not NxDomain or NODATA
it) records. If you need unioning of distinct spaces using custom rules for resolving conflicts, you need a special nameserver running on localhost or somewhere else you control that performs this logic.

## 改进

+ localdns(本地dns)
+ smartdns(返回连接最快的)
+ namasq(并发解析，使用最快解析的)

## 趣事

https://www.openwall.com/lists/musl/2023/01/10/1
https://www.openwall.com/lists/musl/2023/01/10/2
https://www.openwall.com/lists/musl/2023/01/10/3
https://www.openwall.com/lists/musl/2023/01/11/1
https://www.openwall.com/lists/musl/2023/01/11/2
https://www.openwall.com/lists/musl/2023/01/11/3
https://www.openwall.com/lists/musl/2023/01/11/4

