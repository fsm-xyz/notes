# os

## CPU

* 处理器
* CPU核数
* CPU线程数

### 超线程

超线程技术：把一个核模拟成2个，实际同一时刻只能运行一个

比如1个处理器包含2个核，使用超线程技术，则为双核四线程

## CRLF和LF

CR: 回车
LF: 换行

vim -b filename会看到^M,
如何打出:
> Ctrl+V, Ctrl+M打出来 CR
> Ctrl+V, Ctrl+J  LF

vim做了兼容
 :set fileformats=unix,dos
当既有CRLF，LF， CR就会出现^M等

[参考](https://blog.csdn.net/eroswang/article/details/2966082)

## 内核交互

1. 系统调用
2. 系统命令
3. C库
4. 内核函数

### 系统调用

1. 使用C库函数
2. 直接syscall
