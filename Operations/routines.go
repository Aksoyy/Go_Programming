package main

import (
	"fmt"
	"time"
)

func main() {
	// Work("test test test test")
	// Work("deneme")

	go Work("test test test test")
	go Work("deneme")
	time.Sleep(time.Second * 3)
}

func Work(msg string) {
	fmt.Println(msg)
	time.Sleep(time.Second * 2)
}

// ➜ Operations git:(master) ✗ go run routines.go
// test test test test
// deneme
// ➜ Operations git:(master) ✗ go run routines.go
// deneme
// test test test test
