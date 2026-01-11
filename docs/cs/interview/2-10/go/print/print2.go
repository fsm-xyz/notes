package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	a := make(chan int)
	b := make(chan int)

	wg.Go(func() {
		for i := range a {
			fmt.Printf("A received: %d, sending: %d\n", i, i+1)
			b <- i + 1
		}
		close(b)
		fmt.Println("Goroutine A finished.")
	})

	wg.Go(func() {
		fmt.Println("B starting, sending: 1")
		a <- 1

		for i := range b {
			fmt.Printf("B received: %d\n", i)
			if i >= 10 {
				fmt.Println("Termination condition (i >= 10) met.")
				close(a)
				break
			}
			fmt.Printf("B sending: %d\n", i+1)
			a <- i + 1
		}
		fmt.Println("Goroutine B finished.")
	})

	fmt.Println("Main: Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Main: All done.")
}
