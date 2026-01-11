# Memory

## 排查方向

1. sync.Mutex上锁逻辑异常，锁等待
2. WaitGroup问题
3. chan等待
4. goroutine泄漏
5. 死循环
6. 资源等待造成的Goroutinue hang住

## 参考

- [知乎](https://zhuanlan.zhihu.com/p/74090074)
- [felix](https://www.felix021.com/blog/read.php?2208)
