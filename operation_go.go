package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	//fmt.Scanln()
	//fmt.Println("done")
}

// direct : 0
// direct : 1
// direct : 2
// goroutine : 0
// goroutine : 1
// Her seferinde goroutine değeri degisebilmektedir.
