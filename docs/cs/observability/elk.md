# ELK

ES + logstash + kibana

ClickHouse + Victoria + OpenOberseve

Grafana Loki

Promtail 

influxDB 3 在架构演进中增强了对日志（Logs）和追踪（Traces）

## trace

ClickHouse

influxDB 3

Loki

OpenObeserve

Elasticsearch

冷热分离。长期数据存放OSS

## logs

SLS + OSS

Elasticsearch

VictoriaLogs

ClickHouse

Elasticsearch LogsDB

Loki + 对象存储

日志： 多行日志容易错乱，1. 用正则匹配，2. 代码包含成一行


# 存储

时序数据库
列式数据库
对象存储

SSD（热数据）+ 对象存储（冷数据）

在 Elasticsearch 中实现数据冷热分离，并加载冷数据，主要涉及到 Elasticsearch 的**索引生命周期管理 (ILM)** 功能。ILM 自动化了索引从热到冷再到归档或删除的全过程，是实现冷热分离最常用和推荐的方式。

### 数据冷热分离的原理

Elasticsearch 数据冷热分离的核心思想是根据数据的新旧，将其存储在不同性能的硬件上。

  * **热数据 (Hot Tier):** 存储最新写入和最常查询的数据。这些数据通常在高性能硬件上，例如 **SSD**。热数据节点主要负责写入和频繁的搜索。
  * **温数据 (Warm Tier):** 存储只读或查询频率较低的数据。这些数据通常在性能稍逊的硬件上，例如 **大容量 HDD**。数据从热阶段自动迁移到温阶段后，只允许查询，不再写入。
  * **冷数据 (Cold Tier):** 存储很少访问但需要保留的数据。这些数据通常在更廉价的存储上，例如 **大容量 HDD 或对象存储**。冷数据节点只负责归档和偶尔的查询，查询性能会较低。

### 如何加载冷数据

加载冷数据（即查询冷数据）的过程是自动的，不需要额外的操作。当你使用 ILM 将索引从热阶段迁移到温阶段或冷阶段时，数据仍然是 Elasticsearch 集群的一部分，并且可以被 **透明地查询**。

你只需要像平常一样发送你的查询请求，Elasticsearch 会自动路由到存储该数据的节点。

-----

### 实现步骤

下面是使用 ILM 实现冷热分离并加载冷数据的详细步骤：

#### 1\. 配置节点角色

首先，你需要为集群中的不同节点设置角色。这是区分热、温、冷节点的关键。

在 `elasticsearch.yml` 配置文件中，为不同节点设置 `node.roles`：

  * **热节点：** 具备数据写入和搜索能力。
    ```yaml
    node.roles: [ master, data_hot ]
    ```
  * **温节点：** 具备搜索能力，但通常不再接受写入。
    ```yaml
    node.roles: [ data_warm ]
    ```
  * **冷节点：** 仅负责归档和搜索。
    ```yaml
    node.roles: [ data_cold ]
    ```

配置完成后，重启节点让配置生效。

#### 2\. 创建 ILM 策略

接下来，你需要创建一个 ILM 策略来定义索引生命周期的每个阶段以及迁移条件。你可以通过 Kibana UI 或者 API 来创建。

这个策略定义了索引何时从一个阶段移动到下一个阶段。

**示例：一个简单的 ILM 策略**

```json
PUT /_ilm/policy/my-ilm-policy
{
  "policy": {
    "phases": {
      "hot": {
        "actions": {
          "rollover": {
            "max_age": "7d",
            "max_docs": 10000000
          }
        }
      },
      "warm": {
        "min_age": "7d",
        "actions": {
          "set_priority": {
            "priority": 50
          },
          "shrink": {
            "number_of_shards": 1
          }
        }
      },
      "cold": {
        "min_age": "30d",
        "actions": {
          "set_priority": {
            "priority": 0
          },
          "forcemerge": {
            "max_num_segments": 1
          }
        }
      },
      "delete": {
        "min_age": "90d",
        "actions": {
          "delete": {}
        }
      }
    }
  }
}
```

  * **`hot` 阶段：** 当索引达到 7 天或者包含 1000 万个文档时，触发 `rollover` 操作，创建一个新的索引。
  * **`warm` 阶段：** 在 `rollover` 之后，当索引达到 7 天时，自动将索引移动到温节点。
  * **`cold` 阶段：** 当索引达到 30 天时，自动将索引移动到冷节点。
  * **`delete` 阶段：** 当索引达到 90 天时，自动删除索引。

#### 3\. 将策略应用于索引模板

最后，你需要将这个 ILM 策略与一个索引模板关联起来。这样，所有基于该模板创建的新索引都会自动遵循这个生命周期策略。

```json
PUT /_index_template/my-data-template
{
  "index_patterns": ["my-data-*"],
  "template": {
    "settings": {
      "index": {
        "lifecycle": {
          "name": "my-ilm-policy"
        },
        "routing": {
          "allocation": {
            "require": {
              "data": "hot"
            }
          }
        }
      }
    }
  }
}
```

  * **`index_patterns`:** 匹配 `my-data-` 开头的所有索引。
  * **`lifecycle.name`:** 将 `my-ilm-policy` 策略应用于这些索引。
  * **`routing.allocation.require`:** 这是一个关键设置，它指定新创建的索引必须首先分配到 **`data: hot`** 的节点上。之后，ILM 会根据策略自动将索引迁移到其他类型的节点。

### 如何查询冷数据？

一旦以上步骤设置完成，你就可以像往常一样查询数据了。当你的查询请求到来时，Elasticsearch 的协调节点（Coordinating Node）会知道哪个索引（例如 `my-data-2023-01-01`）处于温或冷阶段，并根据其分片分配情况，自动将请求路由到温或冷数据节点进行查询。

对于用户而言，这个过程是完全透明的，你不需要关心数据具体存储在哪里。只不过，查询冷数据时，由于其存储在性能较低的硬件上，查询延迟可能会比热数据要高。