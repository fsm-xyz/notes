# 数据库的分表分库

单机单实例, 单机多实例, 多机单实例, 多机多实例

## 分表字段uid

db_num
table_num
total_table_num = db_num * table_num

### 计算规则

#### 先分表

获取具体的表
uid % total_table_num

获取具体的库
(uid % total_table_num) / table_num

#### 先分库

求库
uid % db_num
求表
(uid / db_num) % table_num

### 按时间分

### 按id数量区间
