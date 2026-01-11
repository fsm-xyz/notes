package main

import (
	"fmt"
	"time"

)

func main(){
	t := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-t)
	}()

	t<-1
	close(t)
	fmt.Println("close t")

	if t, ok := <- t ; ok{
		fmt.Println(t, ok)
	}
}