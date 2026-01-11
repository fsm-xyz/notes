# Airflow云原生架构演进历程

**云原生架构演进历程**。从追求“极致的技术优雅”（异步回调、自定义调度器、资源感知）出发，在经历了各种状态同步、清理失效、生命周期管理的坑之后，最终回归到“稳健与灵活性并重”的生产实践。

原生架构是Celery+ KubernetesPodOperator

## 🛠️ Airflow 3.0 & Kubernetes 调度选型演进之路

### 第一阶段：追求极致资源效率 (Deferrable + Volcano)

* **初衷**：解决 `KubernetesExecutor` 下 Worker Pod 随任务 1:1 创建导致的资源浪费和 Pending Pod 堆积问题。
* **尝试**：利用 **Deferrable Operators + Triggers** 实现异步等待，结合 **Volcano** 调度器实现 AI/大负载任务的排队与感知。
* **痛点**：
* **状态管理复杂**：自定义 `VcJob` 的生命周期与 Airflow 状态同步极其困难（重试、删除、UI 显示）。
* **清理残留**：异步回调过程中 Pod 经常删除不干净，导致“僵尸任务”占用资源。
* **架构过重**：为了节省 Worker 资源，却引入了极大的运维复杂度。


### 第二阶段：调度器博弈 (Volcano vs. Koordinator)

* **痛点**：Volcano 开源版文档滞后，负载感知功能闭源，社区互动体感一般。
* **选型**：转向阿里云开源的 **Koordinator**。
* **优势**：组件模块化（灵活组装）、文档齐全、社区响应快，且具备成熟的**负载感知调度**和**资源超卖（Colocation）**能力。

### 第三阶段：返璞归真，回归主流 (CeleryKubernetesExecutor)

* **反思**：Airflow 的 `Trigger` 模式在 Kubernetes 环境下虽优雅但尚未完全成熟。
* **决策**：回归 **CeleryKubernetesExecutor**，平衡了常驻 Worker 处理轻量任务（高频、低延迟）与 Pod 动态扩展处理重量任务（高隔离、高资源）的需求。

---

## 🚀 最终架构：Airflow 3.0 在 K8s 上的最优实践

你总结的这套方案实际上是目前 Kubernetes 上运行 Airflow 的“天花板”配置：

### 1. 核心架构选型

* **执行器**：`CeleryKubernetesExecutor` (CKE)。
* **资源调度**：引入 **Koordinator** 开启负载感知调度，解决 Pod 堆积在高峰节点的难题。
* **资源利用**：利用 **ack-kube-queue** (或类似 ElasticQuota) 实现多租户/多任务排队。

### 2. 任务执行模式 (按推荐度排序)

* **方案 A (最推荐)**：`KubernetesExecutor` + **镜像融合/Sidecar**。
* *逻辑*：Worker 与 Task 容器同生共死，任务结束 Pod 即销毁。
* *优点*：彻底解决常驻 Worker 数量问题，实现“有多大锅下多大米”。


* **方案 B (性能平衡)**：`CeleryExecutor` + `KubernetesPodOperator` (KPO)。
* *逻辑*：由常驻 Worker 派发独立的 Pod 执行。
* *优点*：隔离性好，适合复杂环境，但 Pod 启动稍慢。



### 3. 资源治理策略

* **负载感知**：通过 Koordinator 实时监控节点水位，避免热点节点崩溃。
* **超卖与 Quota**：使用 `ElasticQuotaTree` 解决多命名空间资源分配，通过权重控制队列优先级。

---

### 💡 感悟精炼：

> **“理想是异步回调 + 资源感知的完美调度，现实是状态一致性与运维成本的拉锯战。”**
> 在 Airflow 3.0 时代，最稳健的路径是：**成熟执行器 (CKE) + 云原生精细化调度 (Koordinator/Kube-Queue) + 任务执行容器化 (Mirror/KPO)**。
