# 删除

## 批量某一类删除

```bash
redis-cli keys "name*" | xargs redis-cli del
```

## 本db的删除

```bash
flushdb
```

## 全部的db删除

```bash
flushall
```

注意:
    慎用，有危险性，会阻塞其他的操作
