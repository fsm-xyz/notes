# k8s

API Server

etcd	分布式键值存储，保存集群状态	数据一致性算法（Raft）、备份恢复策略

Controller Manager  维护资源状态（如Deployment副本数）

Scheduler   将 Pod 绑定到合适 Node	调度算法流程（预选+优选+绑定）

kubelet     节点代理，管理 Pod 生命周期

kube-proxy  实现 Service 负载均衡（iptables/IPVS）

POD     pause 容器实现网络共享

Pending → Running → Succeeded/Failed

关键问题：

Q：如何保证 Pod 内容器启动顺序？
答：使用 Init Container（如先启动数据库迁移容器，再启应用容器）


Deployment

StatefulSet

DaemonSet

CronJob


ClusterIP

NodePort

LoadBalancer

Ingress 控制器

PV

PVC

StorageClass	动态创建 PV 的模板（如AWS EBS）

ConfigMap

Secret

ENV


gRPC无法使用svc进行负载均衡

NetworkPolicy


RBAC


服务发现: CoreDNS  + 环境变量

Request/Limit. HPA. VPA, Descheduler


## CNI

Flannel， Calico

### Flannel

+ VXLAN 模式， 以太帧将进行封装
+ host-gw 模式（Host Gateway）完全基于路由表转发，无隧道封装
+ UDP 模式

### Calico

纯三层容器网络，BGP 协议实现跨主机容器通信，并基于 iptables 实现网络策略

IPIP和BGP模式

IPIP 模式通过隧道解决跨网段问题
BGP 模式要求底层网络支持路由可达（如交换机配置）

+ Felix(管理容器网络接口，配置路由规则,通过 iptables 设置 ACL 规则，实现网络策略隔离，监控并上报节点网络状态至 etcd)
+ BGP Client (BIRD)  监听内核路由变化，通过 BGP 协议将本节点容器路由信息，小规模集群：节点间直接互联，大规模集群：通过 BGP Route Reflector (RR) 集中分发路由
+ etcd 存储集群网络元数据（如 IP 分配、路由规则、策略配置），确保状态一致性


路由规模问题

每个容器

优化：聚合路由（默认按 /26 块分配），或使用 RR 集中管理37。

路由黑洞

当节点借用其他节点的 IP 块时，可能因黑洞路由（blackhole /28 proto bird）导致流量丢弃。

规避：确保 IP 池充足，或限制每节点 Pod 数量（kubelet --max-pods）7。

性能瓶颈

iptables 规则过多时影响转发效率。

替代方案：启用 eBPF 数据平面（绕过 iptables，提升性能）


eBPF

### Clium

Cilium 是基于 eBPF（扩展伯克利包过滤器） 的云原生网络、安全与可观测性方案

Overlay 隧道模式

Native 路由模式（BGP/Underlay）

多集群方案（ClusterMesh）

### Terway(阿里云)

VPC模式
ENI模式
ENI多IP
ENI-Trunking

### 双网卡(虚拟网卡)

### BGP 路由表

### 虚拟交换机


pending

Runing

succeed

Failed

UNKNOWN


kubernetes是目前最热门的技术之一，也是很多公司技术栈中不可或缺的一部分。因此，在面试中，kubernetes相关的知识也成了考察候选人能力的重要方面。

以下是为kubernetes初学者精心整理的一些常见的面试题以及对应的答案。

---

### **基础概念**

#### 1. 什么是Kubernetes？

Kubernetes（简称K8s）是一个开源的容器编排平台，用于自动化应用程序的部署、扩展和管理。它可以将容器化的应用程序部署到集群中，并提供了一套机制来确保应用程序的高可用性、可扩展性和弹性。

---

### **核心组件**

#### 2. Kubernetes的核心组件有哪些？

Kubernetes的核心组件主要分为两类：**控制平面（Control Plane）**组件和**节点（Node）**组件。

**控制平面组件：**
- **kube-apiserver**：集群的API服务器，提供了所有集群操作的入口。
- **etcd**：一个高可用的键值存储系统，用于保存集群的所有状态和配置数据。
- **kube-scheduler**：负责将Pod调度到合适的节点上。
- **kube-controller-manager**：运行各种控制器，如节点控制器、副本控制器、端点控制器等，用于维护集群的期望状态。

**节点组件：**
- **kubelet**：运行在每个节点上的代理，负责管理Pod和容器的生命周期。
- **kube-proxy**：为集群中的服务提供网络代理和负载均衡功能。
- **容器运行时**：负责运行容器，如Docker、containerd、CRI-O等。

---

### **核心对象**

#### 3. 什么是Pod？Pod有什么特点？

**Pod**是Kubernetes中最小的、可部署的计算单元，它可以包含一个或多个紧密相关的容器。Pod中的容器共享网络命名空间、存储卷等资源。

**特点：**
- **原子性**：Pod是调度的最小单位。
- **共享**：同一个Pod中的容器共享网络、存储等资源。
- **短暂性**：Pod是短暂的，当它被杀死后不会自动重启，需要通过控制器（如Deployment）来管理其生命周期。

---

### **部署和管理**

