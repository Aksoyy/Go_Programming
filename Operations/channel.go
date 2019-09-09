package main

import (
	"fmt"
	"time"
)

func Work(msg string, ch chan string) {
	ch <- msg
	time.Sleep(time.Second * 2)
}

func main() {
	ch1 := make(chan string)
	go Work("work-1", ch1)
	go Work("work-2", ch1)

	for i := 0; i < 2; i++ {
		msg := <-ch1
		fmt.Println(msg)
	}
}
