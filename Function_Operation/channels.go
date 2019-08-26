package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping_test" }()

	msg := <-messages
	fmt.Println(msg)
}

// ping_test
