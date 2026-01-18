**不需要。**

这是一个非常重要的误区。**方案 A（域名后加点）完全不需要修改 CoreDNS 或 LocalDNS 的任何配置。**

这个优化动作是发生在 **客户端（你的 Pod）** 内部的，确切地说是发生在 Pod 的 **操作系统层（glibc/musl）**。

以下是详细的原理解析，帮助你理解为什么服务端（DNS Server）不需要动：

### 1\. “点”是给谁看的？

那个末尾的点（`.`），是给 **Pod 内部的 DNS 解析器**（Resolver，通常是操作系统自带的库）看的信号。

  * **没有点：** 操作系统解析器觉得这是一个“相对路径”，它会很勤快地帮你把 `/etc/resolv.conf` 里的 `search` 后缀一个个拼上去试。
  * **有点：** 操作系统解析器看到末尾有点，就知道这是“绝对路径”（FQDN），它会**立即停止**任何拼接尝试，直接把这个域名发出去。

### 2\. CoreDNS 实际收到了什么？

无论你是否在代码里加点，**CoreDNS 最终收到的成功请求是一模一样的**。区别仅在于“废话”请求有没有发出来。

让我们模拟一下请求 `google.com` 的过程：

#### 场景一：你没加点（`google.com`），且 CoreDNS 没改配置

1.  **Pod 内部：** 发现没点，开始根据 search 域拼接。
2.  **Pod -\> CoreDNS：** 发送 `google.com.default.svc.cluster.local.`
      * **CoreDNS：** 查不到，返回 NXDOMAIN（不存在）。
3.  **Pod -\> CoreDNS：** 发送 `google.com.svc.cluster.local.`
      * **CoreDNS：** 查不到，返回 NXDOMAIN。
4.  ...（重复几次无效交互）...
5.  **Pod -\> CoreDNS：** 最终发送 `google.com.`
      * **CoreDNS：** 查到了！返回 IP。

#### 场景二：你加了点（`google.com.`），CoreDNS 依然没改配置

1.  **Pod 内部：** 发现有点，跳过拼接。
2.  **Pod -\> CoreDNS：** 直接发送 `google.com.`
      * **CoreDNS：** 查到了！返回 IP。

**结论：**
在场景二中，CoreDNS 只是少处理了前几次垃圾请求，直接处理了最后那次正确的请求。**它不需要做任何特殊配置来“支持”带点的域名**，因为 DNS 协议标准里，所有域名的本质都是带点的。

### 3\. 为什么会产生这种误解？

有些用户可能会混淆 **CoreDNS 的 Rewrite 插件** 和这个优化。

  * **什么时候需要改 CoreDNS？**
    如果你想让 `google.com` 解析到你集群内部的某个 Service IP（做域名劫持或重写），那你需要在 CoreDNS 里配置 `hosts` 或 `rewrite`。
  * **什么时候不需要改？**
    如果只是为了加速解析外部域名（做性能优化），只要在客户端（代码/配置）加点即可。

### 4\. 如何验证？

你可以直接进入一个 Pod 内部，用 `nslookup` 验证一下，完全不需要动服务器配置：

```bash
# 进入一个 Pod
kubectl exec -it <pod-name> -- sh

# 测试不加点 (你会感觉到明显的卡顿，或者看到 search 过程)
# 注意：nslookup 命令可能会自动显示它尝试的过程
nslookup google.com

# 测试加点 (响应通常会更快，且显示的 Server 是一样的)
nslookup google.com.
```

-----

### 总结

  * **修改位置：** 仅在业务代码的配置文件（如 DB 连接串、API 地址）中修改。
  * **CoreDNS/LocalDNS：** **完全不动。**

**下一步建议：**
既然明确了不需要改服务端，你现在就可以找一个不重要的测试环境服务，把它的外部数据库连接地址后面加个点，重启服务观察一下业务日志是否正常，确认没有兼容性问题。需要我提供关于“如何检查应用是否支持 FQDN”的排查思路吗？


如果你的集群使用了 **NodeLocal DNSCache**，需要同时配置它，因为 Pod 的 DNS 请求会先到 NodeLocal DNS，而不是直接到 CoreDNS。

## 架构说明

