package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	// channel creating

	queue <- "one"
	queue <- "two"
	// add strings in channel
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// one
// two
