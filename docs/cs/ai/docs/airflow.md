# Airflow

## 配置

### 并发

还有个共享的pool可以限制task的数量, slots

```sh
全局最大并发 Task 数	parallelism（核心）     64
单个 DAG 最大并发 Task 数	dag_concurrency	32
单 DAG 最大活跃运行数	max_active_runs_per_dag	16
单 DAG 运行实例的最大活跃 Task 数	max_active_tasks_per_dagrun（3.x 新增优化）
Worker 级并发限制	worker_concurrency      32
```

### 超时时间

```sh

# 启动的超时时间
startup_timeout_seconds=720000, 
# 执行的超时时间
execution_timeout=timedelta(minutes=30),
```

### XCOM

在CeleryExecutor模式下，KubernetesPodOperator传递数据需要XCOM,但是XCOM需要sidecar模式，默认镜像国内无法访问，需要指定自定义镜像

AI总结的环境变量和配置都不对，2种方式页面上设置Connection或者helm下设置这个

```sh
airflowLocalSettings: |-
  {{- if semverCompare ">=2.2.0 <3.0.0" .Values.airflowVersion }}
  {{- if not (or .Values.webserverSecretKey .Values.webserverSecretKeySecretName) }}
  from airflow.www.utils import UIAlert

  DASHBOARD_UIALERTS = [
    UIAlert(
      'Usage of a dynamic webserver secret key detected. We recommend a static webserver secret key instead.'
      ' See the <a href='
      '"https://airflow.apache.org/docs/helm-chart/stable/production-guide.html#webserver-secret-key" '
      'target="_blank" rel="noopener noreferrer">'
      'Helm Chart Production Guide</a> for more details.',
      category="warning",
      roles=["Admin"],
      html=True,
    )
  ]
  {{- end }}
  {{- end }}

      # 这块儿
  from airflow.providers.cncf.kubernetes.utils.xcom_sidecar import PodDefaults
  PodDefaults.SIDECAR_CONTAINER.image_pull_policy = "IfNotPresent"
  PodDefaults.SIDECAR_CONTAINER.image = "alpine"
```

### keda

通过这个进行动态扩容worker,不像原始的HPA只能实现CPU和内存的扩容
helm/airflow/
scheduler把任务进入队列，然后keda-metrices-api-server触发轮训, 调用keda-oprator获取资源，发现需要进行扩容，创建HPA,然后设置数量

keda-operator它根据 ScaledObject 里定义的 pollingInterval（默认 30s），定期去连你的 Postgres/MySQL 数据库。

如果查询结果 > 0：KEDA 认为“有活干了”，直接修改 Worker Deployment 的副本数，将 replicas 从 0 改为 1。

一旦副本数到达 1，后续的扩容工作转交给 Kubernetes 原生的 HPA 负责，KEDA 退居幕后充当“情报员”。

HPA 定期（默认 15s）向 KEDA Metrics Server 发起 HTTP/gRPC 请求。

做什么： 接到 HPA 请求后，它立马连接 Postgres/MySQL 数据库，执行 SQL 查询，把查到的数字返回给 HPA。

Worker 启动:

Kubernetes 调度并启动新的 Worker Pod。

Worker 启动 airflow celery worker 进程。

抢任务:
helm/airflow/
Worker 连接 Redis/RabbitMQ（Broker），抢走 queued 状态的任务。

Worker 更新数据库，将任务状态从 queued 改为 running。



结合你之前遇到的报错（HPA 获取指标超时）和 Airflow 的架构，KEDA 的扩缩容流程其实是一个 **"监视 -\> 汇报 -\> 决策 -\> 执行"** 的闭环。

理解这个流程能帮你瞬间明白为什么之前会报 `timeout` 错误，以及为什么 SQL 写错了就不动。

以下是 KEDA 在 Airflow 场景下的完整工作流：

-----

### 核心流程图解

```mermaid
graph TD
    DB[(Airflow Database)] -- 1. Polling (SQL) --> KEDA_OP[KEDA Operator]
    DB -- 4. Query Metric --> KEDA_METRIC[KEDA Metrics Server]
    
    KEDA_OP -- 2. Activate (0->1) --> WORKER[Airflow Worker Deployment]
    KEDA_OP -- 3. Create/Manage --> HPA[Kubernetes HPA]
    
    HPA -- 4. Ask: "How many tasks?" --> KEDA_METRIC
    KEDA_METRIC -- 5. Return: "100 tasks" --> HPA
    
    HPA -- 6. Scale (1->N) --> WORKER
    
    WORKER -- 7. Consume Task --> BROKER[Redis/RabbitMQ]
    WORKER -- 8. Update State --> DB
```

