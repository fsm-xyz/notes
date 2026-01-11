# ES命令

## 迁移

### 获取当前集群的所有settings

```shell
curl -XGET localhost:9200/_cluster/settings
```

### 获取当前集群的所有mapping

```shell
curl -XGET localhost:9200/_mapping
```

### 获取当前集群的所有脚本

```shell
curl 'http://localhost:9200/_cluster/state/metadata?pretty&filter_path=**.stored_scripts'
```

### 获取当前集群的所有index

```shell
curl -XGET localhost:9200/_cat/indices?v
```

### 删除文档

```shell
curl -XDELETE -H 'Content-Type: application/json' 'http://localhost:9200/{index}'
```

### elasticdump

--type
    + mapping
    + settings
    + analyzer

```shell
elasticdump   --input=http://localhost:9200/{index} --output=a.json  --type=mapping
```