#### 4. Deployment和Pod有什么区别？

- **Pod**是最小的部署单元，它本身不具备自我修复能力。
- **Deployment**是一个更高层次的API对象，用于管理Pod和ReplicaSet。它可以确保指定数量的Pod副本始终在运行，并提供了**滚动更新**和**回滚**等功能，让应用程序的部署和升级变得更加平滑和安全。

#### 5. 什么是Service？它的作用是什么？

**Service**是Kubernetes中用于暴露一组Pod的抽象。它为Pod提供一个稳定的网络地址（IP地址和端口），即使Pod被重新创建或迁移，Service的IP地址也不会改变。

**作用：**
- **服务发现**：Service提供了稳定的网络入口，让其他Pod或外部应用可以方便地访问这些Pod。
- **负载均衡**：Service可以将流量分发到其背后的多个Pod上。

#### 6. Service有哪些类型？

常见的Service类型有四种：
- **ClusterIP**：默认类型，为Service分配一个集群内部IP，只能在集群内部访问。
- **NodePort**：在每个节点上暴露一个静态端口，可以通过`NodeIP:NodePort`从集群外部访问。
- **LoadBalancer**：在云服务商上创建外部负载均衡器，用于将外部流量路由到Service。
- **ExternalName**：通过返回CNAME记录，将Service映射到外部域名。

---

### **网络和存储**

#### 7. Kubernetes网络模型是怎样的？

Kubernetes的网络模型遵循以下原则：
- 每个Pod都有一个独立的IP地址。
- 同一节点上的Pod可以直接通过IP地址通信。
- 不同节点上的Pod可以直接通过IP地址通信，无需进行NAT转换。
- 节点上的代理（kube-proxy）可以访问所有Pod。

要实现这个模型，通常需要使用**网络插件（CNI）**，如Flannel、Calico、Cilium等。

#### 8. 什么是PV和PVC？

- **PV（PersistentVolume）**：由管理员或存储提供者提供的、集群中的一块存储资源。它独立于Pod的生命周期。
- **PVC（PersistentVolumeClaim）**：用户对存储资源的请求。当一个Pod需要存储时，它会通过PVC来申请。

PV和PVC的设计实现了存储的**解耦**，让应用程序开发者无需关心底层存储的具体实现。

---

### **安全和配置**

#### 9. ConfigMap和Secret有什么区别？

- **ConfigMap**：用于存储非敏感的配置数据，如配置文件、环境变量等。
- **Secret**：用于存储敏感信息，如密码、API密钥、令牌等。Secret中的数据会被加密存储在etcd中。

它们都可以在Pod中以**环境变量**或**文件**的形式被挂载使用。

#### 10. 如何管理Pod的资源？

可以通过在Pod的YAML文件中设置`resources`字段来管理Pod的资源。

- **`requests`**：Pod所需的**最小**资源量。调度器会根据这个值来决定将Pod调度到哪个节点上。
- **`limits`**：Pod可以使用的**最大**资源量。如果Pod的使用量超过这个值，可能会被**OOM Killer**杀死或受到限制。

---
这是一个很好的问题。**ConfigMap** 的更新机制是 Kubernetes 中的一个重要概念。

简单来说，ConfigMap 默认情况下不会动态更新到容器内部，它的行为取决于你使用它的方式。

---

### ConfigMap 的更新方式

ConfigMap 有两种主要的使用方式：

1.  **作为环境变量**
2.  **作为文件卷**

#### 1. 作为环境变量

当你在 Pod 中以环境变量的形式引用 ConfigMap 的键值时，**如果 ConfigMap 更新了，环境变量的值不会自动更新**。你需要重启 Pod 才能让新的值生效。

这是因为 Pod 的环境变量是在创建时从 ConfigMap 中提取并设置的，之后就不会再检查 ConfigMap 的变化。

#### 2. 作为文件卷

当你在 Pod 中以文件卷的形式挂载 ConfigMap 时，情况则不同：

-   kubelet 会监控 ConfigMap 的变化。
-   当 ConfigMap 更新后，kubelet 会自动更新 Pod 中挂载的文件。
-   更新通常会有延迟，**大约在 10 秒到 1 分钟之间**。

**一个重要的注意事项：** 应用程序本身需要有能力监听文件变化。如果你的应用程序在启动时只读取一次配置文件，那么即使文件更新了，它也不会使用新配置。

---

### 最佳实践

在大多数需要动态更新配置的场景中，将 ConfigMap 作为**文件卷**挂载是更推荐的做法。

如果你希望应用程序能够对配置变化做出响应，你的应用程序代码需要实现以下逻辑：

-   **监听文件系统事件**：使用 `inotify`（Linux）或其他类似机制来监听挂载的配置文件的变化。
-   **重新加载配置**：当检测到文件变化时，应用程序应该重新加载配置并应用更改，而无需重启。

如果你的应用程序无法实现这一点，那么唯一的办法就是重启 Pod 来应用新的配置。为此，你可以使用像 **Deployment** 这样的控制器来管理 Pod，更新 ConfigMap 后，可以通过 `kubectl rollout restart` 命令来触发滚动重启。