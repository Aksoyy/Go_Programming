package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Work(work_id int, msg string, wg *sync.WaitGroup, mx *sync.Mutex, results map[string]string) {
	mx.Lock()
	time.Sleep(time.Second * 1)
	worker_id := msg + "-" + strconv.Itoa(work_id)
	results[worker_id] = msg
	mx.Unlock()

	wg.Done()
}

func main() {

	number1, number2 := 5, 4
	wg := &sync.WaitGroup{}
	mx := &sync.Mutex{}
	results := make(map[string]string, number1+number2)

	wg.Add(number1)
	for i := 0; i < number1; i++ {
		go Work(i, "work-1", wg, mx, results)
	}
	wg.Wait()

	fmt.Println(results)

	wg.Add(number2)
	for i := 0; i < number2; i++ {
		go Work(i, "work-2", wg, mx, results)
	}
	wg.Wait()
	fmt.Println(results)
}

// map[work-1-0:work-1 work-1-1:work-1 work-1-2:work-1 work-1-3:work-1 work-1-4:work-1]
// map[work-1-0:work-1 work-1-1:work-1 work-1-2:work-1 work-1-3:work-1 work-1-4:work-1 work-2-0:work-2 work-2-1:work-2 work-2-2:work-2 work-2-3:work-2]

// map[work-1-0:work-1 work-1-1:work-1 work-1-2:work-1 work-1-3:work-1 work-1-4:work-1]
// map[work-1-0:work-1 work-1-1:work-1 work-1-2:work-1 work-1-3:work-1 work-1-4:work-1 work-2-0:work-2 work-2-1:work-2 work-2-2:work-2 work-2-3:work-2]
