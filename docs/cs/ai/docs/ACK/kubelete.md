# ACK

保障系统资源不被业务负载占用解决方案

分配预留

硬性保证

## 分配预留

在阿里云 ACK (Container Service for Kubernetes) 集群中，为了防止业务容器过度占用资源导致系统组件（如 `kubelet`、`containerd`、`csi-plugin` 等）崩溃或无法上报状态，通常采用 **Kubelet 资源预留 (Resource Reservation)** 的机制。

你可以通过以下几种方式来实现节点 CPU 的保留：

---

### 1. 扩容/创建节点池时设置（推荐）

在 ACK 控制台创建或修改节点池时，可以通过**自定义参数**来配置 `kube-reserved`。

* **kube-reserved**：为 Kubernetes 系统组件（kubelet, container runtime 等）预留的资源。
* **system-reserved**：为 OS 系统进程（ssh, systemd 等）预留的资源。

**操作步骤：**

1. 登录容器服务管理控制台。
2. 进入 **节点池**，点击 **创建节点池** 或编辑现有节点池的 **配置参数**。
3. 在 `Kubelet` 参数配置中，添加或修改：
* `kube-reserved`: `cpu=1,memory=2Gi` (根据实例规格调整)
* `system-reserved`: `cpu=500m,memory=1Gi`


4. 确认后，新加入该节点池的节点会自动生效。

---

### 2. 修改现有节点的 Kubelet 配置

如果你需要对存量节点进行调整，可以通过修改节点上的 `Kubelet` 配置文件：

1. **登录节点**：使用 SSH 登录目标节点。
2. **编辑配置**：修改 `/etc/kubernetes/kubelet.conf`（ACK 默认路径可能略有不同，通常在 `/var/lib/kubelet/config.yaml`）。
3. **添加参数**：
```yaml
kubeReserved:
  cpu: "1000m"
  memory: "2Gi"
systemReserved:
  cpu: "500m"
  memory: "1Gi"
evictionHard:
  memory.available: "500Mi"

```

4. **重启 Kubelet**：
```bash
systemctl restart kubelet

```

### 3. 最佳实践建议

* **比例预留**：对于小规格实例（如 4核），建议预留 0.5-1 核；对于大规格实例（如 32核+），建议预留 2-4 核。
* **结合驱逐策略**：除了 CPU，**内存预留**更为关键。建议设置 `evictionHard`（硬驱逐阈值），在内存极度紧缺时保护系统不进入 OOM 状态。

我明白你的意思了。你遇到的不是“调度时”没地方放 Pod，而是**运行时（Runtime）**由于 Pod 设置了较高的 `limit` 或者压根没设 `limit`，导致大量业务 Pod 疯狂抢占 CPU，甚至把 `kubelet` 等系统进程的 CPU 时间片也给抢走了，导致系统组件“假死”，无法及时上报心跳。

在 ACK 这种基于 Linux Cgroups 的环境下，要解决“运行时溢出”对系统组件的影响，核心在于**强制隔离系统进程与业务进程的 Cgroup 控制组**。

以下是针对“运行时溢出”的深度配置方案：

---

## 硬性保证

### 1. 开启 Kubelet 的 Cgroup 强制预留 (Enforcement)

默认情况下，Kubernetes 只是在账面上减掉预留资源。要防止 Pod 运行时侵占系统 CPU，必须开启 `enforce-node-allocatable` 参数。

在 ACK 节点池的 Kubelet 自定义参数中确保以下配置：

* **参数名**：`--enforce-node-allocatable`
* **推荐值**：`pods,system-reserved,kube-reserved`

**原理**：

* 设置后，所有 Pod 都会被限制在一个总的 Cgroup 子树下（`kubepods.slice`）。
* 该子树的总 CPU 配额等于 `NodeAllocatable`。
* **即便所有 Pod 的 limit 加起来超标，它们也只能在分配给 Pod 的那部分 CPU 范围内竞争，绝对无法碰到系统组件预留的那部分 CPU。**

---

### 2. 启用 CPU Manager（针对关键系统组件隔离）

如果你的某些系统组件非常敏感，可以使用 Kubelet 的 **静态 CPU 管理策略 (Static Policy)**。

* **配置参数**：`--cpu-manager-policy=static`
* 将 `kube-system` 下的关键组件（如 `csi-plugin` 或日志组件）设置为 **Guaranteed QoS**（即 `request` 等于 `limit` 且为整数核）。
* 这样 Kubelet 会给这些组件分配**独占 CPU 核心**，业务 Pod 运行时无法使用这些物理核，从物理层面上实现隔离。

---

## 3. 使用 ACK 的“弹性资源限制” (ack-koordinator)

如果你开启了 ACK 的“动态资源超卖”功能，单纯靠 Cgroup 限制可能不够灵活。建议部署 `ack-koordinator`（原 `ack-slo-manager`）：

* **CPU Suppress (CPU 压制)**：它可以实时监控节点 CPU 总利用率。当系统负载超过安全水位（例如 90%）时，它会自动通过 Cgroup **动态调低** 低优先级（BestEffort/Burstable）Pod 的 `cpu.cfs_quota_us`，强制给系统组件让出路来。
* **配置方式**：
在 `kube-system` 命名空间下配置 `ack-slo-config`：
```yaml
resource-threshold-config: |
  {
    "clusterStrategy": {
      "enable": true,
      "cpuSuppressThresholdPercent": 85
    }
  }

```

---

## 4. 关键：修正“超卖率”引发的 CPU Throttling

如果 Pod 资源溢出导致了节点卡顿，通常是因为宿主机开启了 CPU 限制（CFS Quota）。

* **建议方案**：如果业务对延迟敏感，可以尝试开启 **ACK 容器 CPU QoS**。它利用了内核的 `Bfield` 或 `Group Identity` 技术，确保系统组件和高优先级 Pod 在内核调度器中拥有更高的**调度权重 (Priority)**，即使 CPU 跑满，内核也会优先处理系统进程的指令。

---

### 总结建议的操作顺序：

1. **第一步**：检查并设置 `enforce-node-allocatable=pods,kube-reserved`，这是最硬的物理防线。
2. **第二步**：提高 `kube-reserved` 的 CPU 预留值（建议对于 16核以上机器至少预留 1核-2核）。
3. **第三步**：如果节点仍然容易因超卖假死，部署 `ack-koordinator` 设置 `cpuSuppressThresholdPercent` 安全水位线。

**你需要我帮你写一个针对特定节点池修改 `enforce-node-allocatable` 的具体操作步骤吗？**