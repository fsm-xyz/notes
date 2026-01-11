package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range ch2 {
			fmt.Println("g1", i)
			ch1 <- i
		}
	}()

	ch2 <- 1
	for i := range ch1 {
		fmt.Println("g2", i)
		if i == 9 {
			close(ch2)
			close(ch1)
			break
		}
		ch2 <- i + 1
	}
}
