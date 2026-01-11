# prometheus

## 数据内容

+ 指标: name和label组成, <metric name>{<label name>=<label value>, ...}, metric name的命名规则为：应用名称开头_监测对像_数值类型_单位
+ 时间戳: 精确到毫秒的时间戳
+ 样本值: float64

## 数据类型

+ Counter
+ Gauge(仪表盘类型)
+ Histogram(直方图类型)
+ Summary(摘要类型)

### Couter

只增不减

+ 总和

+ 获取增长率

rate(http_requests_total[5m])

+ 查询当前系统中，访问量前10的HTTP地址

topk(10, http_requests_total)

### Gauge(仪表盘类型)

可增可减

+ 计算CPU温度在两小时内的差异

dalta(cpu_temp_celsius{host="zeus"}[2h])

+ 预测系统磁盘空间在4小时之后的剩余情况

predict_linear(node_filesystem_free{job="node"}[1h], 4*3600)

### Histogram(直方图类型)

+ 事件发生的总次数，basename_count

io_namespace_http_requests_latency_seconds_histogram_count{path="/",method="GET",code="200",} 2.0

+ 所有事件产生值的大小的总和，basename_sum

io_namespace_http_requests_latency_seconds_histogram_sum{path="/",method="GET",code="200",} 13.107670803000001

+ 事件产生的值分布在bucket中的次数，basename_bucket{le=“上包含”}
io_namespace_http_requests_latency_seconds_histogram_bucket{path="/",method="GET",code="200",le="0.005",} 0.0

### Summary(摘要类型)


Summary类型和Histogram类型相似
+ 都包含 < basename>_sum和< basename>_count;
+ Histogram需要通过< basename>_bucket计算quantile，而Summary直接存储了quantile的值。

+ 事件发生总的次数

io_namespace_http_requests_latency_seconds_summary_count{path="/",method="GET",code="200",} 12.0

+ 事件产生的值的总和

io_namespace_http_requests_latency_seconds_summary_sum{path="/",method="GET",code="200",} 51.029495508

+ http请求响应时间的中位数是3.052404983s

io_namespace_http_requests_latency_seconds_summary{path="/",method="GET",code="200",quantile="0.5",} 3.052404983

+ 这12次http请求响应时间的9分位数是8.003261666s

io_namespace_http_requests_latency_seconds_summary{path="/",method="GET",code="200",quantile="0.9",} 8.003261666