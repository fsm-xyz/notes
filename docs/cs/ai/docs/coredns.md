# CoreDNS

k8s集群默认是cluster.local域名，但是有些情况下有人安装的时候设置了别的域名，导致一些开源项目里面写的cluster.local域名就不能使用了

ACK默认安装的时候选择了就没法修改，所以必须想别的办法处理

## 重写 

为集群内服务设置“别名”（Rewrite）

```sh
.:53 {
    errors
    health
    # 将 api.mycompany.com 重写为集群内部服务域名
    rewrite name api.mycompany.com my-service.default.svc.cluster.local
    kubernetes cluster.local in-addr.arpa ip6.arpa {
       pods insecure
       fallthrough in-addr.arpa ip6.arpa
    }
    # ... 其他配置
}
```

## 静态

配置静态 IP 解析（Hosts）

```sh
.:53 {
    # ... 其他配置/etc/resolv.conf
    hosts {
        192.168.1.100 db.internal
        10.0.0.5      auth.legacy.system
        fallthrough
    }
    # ...
}
```

## 连接外部/私有 DNS 服务器（Stub Domains / Forward）

```sh
# 默认块保持不变
.:53 {
    errors
    health
    kubernetes cluster.local in-addr.arpa ip6.arpa {
       pods insecure
       fallthrough in-addr.arpa ip6.arpa
    }
    forward . /etc/resolv.conf
    cache 30
}

# 新增块：专门处理 corp.example.com 后缀
corp.example.com:53 {
    errors
    cache 30
    forward . 10.10.0.1
}
```

## 原生多域名配置（最干净、推荐）

```sh
.:53 {
    errors
    health
    # 增加cluster.local2后缀
    kubernetes cluster.local cluster.local2 in-addr.arpa ip6.arpa {
       pods insecure
       fallthrough in-addr.arpa ip6.arpa
    }
    forward . /etc/resolv.conf
    cache 30
}
```

## 问题原因

ACK安装的时候设置域名后缀导致的问题， `/etc/resolv.conf`这个文件

这个search域会自动添加后缀到短域名，最大五段

```sh
nameserver xxx.xxx.xxx
search default.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
```

## 总结

CoreDNS 配置：决定了 DNS 服务器能解析哪些全域名。

/etc/resolv.conf (search)：决定了客户端怎么把短域名变成全域名去询问 DNS 服务器。

如果你只改了 CoreDNS 没改 search，那你必须使用完整域名（FQDN）来访问新后缀的服务。
