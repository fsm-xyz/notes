// 要求：
// 1. 只能编辑 foo 函数
// 2. foo 必须要调用 slow 函数
// 3. foo 函数在 ctx 超时后必须立刻返回
// 4. 【加分项】如果 slow 结束的比 ctx 快，也立刻返回

package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	fmt.Println("Before Go Num: ", runtime.NumGoroutine())
	foo(ctx)
	fmt.Println("After Go Num: ", runtime.NumGoroutine())
	time.Sleep(10 * time.Second)
	fmt.Println("End Go Num: ", runtime.NumGoroutine())
}

// 您需要实现 foo 函数，要求 foo 在 ctx 超时后立即返回
// foo 必须调用 slow 函数
func foo(ctx context.Context) {
	done := make(chan struct{}, 1)
	go func() {
		slow()
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return
	case <-done:
		return
	}
}

func foo2(ctx context.Context) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    
    go func() {
        slow()
        cancel()
    }()
    
    <-ctx.Done()
}

// 您不能修改 slow 函数
func slow() {
	n := rand.Intn(10)
	fmt.Printf("sleep %ds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
