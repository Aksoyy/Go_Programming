// subs içerisinde bulunan herkese, notify fonksiyonunu kullanarak,
// dead-lock oluimadan, goroutineler ile mesaj gönderebilen kodu yazın.
// (sync.WaitGroup)
package main

import (
	"fmt"
	"sync"
	"time"
)

func Notify(to, msg string) {
	fmt.Println("sending message:", msg, "to:", to)
}

func Work(person, msg string, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	Notify(person, msg)
	wg.Done()
}

func main() {
	subs := []string{
		"user1",
		"user2",
		"user3",
	}
	msg := "hello"

	wg := &sync.WaitGroup{}
	wg.Add(3)
	for _, sub := range subs {
		go Work(sub, msg, wg)
	}
	wg.Wait()
}

// sending message: hello to: user3
// sending message: hello to: user1
// sending message: hello to: user2
