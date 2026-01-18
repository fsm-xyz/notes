# Airflow 配置优先级

根据官方文档，Airflow 配置的通用优先级顺序如下：
1. 环境变量 (AIRFLOW__DATABASE__SQL_ALCHEMY_CONN)
2. 命令环境变量 (AIRFLOW__DATABASE__SQL_ALCHEMY_CONN_CMD)
3. Secret 环境变量 (AIRFLOW__DATABASE__SQL_ALCHEMY_CONN_SECRET)
4. airflow.cfg 配置文件中的设置
5. Airflow 内置默认值

### 简单来说：**环境变量优先级最高！**

环境变量不会改变配置文件的内容，它们会覆盖配置文件中的值。Airflow 在读取配置值时会首先检查是否定义了环境变量，如果有，就会使用环境变量而不是配置文件。

## 具体说明

### 1. **环境变量格式**
```bash
AIRFLOW__{SECTION}__{KEY}
```

例如：
```bash
# 覆盖 [celery] 部分的 worker_concurrency
export AIRFLOW__CELERY__WORKER_CONCURRENCY=32

# 覆盖 [core] 部分的 parallelism
export AIRFLOW__CORE__PARALLELISM=64
```

### 2. **配置文件格式**
```ini
[celery]
worker_concurrency = 16

[core]
parallelism = 32
```

### 3. **优先级实例**

假设你有以下配置：

**airflow.cfg 文件：**
```ini
[celery]
worker_concurrency = 16
```

**环境变量：**
```bash
export AIRFLOW__CELERY__WORKER_CONCURRENCY=32
```

**实际生效值：** `32` （环境变量覆盖了配置文件）

## 验证当前生效的配置

使用以下命令查看实际生效的配置：

```bash
# 方法 1: 列出所有配置（显示实际生效值）
airflow config list

# 方法 2: 获取特定配置项
airflow config get-value celery worker_concurrency

# 方法 3: 在 Python 中验证
python -c "from airflow.configuration import conf; print(conf.get('celery', 'worker_concurrency'))"
```

## 最佳实践建议

没有"最好"的方式。使用环境变量覆盖适合需要轻松更改的场景，修改和挂载配置文件可能适合更"持久"的方式，但你也可以构建包含替换后的 airflow.cfg 的自定义镜像（这是最"持久"的方式）。

### 推荐策略：

1. **开发/测试环境**：使用环境变量，方便快速调整
2. **生产环境**：
   - 基础配置：写在 `airflow.cfg` 文件中
   - 敏感信息：使用环境变量或 Secret
   - 需要动态调整的参数：使用环境变量

### 在你的 Kubernetes 环境中

从你的日志可以看出使用的是 Kubernetes，建议：

```yaml
# 在 Pod 配置中设置环境变量
env:
  - name: AIRFLOW__CELERY__WORKER_CONCURRENCY
    value: "32"
  - name: AIRFLOW__CORE__PARALLELISM
    value: "64"
```

这样可以覆盖默认配置，且易于管理和版本控制。
