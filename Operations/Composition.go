package main

import "fmt"

type Hesap struct{}

func (h *Hesap) HesapBul() float32 {
	fmt.Println("Test")
	return 0
}

func (h *Hesap) HesapOde(miktar float32) bool {
	fmt.Println("Test2")
	return true
}

type KrediHesap struct {
	Hesap
}

func main() {
	test := &KrediHesap{}
	test.HesapBul()    // Test
	test.HesapOde(400) // Test2
}
