package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Work(work_id int, msg string, wg *sync.WaitGroup, results map[string]string) {
	time.Sleep(time.Second * 1)
	worker_id := msg + "-" + strconv.Itoa(work_id)
	results[worker_id] = msg
	wg.Done()
}

func main() {

	nWorkers1, nWorkers2 := 5, 4
	wg := &sync.WaitGroup{}
	results := make(map[string]string, nWorkers1+nWorkers2)

	wg.Add(nWorkers1)

	for i := 0; i < nWorkers1; i++ {
		go Work(i, "work-1", wg, results)
	}

	wg.Wait()
	fmt.Println(results)
	wg.Add(nWorkers2)

	for i := 0; i < nWorkers2; i++ {
		go Work(i, "work-2", wg, results)
	}
	wg.Wait()
	fmt.Println(results)
}

// ➜  Operations git:(master) ✗ go run mutex.go
// map[work-1-0:work-1 work-1-1:work-1 work-1-2:work-1 work-1-3:work-1 work-1-4:work-1]
// map[work-1-0:work-1 work-1-1:work-1 work-1-2:work-1 work-1-3:work-1 work-1-4:work-1 work-2-0:work-2 work-2-1:work-2 work-2-2:work-2 work-2-3:work-2]

// ➜  Operations git:(master) ✗ go run mutex.go
// fatal error: concurrent map writes
// goroutine 36 [running]:
// runtime.throw(0x10c93a9, 0x15)
// runtime.mapassign_faststr(0x10ab4e0, 0xc000088000, 0xc000018070, 0x8, 0x1)
// main.Work(0x3, 0x10c716a, 0x6, 0xc000086000, 0xc000088000)
// runtime.goexit()
// created by main.main
// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000086008)
// sync.(*WaitGroup).Wait(0xc000086000)
// main.main()
// exit status 2
