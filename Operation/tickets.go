package main

import "time"
import "fmt"

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// Tick at 2019-08-29 13:55:01.883827 +0300 +03 m=+0.505442651
// Tick at 2019-08-29 13:55:02.380503 +0300 +03 m=+1.002104251
// Tick at 2019-08-29 13:55:02.88258 +0300 +03 m=+1.504166670
// Ticker stopped
