package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			chan1 <- "one"
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			chan2 <- "two"
		}
	}()

	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 1):
			fmt.Println("Test message")
			return
		}
	}
}

// Test message
