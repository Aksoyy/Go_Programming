package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Dict example\n")
	// dict_ := map[string]string{"deneme": "deneme2"}
	// fmt.Println(dict_)

	godur, _ := time.ParseDuration("10ms")

	go func() {
		for index := 1; index < 10; index += 1 {
			fmt.Println("Ilk cikti")
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 1; i < 10; i += 1 {
			fmt.Println("Ikinci cikti")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

// Ikinci cikti
// Ilk cikti
// Ikinci cikti
// Ilk cikti
// Ikinci cikti
// Ilk cikti
// Ilk cikti
// Ikinci cikti
// Ilk cikti
// Ikinci cikti
// Ikinci cikti
// Ilk cikti
// Ilk cikti
// Ikinci cikti
// Ikinci cikti
// Ilk cikti
// Ikinci cikti
// Ilk cikti
