# [logrus](https://github.com/sirupsen/logrus)

## 主要结构体

Hook            定义hook
Entry           存储数据的
Fields          需要输出的字段
Logger          真正做打印日志的
Formatter       格式日志内容

## Logger

可以定义全局的log，统一输出
var logger = log.New()

## Formatter

自带JSONFormatter, TextFormatter
自定义需要实现Format

```go
type Formatter interface {
    Format(*Entry) ([]byte, error)
}
```

## Fields

结构化输出字段

```go
log.WithFields(log.Fields{
    "name": name,
    "id": id,
}).Fatal("failed")
## Hook
```

定义结构实现这2个接口
每次写入日志时拦截，修改logrus.Entry

```go
type Hook interface {
    Levels() []Level
    Fire(*Entry) error
}
```

```go
type NameHook struct {}

func (n *NameHook) Levels() log.Levels {
    return log.Levels
}

func (n *NameHook) Fire(entry *log.Entry) error {
    entry.Data["name"] = "formych"
    return nil
}
```

添加这个hook
log.AddHook(hook)添加相应的hook

## 设置行号和调用函数名

log.SetReportCaller(true)

## 日志造成os.Exit

使用RegisterExitHandler保证平滑关闭

## 发送日志

将日志发送到日志中心也是logrus所提倡的，第三方hook

* logrus_amqp：Logrus hook for Activemq。
* logrus-logstash-hook:Logstash hook for logrus。
* mgorus:Mongodb Hooks for Logrus。
* logrus_influxdb:InfluxDB Hook for Logrus。
* logrus-redis-hook:Hook for Logrus 。

## 自己实现日志文件分割
