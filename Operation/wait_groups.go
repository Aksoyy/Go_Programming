package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}

// Worker 5 starting
// Worker 1 starting
// Worker 2 starting
// Worker 3 starting
// Worker 4 starting
// Worker 5 done
// Worker 4 done
// Worker 2 done
// Worker 3 done
// Worker 1 done
