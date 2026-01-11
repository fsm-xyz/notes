# VictoriaMetrics

基于LSM的时序数据库，高压缩比(Gorilla + ZSTD)

LSM树是一种非常高效的数据结构，专为写入密集型的应用场景（如数据库、搜索引擎、日志系统）而设计。它牺牲了一定的读取性能，来换取极高的写入吞吐量

化随机为顺序

vminsert，vmstorage, vmselect, vmagent, vmalert

Counter, Histogram, Summary

metric和label

先存为WAL实现持久性，然后写入内存mergeTable, 写入不可变的memTable, 后台进行落盘，写入0层的table, 后台再进合并分成table, 成成更大的

LevelDB / RocksDB


## 数据模型

单值模型（Single-Value）：每个时间序列（Metric + Labels）独立存储，仅支持浮点数值。相比多值模型（如 InfluxDB），更适配监控场景的高维度标签查询

列式数据库，适合做聚合计算

一致性哈希实现负载均衡

无WAL(Write-AheadLog)


索引与数据分离：

索引：生成唯一 TSID（Time Series ID），映射 Metric+Labels 到数据块。

数据：按 TSID 分组，列式存储时间戳与数值，使用 Delta-of-Delta + Gorilla 压缩算法减少存储空间