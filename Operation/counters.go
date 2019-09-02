package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {

				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

// ➜ go run counters.go
// ops: 36300
// ➜ go run counters.go
// ops: 36150
// ➜ go run counters.go
// ops: 36299
// ➜ go run counters.go
// ops: 36402
