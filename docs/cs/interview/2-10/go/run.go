package main

import (
	"fmt"
	"sync"
)

func main() {
	// ChanRun()
	WgRun()
}

var (
	curr     = 2
	num      = 100
	taskChan = make(chan Task, num)
	done     = make(chan struct{})

	tasks = []Task{
		func() { fmt.Println("Task 1") },
		func() { fmt.Println("Task 2") },
		func() { fmt.Println("Task 3") },
		func() { fmt.Println("Task 4") },
		func() { fmt.Println("Task 5") },
	}
)

type Task func()

func ChanRun() {
	go run()
	for _, t := range tasks {
		taskChan <- t
	}
	close(taskChan)
	for i := 0; i < len(tasks); i++ {
		<-done
	}
}
func run() {
	for i := 0; i < curr; i++ {
		go func() {
			for t := range taskChan {
				t()
				done <- struct{}{}
			}
		}()
	}
}

func WgRun() {
	wg := sync.WaitGroup{}

	for i := 0; i < curr; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range taskChan {
				t()
			}
		}()
	}

	for _, t := range tasks {
		taskChan <- t
	}

	close(taskChan)

	wg.Wait()
}
