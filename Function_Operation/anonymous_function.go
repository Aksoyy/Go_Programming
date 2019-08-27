package main

import "fmt"

func main() {

	myFunc := func(items ...int) (p1 int, p2 int) {
		for _, item := range items {
			p2 += item
		}
		p1 = len(items)
		return
	}

	a1, a2 := myFunc(1, 2, 3, 4, 5)
	fmt.Println("Miktar:", a1, " Sonuc:", a2)
}

// Miktar:5 Sonuc:15
// Miktar: 5  Sonuc: 15
