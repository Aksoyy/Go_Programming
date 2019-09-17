# Metglobal GO Sınavı

## Açıklamalar
- İlk 4 soruyu bilgisayarınızda çalıştırmak için `go run soru<soru_numarası>.go`, 5. Soruyu çalıştırmak için `go test soru5_test.go` komutunu kullanabilirsiniz.

- Kodları yazmak için https://play.golang.org/, kaynak için bit.ly/2keioJP adresini kullanabilirsiniz.

- Cevaplarınızı sadece soru dosyalarına yazın.

- Soru dosyalarını cevaplarınızla doldurduktan sonra, tüm klasörü `.zip` formatında sıkıştırıp
`kaan.kumkale@metglobal.com` adresine gönderin.


##  Sorular

1- `Read` fonksiyonu içerisinde dosyanın kapalı olup olmadığını 
kontrol edin. Dosya kapalı ise `"file: file is closed"` 
mesajı ile birlikte yeni bir hata döndürün.


```go
package main

type File struct {
	closed bool
}

func Read(file File) {}

func main() {
	f := File{closed: true}
	err := Read(f)
	fmt.Println(err)
}
```

2- `Queue` tipi için aşağıdaki kullanımı sağlayan metodları yazın.
```go
package main

type Queue struct {
	arr []int
}

// Append adds the given element at the end of the queue
func Append() {}

// AppendLeft adds the given element at the head of the queue
func AppendLeft() {}

// Pop removes the latest element in the queue
func Pop() {}

// Pop removes the first element in the queue
func PopLeft() {}

func main() {
	q := &Queue{}
	q.Append(4)     // q: [4]
	q.AppendLeft(1) // q: [1, 4]
	e := q.Pop()    // q: [1], e: 4
	e = q.PopLeft() // q: [0], e: 1
}
```


3- `subs` içerisinde bulunan herkese, `Notify` fonksiyonunu kullanarak, dead-lock oluşmadan, goroutineler ile mesaj gönderebilen kodu yazın.

```go
package main

import "fmt"

func Notify(to, message string) {
	fmt.Println("sending message:", message, "to:", to)
}

func main() {
	subs := []string{
		"user1",
		"user2",
		"user3",
	}
	msg := "hello"
	for _, sub := range subs {
		Notify(sub, msg)
	}
}
```

4- `select` yapısını kullanarak, `Hammer` fonksiyonunun 
3 saniyeden fazla sürmesi durumunda `main` fonksiyonunu bitirin.

```go
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

	res := <-ch
	fmt.Println("Finished after", res, "seconds.")
}
```

5- `Stringify` fonksiyonu çalıştırılınca oluşabilecek herhangi **2** davranışı test edin.

```go
package main

import (
	"errors"
	"fmt"
	"testing"
)

// Stringify turns the supported value types into a string
// with a specified format
func Stringify(value interface{}) (string, error) {
	switch value.(type) {
	case string:
		return value.(string), nil
	case float32, float64:
		return fmt.Sprintf("%.2f", value), nil
	case int:
		return fmt.Sprintf("%d", value), nil
	default:
		return "", errors.New("Invalid type")
	}
}

func TestStringify(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		t.Errorf("error")
	})
}
```