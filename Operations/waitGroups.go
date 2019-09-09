package main

import (
	"fmt"
	"sync"
	"time"
)

func Work(msg string, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println(msg)
	wg.Done()
}

func main() {
	number1, number2 := 5, 4
	wg := &sync.WaitGroup{}

	wg.Add(number1)
	for i := 0; i < number1; i++ {
		go Work("work-1", wg)
	}
	wg.Wait()

	wg.Add(number2)
	for i := 0; i < number2; i++ {
		go Work("work-2", wg)
	}
	wg.Wait()
}

// work-1
// work-1
// work-1
// work-1
// work-1
// work-2
// work-2
// work-2
// work-2
