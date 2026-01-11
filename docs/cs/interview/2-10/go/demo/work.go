package main

import (
	"fmt"
	"sync"
)

// 启动 3 个 worker goroutine
// 处理任务主线程通过 channel
//  发送 10 个任务（任务内容为整数 1-10）
// 每个 worker 处理任务时，
// 输出 “worker  处理任务 ”所有任务处理完成后，
// 主线程输出 “所有任务处理完毕” 并退出

func main() {
	workerNums := 3
	workerQ := make(chan int, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		workerQ <- i
	}
	close(workerQ)
	// wg.Add(workerNums)
	for i := 0; i < workerNums; i++ {
		// go func() {
		// 	defer wg.Done()
		// 	for w := range workerQ {
		// 		fmt.Printf("worker:%d,  处理任务: %d\n", i, w)
		// 	}
		// }()
		wg.Go(func() {
			for w := range workerQ {
				fmt.Printf("worker:%d,  处理任务: %d\n", i, w)
			}
		})

		// 利泽
		// 数字货币研究所

	}
	wg.Wait()

	a := []int{1, 3, 5}
	b := []int{2, 4, 6, 8}
	fmt.Println(merge(a, b))
}

