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
			time.Sleep(2 * time.Second)
			chan1 <- "one"
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			chan2 <- "two"
		}
	}()

	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}

// two
// one
// two
// two
// ^Csignal: interrupt