```
使用 NodeLocal DNS 的请求流程：

┌─────────┐      ┌──────────────────┐      ┌──────────┐      ┌─────────────┐
│   Pod   │ ───> │ NodeLocal DNS    │ ───> │ CoreDNS  │ ───> │ 阿里云 DNS   │
│         │      │ (169.254.20.10)  │      │          │      │             │
└─────────┘      └──────────────────┘      └──────────┘      └─────────────┘
```

## 解决方案

### 步骤 1：配置 NodeLocal DNSCache

编辑 NodeLocal DNS ConfigMap：

```bash
kubectl edit configmap node-local-dns -n kube-system
```

修改 Corefile 配置：

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: node-local-dns
  namespace: kube-system
data:
  Corefile: |
    cluster.local:53 {
        errors
        cache {
            success 9984 30
            denial 9984 5
        }
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__CLUSTER__DNS__ {
            force_tcp
        }
        prometheus :9253
    }
    
    in-addr.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__CLUSTER__DNS__ {
            force_tcp
        }
        prometheus :9253
    }
    
    ip6.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__CLUSTER__DNS__ {
            force_tcp
        }
        prometheus :9253
    }
    
    # 添加阿里云域名转发规则
    aliyuncs.com:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . 100.100.2.136 100.100.2.138
        prometheus :9253
    }
    
    .:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__UPSTREAM__SERVERS__ {
            force_tcp
        }
        prometheus :9253
    }
```

### 步骤 2：重启 NodeLocal DNS

```bash
# 查看 NodeLocal DNS Pods
kubectl get pods -n kube-system -l k8s-app=node-local-dns

# 删除 Pods 让其重建（DaemonSet 会自动重建）
kubectl delete pods -n kube-system -l k8s-app=node-local-dns

