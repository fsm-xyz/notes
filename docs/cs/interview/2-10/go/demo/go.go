package main
import (
"fmt"
"time"
)
func main() {
	ch1 := make(chan int)

	go fmt.Println(<-ch1)
	// 1. 主goroutine执行到这一行
	// 2. Go首先计算参数：<-ch1
	// 3. 主goroutine尝试从ch1接收数据，但通道为空
	// 4. 主goroutine阻塞在这里，永远不会继续执行后面的代码

	// go func(){
	// 	fmt.Println(<-ch1)
	// }()
	ch1 <- 5
	time.Sleep(1 * time.Second)
}
