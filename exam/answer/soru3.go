// subs içerisinde bulunan herkese, notify fonksiyonunu kullanarak,
// dead-lock oluimadan, goroutineler ile mesaj gönderebilen kodu yazın.
// (sync.WaitGroup)
package main

import (
	"fmt"
	"sync"
)

func Notify(to, message string, wg *sync.WaitGroup) {
	fmt.Println("sending message:", message, "to:", to)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	subs := []string{
		"user1",
		"user2",
		"user3",
	}
	wg.Add(len(subs))

	for _, sub := range subs {
		go Notify(sub, "hello", wg)
	}

	wg.Wait()
}
