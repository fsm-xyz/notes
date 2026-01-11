# 调试

## 接口

```sh
# 强制开启trace日志
curl -X POST 127.0.0.1:15000/logging?level=trace

# 列出所有已加载的TLS证书，包括文件名，序列号，主题备用名称以及符合证书原型定义的 JSON格式到期的天数
curl -X POST 127.0.0.1:15000/certs

# 列出所有已配置的集群管理器集群，此信息包括每个群集中发现的所有上游主机以及每个主机统计信息
curl -X POST 127.0.0.1:15000/clusters

# 将当前Envoy从各种组件加载出来的的配置转储为JSON进行输出
curl -X POST 127.0.0.1:15000/config_dump

# 如果启用了互斥跟踪，则以 JSON格式输出当前的Envoy互斥争用统计信息（MutexStats）
curl -X POST 127.0.0.1:15000/contention

# 启用或禁用CPU Profiler。需要使用gperftools进行编译。输出文件可以由admin.profile_path配置
curl -X POST 127.0.0.1:15000/cpuprofiler

# drain listeners
curl -X POST 127.0.0.1:15000/drain_listeners

# cause the server to fail health checks
curl -X POST 127.0.0.1:15000/healthcheck/fail

# cause the server to pass health checks
curl -X POST 127.0.0.1:15000/healthcheck/ok

# enable/disable the heap profiler
curl -X POST 127.0.0.1:15000/heapprofiler

# print out list of admin commands
curl -X POST 127.0.0.1:15000/help

# print the hot restart compatibility version
curl -X POST 127.0.0.1:15000/hot_restart_version

# dump current Envoy init manager information (experimental)
curl -X POST 127.0.0.1:15000/init_dump

# print listener addresses
curl -X POST 127.0.0.1:15000/listeners

# query/change logging levels
curl -X POST 127.0.0.1:15000/logging

# print current allocation/heap usage
curl -X POST 127.0.0.1:15000/memory

# 退出服务器
curl -X POST 127.0.0.1:15000/quitquitquit

# print server state, return 200 if LIVE, otherwise return 503
curl -X POST 127.0.0.1:15000/ready

# reopen access logs
curl -X POST 127.0.0.1:15000/reopen_logs

# reset all counters to zero
curl -X POST 127.0.0.1:15000/reset_counters

# print runtime values
curl -X POST 127.0.0.1:15000/runtime

# modify runtime values
curl -X POST 127.0.0.1:15000/runtime_modify

# print server version/status information
curl -X POST 127.0.0.1:15000/server_info

# print server stats
curl -X POST 127.0.0.1:15000/stats

# print server stats in prometheus format
curl -X POST 127.0.0.1:15000/stats/prometheus

# Show recent stat-name lookups
curl -X POST 127.0.0.1:15000/stats/recentlookups

# clear list of stat-name lookups and counter
curl -X POST 127.0.0.1:15000/stats/recentlookups/clear

# disable recording of reset stat-name lookup names
curl -X POST 127.0.0.1:15000/stats/recentlookups/disable

# enable recording of reset stat-name lookup names
curl -X POST 127.0.0.1:15000/stats/recentlookups/enable
```