# 或者使用 rollout restart（如果支持）
kubectl rollout restart daemonset node-local-dns -n kube-system
```

### 步骤 3：同时配置 CoreDNS（作为后备）

```bash
kubectl edit configmap coredns -n kube-system
```

```yaml
data:
  Corefile: |
    .:53 {
        errors
        health {
            lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
            pods insecure
            fallthrough in-addr.arpa ip6.arpa
            ttl 30
        }
        prometheus :9153
        # 添加阿里云域名转发
        forward aliyuncs.com 100.100.2.136 100.100.2.138
        forward . /etc/resolv.conf {
            max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
```

重启 CoreDNS：

```bash
kubectl rollout restart deployment coredns -n kube-system
```

## 验证配置

### 1. 检查 NodeLocal DNS 配置

```bash
# 查看 ConfigMap
kubectl get configmap node-local-dns -n kube-system -o yaml

# 查看 Pod 日志
kubectl logs -n kube-system -l k8s-app=node-local-dns --tail=50
```

### 2. 在 Pod 中测试

```bash
# 进入测试 Pod
kubectl run test-dns --image=busybox:1.28 --rm -it -- sh

# 查看 DNS 配置
cat /etc/resolv.conf
# 应该显示：nameserver 169.254.20.10

# 测试解析
nslookup tars-datasets-shanghai.oss-cn-shanghai-internal.aliyuncs.com

# 测试带点的域名
nslookup tars-datasets-shanghai.oss-cn-shanghai-internal.aliyuncs.com.
```

### 3. 查看 DNS 查询日志

```bash
# 查看 NodeLocal DNS 日志
kubectl logs -n kube-system -l k8s-app=node-local-dns -f | grep aliyuncs

# 查看 CoreDNS 日志（作为对比）
kubectl logs -n kube-system -l k8s-app=kube-dns -f | grep aliyuncs
```

## 常见问题排查

### 问题 1：配置不生效

```bash
# 确认 ConfigMap 已更新
kubectl describe configmap node-local-dns -n kube-system

# 强制删除所有 NodeLocal DNS Pods
kubectl delete pods -n kube-system -l k8s-app=node-local-dns --force --grace-period=0

# 等待重建完成
kubectl get pods -n kube-system -l k8s-app=node-local-dns -w
```

### 问题 2：找不到 __PILLAR__ 变量

如果你的集群使用的是不同的部署方式，可能需要替换占位符：

```yaml
# 将 __PILLAR__CLUSTER__DNS__ 替换为实际的 CoreDNS 服务 IP
forward . 10.96.0.10  # 你的 CoreDNS Service ClusterIP

# 将 __PILLAR__UPSTREAM__SERVERS__ 替换为上游 DNS
forward . /etc/resolv.conf
# 或
forward . 8.8.8.8 8.8.4.4
```

查找实际值：

```bash
# 查找 CoreDNS Service IP
kubectl get svc kube-dns -n kube-system -o jsonpath='{.spec.clusterIP}'

# 查看节点的 DNS 配置
kubectl debug node/<node-name> -it --image=busybox:1.28
cat /etc/resolv.conf
```

## 最佳配置示例（完整版）

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: node-local-dns
  namespace: kube-system
data:
  Corefile: |
    cluster.local:53 {
        errors
        cache {
            success 9984 30
            denial 9984 5
        }
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 10.96.0.10 {
            force_tcp
        }
        prometheus :9253
        health 169.254.20.10:8080
    }
    
    in-addr.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 10.96.0.10 {
            force_tcp
        }
        prometheus :9253
    }
    
    ip6.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 10.96.0.10 {
            force_tcp
        }
        prometheus :9253
    }
    
    # 阿里云内网域名
    aliyuncs.com:53 {
        errors
        cache {
            success 9984 30
            denial 9984 5
        }
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 100.100.2.136 100.100.2.138
        prometheus :9253
    }
    
    .:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . /etc/resolv.conf
        prometheus :9253
    }
```

## 预期效果

配置完成后：

1. ✅ Pod 的 DNS 请求首先到达 NodeLocal DNS
2. ✅ 阿里云域名直接被 NodeLocal DNS 转发到阿里云 DNS
3. ✅ 缓存在本地，减少延迟
4. ✅ 不再看到 NXDOMAIN 日志（或只看到一次）
5. ✅ 域名解析成功

配置后，那些 `NXDOMAIN` 日志会显著减少或消失！


如果你的集群使用了 **NodeLocal DNSCache**，需要同时配置它，因为 Pod 的 DNS 请求会先到 NodeLocal DNS，而不是直接到 CoreDNS。

## 架构说明

```
使用 NodeLocal DNS 的请求流程：

┌─────────┐      ┌──────────────────┐      ┌──────────┐      ┌─────────────┐
│   Pod   │ ───> │ NodeLocal DNS    │ ───> │ CoreDNS  │ ───> │ 阿里云 DNS   │
│         │      │ (169.254.20.10)  │      │          │      │             │
└─────────┘      └──────────────────┘      └──────────┘      └─────────────┘
```

## 解决方案

### 步骤 1：配置 NodeLocal DNSCache

编辑 NodeLocal DNS ConfigMap：

```bash
kubectl edit configmap node-local-dns -n kube-system
```

修改 Corefile 配置：

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: node-local-dns
  namespace: kube-system
data:
  Corefile: |
    cluster.local:53 {
        errors
        cache {
            success 9984 30
            denial 9984 5
        }
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__CLUSTER__DNS__ {
            force_tcp
        }
        prometheus :9253
    }
    
    in-addr.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__CLUSTER__DNS__ {
            force_tcp
        }
        prometheus :9253
    }
    
    ip6.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__CLUSTER__DNS__ {
            force_tcp
        }
        prometheus :9253
    }
    
    # 添加阿里云域名转发规则
    aliyuncs.com:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . 100.100.2.136 100.100.2.138
        prometheus :9253
    }
    
    .:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10
        forward . __PILLAR__UPSTREAM__SERVERS__ {
            force_tcp
        }
        prometheus :9253
    }
```

### 步骤 2：重启 NodeLocal DNS

```bash
# 查看 NodeLocal DNS Pods
kubectl get pods -n kube-system -l k8s-app=node-local-dns

# 删除 Pods 让其重建（DaemonSet 会自动重建）
kubectl delete pods -n kube-system -l k8s-app=node-local-dns

