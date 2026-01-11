package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 测试基于 Cond 的版本
	fmt.Println("=== 基于 Cond 的限制器 ===")
	testCondLimiter()

	fmt.Println("\n=== 基于 Channel 的限制器 ===")
	testChannelLimiter()
}

func testCondLimiter() {
	l := NewCondLimiter(3)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			l.Allow()
			fmt.Printf("Goroutine %d 开始执行\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine %d 执行完成\n", id)
			l.Release()
		}(i)
	}

	wg.Wait()
}

func testChannelLimiter() {
	l := NewChannelLimiter(3)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			l.Allow()
			fmt.Printf("Goroutine %d 开始执行\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine %d 执行完成\n", id)
			l.Release()
		}(i)
	}

	wg.Wait()
}

// 基于 Cond 的实现
type CondLimiter struct {
	max  int64
	num  atomic.Int64
	cond *sync.Cond
	mu   sync.Mutex
}

func NewCondLimiter(max int64) *CondLimiter {
	l := &CondLimiter{
		max: max,
	}
	l.cond = sync.NewCond(&l.mu)
	return l
}

func (l *CondLimiter) Allow() {
	l.mu.Lock()
	defer l.mu.Unlock()

	for l.num.Load() >= l.max {
		l.cond.Wait()
	}
	l.num.Add(1)
}

func (l *CondLimiter) Release() {
	l.num.Add(-1)
	l.cond.Broadcast()
}

// 基于 Channel 的实现（更高性能）
type ChannelLimiter struct {
	ch chan struct{}
}

func NewChannelLimiter(max int64) *ChannelLimiter {
	return &ChannelLimiter{
		ch: make(chan struct{}, max),
	}
}

func (l *ChannelLimiter) Allow() {
	l.ch <- struct{}{} // 会阻塞直到有可用槽位
}

func (l *ChannelLimiter) Release() {
	<-l.ch // 释放一个槽位
}

// 更简洁的使用方式：使用 defer
func (l *ChannelLimiter) Do(fn func()) {
	l.Allow()
	defer l.Release()
	fn()
}
