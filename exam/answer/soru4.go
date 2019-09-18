// select yapısını kullanarak,
// Hammer fonksiyonunun 3 saniyeden fazla sürmesi durumunda
// main fonksiyonunu bitirin.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Hammer cranks the CPU for a random duration between 0-10
// and sends the duration to given channel
func Hammer(ch chan int) {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- r
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	go Hammer(ch)

	select {
	case res := <-ch:
		fmt.Println("Finished after", res, "seconds.")
	case <-time.After(time.Second * 1):
		return
	}

}
