package main

import "fmt"

func Cikti(value, value2 string) (veri string) {
	veri = value + value2
	return
}

func main() {
	fmt.Println(Cikti("Hakan", "Aksoy")) // Hakan Aksoy
}
