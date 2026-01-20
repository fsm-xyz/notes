# CPU

## CPU 使用率限流（Cgroup CPU 带宽限制）

PAI 和 ACK 都默认开启

## CPU 核心绑定 / 亲和性（CPU Core Pinning /cpuset 核心隔离）

PAI-DSW 默认开启，ACK 默认关闭
写入在`/sys/fs/cgroup/cpuset/cpuset.cpus`

CPU Manager Policy = static

## 差异

+ PAI-DSW 的调度策略：严格绑定 CPU 核心
+ ACK 的调度策略：cfs,时间片分配


nproc/lscpu/top/htop 查看核数

## 实现

### kubelete

使用 K8s 原生的 `CPU Manager Policy = static` 模式（推荐，和 PAI-DSW 一致）

开启 static 模式后，K8s 的 CPU Manager 会对「申请整核 CPU 的容器」（limits是整数，比如 1、2、4），自动做 cpuset 核心绑定，把容器钉死在固定的 CPU 核心上，和 PAI-DSW 的效果完全一样。

```sh

--cpu-manager-policy=static --kube-reserved=cpu=1,memory=1Gi --system-reserved=cpu=1,memory=1Gi
```

### 命名空间

```bash
apiVersion: v1
kind: Namespace
metadata:
  name: ai-workspace  # 你的AI业务命名空间
  annotations:
    cpu-manager.kubernetes.io/policy: static
```

### pod

```bash
apiVersion: v1
kind: Pod
metadata:
  name: ai-task
spec:
  containers:
  - name: pytorch
    image: pytorch/pytorch:latest
    resources:
      limits:
        cpu: "2"
        memory: "8Gi"
      requests:
        cpu: "2"
        memory: "8Gi"
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: kubernetes.io/hostname
            operator: In
            values: ["ecs-96core-01"]  # 指定宿主机
  tolerations:
  - key: cpuset
    operator: Exists
```