-----

### 详细步骤拆解

#### 第一阶段：激活 (Activation) —— 从 0 到 1

这是 KEDA 独有的能力（HPA 无法缩容到 0）。

1.  **轮询 (Polling):**
      * **谁在做：** `keda-operator`
      * **做什么：** 它根据 `ScaledObject` 里定义的 `pollingInterval`（默认 30s），定期去连你的 Postgres/MySQL 数据库。
      * **查什么：** 执行那条 SQL 语句 (`SELECT count(*) FROM task_instance WHERE state='queued'...`)。
2.  **激活决策:**
      * 如果查询结果 **\> 0**：KEDA 认为“有活干了”，直接修改 Worker Deployment 的副本数，将 `replicas` 从 0 改为 1。
      * 同时，KEDA 会创建一个 Kubernetes HPA 资源（如果还没创建的话）。

#### 第二阶段：水平扩容 (Scaling) —— 从 1 到 N

一旦副本数到达 1，后续的扩容工作转交给 Kubernetes 原生的 **HPA** 负责，KEDA 退居幕后充当“情报员”。

3.  **HPA 索要情报:**

      * **谁在做：** `HorizontalPodAutoscaler` (HPA) Controller。
      * **做什么：** HPA 定期（默认 15s）向 **KEDA Metrics Server** 发起 HTTP/gRPC 请求。
      * **问什么：** “现在的 External Metric (任务数) 是多少？”
      * *💥你之前的 `FailedGetExternalMetric timeout` 报错就发生在这里！HPA 问了，但 KEDA Metrics Server 没回答。*

4.  **情报员查询:**

      * **谁在做：** `keda-metrics-apiserver`
      * **做什么：** 接到 HPA 请求后，它立马连接 Postgres/MySQL 数据库，执行 SQL 查询，把查到的数字返回给 HPA。

5.  **计算与扩容:**

      * **谁在做：** HPA Controller。
      * **算法：** `期望副本数 = ceil(当前排队任务数 / 每个Worker的目标任务数)`。
      * **执行：** HPA 修改 Worker Deployment 的 `replicas`（例如从 1 改成 5）。

#### 第三阶段：消费与执行

6.  **Worker 启动:**
      * Kubernetes 调度并启动新的 Worker Pod。
      * Worker 启动 `airflow celery worker` 进程。
7.  **抢任务:**
      * Worker 连接 Redis/RabbitMQ（Broker），抢走 `queued` 状态的任务。
      * Worker 更新数据库，将任务状态从 `queued` 改为 `running`。

#### 第四阶段：缩容 (Scale Down) —— 从 N 到 0

8.  **指标下降:**
      * 随着任务变为 `running` 或 `success`，数据库里 `state='queued'` 的数量减少。
      * HPA 下一次询问时，得到的数字变小，于是计算出更少的副本数，开始缩减 Worker Pod。
9.  **冷却与归零:**
      * 当任务数归零后，HPA 会把副本数缩减到 1（由 `minReplicaCount` 控制）。
      * **KEDA Operator** 再次接管：如果它发现查询结果持续为 0 超过了 `cooldownPeriod`（默认 5 分钟），它会将副本数从 1 强制改为 0。

-----

### 结合你的问题分析

根据这个流程，我们可以清晰地定位你之前的三个问题：

1.  **问题：`task_instance` 状态全是 NULL，不扩容**

      * **环节：** 步骤 1 (Polling)。
      * **原因：** 数据库里没有 `queued` 的记录，SQL 查出来是 0。KEDA 认为没活干，所以一直在 0 副本待机。

2.  **问题：`dag_run` 是 queued，但 `task` 是 NULL**

      * **环节：** 步骤 0 (Setup)。
      * **原因：** 这是上游 Airflow Scheduler 的问题，还没来得及把 DAG 生成 Task 写入数据库。KEDA 此时是无辜的，它忠实地反映了 Task 表为空的事实。

3.  **问题：`FailedGetExternalMetric timeout`**

      * **环节：** 步骤 3 & 4 (HPA 问情报)。
      * **原因：** HPA 试图联系 `keda-metrics-apiserver`，但在网络层（gRPC连接）或者数据库层（SQL查询）卡死了。导致 HPA 无法做出扩容到 N 的决定。

### 总结

要让这个流程跑通，必须满足三个条件：

1.  **Airflow Scheduler 正常**：把任务塞进数据库（状态变为 `queued`）。
2.  **KEDA 组件通信正常**：HPA 能问到 KEDA，KEDA 能连上 Operator。
3.  **SQL 查询正确**：KEDA 查的 SQL 必须能匹配到那批 `queued` 的任务（队列名匹配、状态匹配）。