# 或者使用 rollout restart（如果支持）
kubectl rollout restart daemonset node-local-dns -n kube-system
```

### 步骤 3：同时配置 CoreDNS（作为后备）

```bash
kubectl edit configmap coredns -n kube-system
```

```yaml
data:
  Corefile: |
    .:53 {
        errors
        health {
            lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
            pods insecure
            fallthrough in-addr.arpa ip6.arpa
            ttl 30
        }
        prometheus :9153
        # 添加阿里云域名转发
        forward aliyuncs.com 100.100.2.136 100.100.2.138
        forward . /etc/resolv.conf {
            max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
```

重启 CoreDNS：

```bash
kubectl rollout restart deployment coredns -n kube-system
```

## 验证配置

### 1. 检查 NodeLocal DNS 配置

```bash
# 查看 ConfigMap
kubectl get configmap node-local-dns -n kube-system -o yaml

# 查看 Pod 日志
kubectl logs -n kube-system -l k8s-app=node-local-dns --tail=50
```

### 2. 在 Pod 中测试

```bash
# 进入测试 Pod
kubectl run test-dns --image=busybox:1.28 --rm -it -- sh

# 查看 DNS 配置
cat /etc/resolv.conf
# 应该显示：nameserver 169.254.20.10

# 测试解析
nslookup tars-datasets-shanghai.oss-cn-shanghai-internal.aliyuncs.com

# 测试带点的域名
nslookup tars-datasets-shanghai.oss-cn-shanghai-internal.aliyuncs.com.
```

### 3. 查看 DNS 查询日志

```bash
# 查看 NodeLocal DNS 日志
kubectl logs -n kube-system -l k8s-app=node-local-dns -f | grep aliyuncs

# 查看 CoreDNS 日志（作为对比）
kubectl logs -n kube-system -l k8s-app=kube-dns -f | grep aliyuncs
```

## 常见问题排查

### 问题 1：配置不生效

```bash
# 确认 ConfigMap 已更新
kubectl describe configmap node-local-dns -n kube-system

# 强制删除所有 NodeLocal DNS Pods
kubectl delete pods -n kube-system -l k8s-app=node-local-dns --force --grace-period=0

# 等待重建完成
kubectl get pods -n kube-system -l k8s-app=node-local-dns -w
```

### 问题 2：找不到 __PILLAR__ 变量

如果你的集群使用的是不同的部署方式，可能需要替换占位符：

```yaml
# 将 __PILLAR__CLUSTER__DNS__ 替换为实际的 CoreDNS 服务 IP
forward . 10.96.0.10  # 你的 CoreDNS Service ClusterIP

# 将 __PILLAR__UPSTREAM__SERVERS__ 替换为上游 DNS
forward . /etc/resolv.conf
# 或
forward . 8.8.8.8 8.8.4.4
```

查找实际值：

```bash
# 查找 CoreDNS Service IP
kubectl get svc kube-dns -n kube-system -o jsonpath='{.spec.clusterIP}'

# 查看节点的 DNS 配置
kubectl debug node/<node-name> -it --image=busybox:1.28
cat /etc/resolv.conf
```

## 最佳配置示例（完整版）

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: node-local-dns
  namespace: kube-system
data:
  Corefile: |
    cluster.local:53 {
        errors
        cache {
            success 9984 30
            denial 9984 5
        }
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 10.96.0.10 {
            force_tcp
        }
        prometheus :9253
        health 169.254.20.10:8080
    }
    
    in-addr.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 10.96.0.10 {
            force_tcp
        }
        prometheus :9253
    }
    
    ip6.arpa:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 10.96.0.10 {
            force_tcp
        }
        prometheus :9253
    }
    
    # 阿里云内网域名
    aliyuncs.com:53 {
        errors
        cache {
            success 9984 30
            denial 9984 5
        }
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . 100.100.2.136 100.100.2.138
        prometheus :9253
    }
    
    .:53 {
        errors
        cache 30
        reload
        loop
        bind 169.254.20.10 10.96.0.10
        forward . /etc/resolv.conf
        prometheus :9253
    }
```

## 预期效果

配置完成后：

1. ✅ Pod 的 DNS 请求首先到达 NodeLocal DNS
2. ✅ 阿里云域名直接被 NodeLocal DNS 转发到阿里云 DNS
3. ✅ 缓存在本地，减少延迟
4. ✅ 不再看到 NXDOMAIN 日志（或只看到一次）
5. ✅ 域名解析成功

配置后，那些 `NXDOMAIN` 日志会显著减少或消失！
