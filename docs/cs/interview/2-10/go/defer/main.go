package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer fmt.Println(i)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
