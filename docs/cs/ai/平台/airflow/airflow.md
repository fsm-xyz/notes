# Airflow

主要基于Airflow3.0的配置以及架构实践

## 架构

Scheduler → Worker Process → Task Execution

```markdown
:::info

更现代，但是不成熟，和UI上的操作结合得不好，需要自己深度管理状态
:::info
```
Scheduler → Worker Process → Deferer  → Task Execution → Trigger

### 组件

+ WebServer(Web UI)
+ ApiServer(API接口)
+ DagProcessor  Airflow 3.0 中负责解析和处理 DAG 文件的独立组件，它将 DAG 解析工作从 Scheduler 中分离出来
+ Scheduler 控制任务的调度
+ Worker 进行监控或者执行任务
+ Flower worker状态的监控
+ Trigger task就绪回调触发
+ Pg 存储任务的元数据
+ Redis存储调度的任务队列


Airflow后面增加了基于事件回调的模式，避免Worker Process常驻，浪费资源，尤其是在kubernetes环境中，但是Trigger + Deferered模式在实践中没有多少人用，自己写了一个基于BashOperator的，实现Volcano + Trigger + Deferered这样的一个在kubernetes中运行，但是跟Airflow UI结合的不好，所以还是回归到原生的UI + 当前的成熟方案CeleryExecutor，KubernetesExecutor，CeleryKubernetesExecutor)，理想很丰满，现实很骨干，等官方UI和Trigger这一套成熟稳定的时候，再使用，现在还回归主流吧

### Executor

#### LocalExecutor

开发环境，任务量小，快速验证用的

#### CeleryExecutor

非云原生时代的架构，依赖外部的消息队列，进行任务分发

Scheduler → Worker Process → Task Execution

BashOperator： 直接在worker内执行

KubernetesOperator: 直接启动taskpod, worker里面监听

#### KubernetesExecutor

云原生推荐

BashOperator： 直接创建一个pod执行

KubernetesOperator: 直接启动 worker pod+ taskpod, worker里面监听,会创建2个oPod

#### CeleryKubernetesExecutor

上面2个的缝合怪，根据任务队列选择执行器来运行任务

个人不推荐

## 选用方案

kubernetesExecutor + Worker镜像融合方案, 这样worker和task一起创建一起销毁，实现有多大锅下多大米，不用考虑worker节点的数量，worker节点合并到业务负载上

##  容量规划

#### 数据

airflow采用的是一个进程监控一个task,一个进程的内存消耗大约250-280M

## 参数配置

### 全局配置

```sh
# airflow.cfg
[core]
# 全部dag的task上限
parallelism = 32

# 每个dag最大的run数量
max_active_runs_per_dag = 4

# 每个dag的最大task
max_active_tasks_per_dag: 32

[celery]
# 每个 worker 的任务数，k8s环境下建议利用hpa或者keda进行扩展
# 设置worker_auto动态扩容的话，很容易爆内存
worker_concurrency = 8

```

### DAG

如果是串行的话，并发上限`max_active_runs`决定，如果存在并行则`max_active_tasks`可以设置更大值，实现更高的并发task执行

```sh
dag = DAG(
    'my_dag',
    start_date=datetime(2024, 1, 1),
    
    # DAG 级别的并发控制
    max_active_runs = 32
    max_active_tasks = 32
    
    # 连续失败 N 次暂停 DAG
    max_consecutive_failed_dag_runs=5,  # 3.0 新增
)

```

### Task级别

```sh
task = PythonOperator(
    task_id='my_task',
    python_callable=my_function,
    
    # 任务级别控制
    max_active_tis_per_dag=5,  # 此任务在一个DAG中最多5个实例同时运行
    max_active_tis_per_dagrun=2,  # 单个DAG Run中此任务最多2个实例

    # on_failure_callback
    # 实现报警
    # sla_miss_callback
)

```


### pools控制

创建Pools, 然后task中使用，进行控制特定任务的并发

