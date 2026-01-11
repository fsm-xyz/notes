# influxdb

## core3.0

arrow 统一内存格式
parquet 高压缩比的
datafusion

无 WAL 设计 + 对象存储直写

ZSTD

## 其他

LSM Tree 写入流程
写入 WAL：数据先追加到 Write-Ahead Log（确保持久化）。

写入 MemTable：内存中的有序结构（如跳表），写满后转为 Immutable MemTable。

刷盘为 SSTable：Immutable MemTable 异步刷到磁盘，形成 Level 0 SSTable（文件间时间范围重叠）。

Compaction：后台合并重叠的 SSTable 到更高层级（Level 1→N），文件更大且无重叠。

TSM Tree 写入流程（InfluxDB 实现）
写入 WAL：保证数据持久化。

写入 Cache：内存中的 Map（Key=时序ID，Value=数据点列表）。

刷盘为 TSM File：Cache 写满后转为 TSM 文件（按时间分块，列式存储）。

Compaction：合并小 TSM 文件，优化查询效率，生成索引文件（.tsi）加速按标签检索。