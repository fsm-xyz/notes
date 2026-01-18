# 容器进程统计

## 方法 1: 使用 kubectl exec 进入 Pod 查看

```bash
# 查看进程列表
kubectl exec <pod-name> -- ps aux

# 统计进程数量
kubectl exec <pod-name> -- ps aux | wc -l

# 如果 Pod 有多个容器，需要指定容器名
kubectl exec <pod-name> -c <container-name> -- ps aux
```

## 方法 2: 使用 top 命令

```bash
# 查看实时进程信息
kubectl exec <pod-name> -- top -b -n 1

# 或使用 ps 的简洁输出
kubectl exec <pod-name> -- ps -ef
```

## 方法 3: 统计特定类型的进程

```bash
# 只显示进程数量（不含表头）
kubectl exec <pod-name> -- ps -ef | tail -n +2 | wc -l

# 查看进程树
kubectl exec <pod-name> -- ps auxf

# 查看特定进程
kubectl exec <pod-name> -- ps aux | grep <process-name>
```

## 方法 4: 如果已经在 Pod 内

如果您已经通过 `kubectl exec -it <pod-name> -- /bin/sh` 进入 Pod：

```bash
# 查看所有进程
ps aux

# 统计进程数
ps aux | wc -l

# 或
ps -ef | wc -l

# 查看进程树结构
pstree -p
```

## 常见问题处理

如果遇到 `ps` 命令不存在的情况（某些精简镜像），可以尝试：

```bash
# 查看 /proc 目录（每个数字目录代表一个进程）
ls /proc | grep -E '^[0-9]+$' | wc -l

# 或直接查看
ls -d /proc/[0-9]*
```