```sh
task = PythonOperator(
    task_id='pooled_task',
    python_callable=my_function,
    pool='my_pool',           # 指定使用的 pool
    pool_slots=2,             # 此任务占用 2 个 pool slots
    priority_weight=10,       # 优先级权重
    wait_for_downstream=True, # 等待下游任务完成
)
```

这个可以实现全局限制运行

任务优先级（Priority Weight）

可以设置进行优先级调度



## 当前问题

1. Airflow不能进行资源感知，智能化调度

airflow不知道底层的资源支持多大的并发，只能根据业务规划好的配置进行调度，资源利用率极低

比如集群只有一个dag运行时，资源闲置，可以动态的把对应的task拉满，加速业务处理

2. Airflow对资源隔离，资源合理分配

比如多个dag时，自带的业务优先级，并不能实现资源隔离，完美的资源分配

目标: 实现资源最大化利用，多业务合理分配，避免饿死， 加速业务处理

3. Airflow的KubernetesPodOperator并不智能

设置的task数量过大， 传统的k8s调度器只会按照requets进行调度，即使request和limit都没问题，但是多个task动态申请内存，运行中容易把节点的内存撑爆，然后节点OOM

目标: 实现动态感知节点负载，负载高就不分配，实现资源保护，避免过载

4. Airflow在k8s环境下，资源申请不合理的话，会导致资源浪费

一些任务request过大，导致无法被调度，就算加机器实现了调度，但是1%的资源利用率机器非常不合理，所以要实现资源超卖，暂时不用的资源可以进行调度，不能像在线业务一样，宁可空着也不能不占

5. NAS的性能，下载处理速度太慢，IO处理太慢，并发太低，需要单独申请(出问题才知道)

目标: 支持本地磁盘挂载调度

## 设计目标

+ 核心目标就是airflow只做dag解析，任务编排和提交
+ 资源调度由专门的调度器进行调度，实现解耦，专业的事情交给专业的人做
+ 支持本地高效磁盘挂载调度
+ 支持自动根据底层资源创建Pod 

### 实现方案

#### 在DAG里实现资源检查

在大规模高并发的情况下这种检查太粗糙，没有分布式一致性，很容易超卖

```sh
with DAG('resource_aware_dag', ...) as dag:
    
    # 1. 预检资源
    check_resources = ResourceAvailabilitySensor(
        task_id='check_resources',
        min_memory_gb=8,
        min_cpu_percent=30,
        min_disk_gb=20,
        poke_interval=60,      # 每 60 秒检查一次
        timeout=3600,          # 1 小时超时
        mode='reschedule',     # 释放 worker slot
    )
    
    # 2. 执行重度任务
    heavy_task = PythonOperator(
        task_id='heavy_task',
        python_callable=resource_intensive_function,
    )
    
    check_resources >> heavy_task
```

#### 第三方调度器

因为k8s的调度器主要面向的是传统微服务体系的，在现在大数据+ML+AI的时代，已经不能满足，所以需要使用面向未来的调度器

以下是主要折腾实践过的华为的Volcano, 阿里的Koordinator

##### Volcano

主要面向批处理，功能很不错，但是文档写的不好，usage感知在他们的公有云明确有，开源代码里面只有设计文档，待验证

+ Queue     支持业务占用不同的Queue，Tree型结构，进行资源精细划分
+ PodGroup  支持Gang调度，实现并行批量计算保证
+ Priority  优先级调度，支持队列优先级，权重调度，避免饿死
+ 支持资源超卖，目前从开源文档没看到usage组件(待验证)
+ VcJob     内嵌task数组，支持的调用优先级，支持指定串行还是并行
+ 抢占  支持优先级抢占

优先级很细，Queue -> VcJob -> Task -> Pod

#### Koordinator

阿里开源，主要面向资源混部，提高资源利用率

+ Koord-Scheduler
+ Koordinator-Manager
+ Koord-Descheduler
+ Koordlet

特性

+ 优先级
+ QoS
+ Job(PodGroup)
+ Device(暂时没折腾GPU深入的调度)
+ 抢占
+ 网络拓补
+ Elastic Quota Management
+ 负载感知调度，重调度，超卖