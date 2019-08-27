package main

import "fmt"

func main() {
	mychanel := make(chan string, 1)
	mychanel <- "Hakan"

	fmt.Println(mychanel, <-mychanel)
}

// 0xc00007e0c0 Hakan
