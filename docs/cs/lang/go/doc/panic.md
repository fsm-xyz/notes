# panic

## 注意点

+ recover只捕获当前goroutinue的panic
+ recover必须放在当前函数里执行，不能跨函数
