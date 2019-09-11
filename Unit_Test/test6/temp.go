package main

import "fmt"

func Cikti(value ...string) (veri string) {
	for _, data := range value {
		veri += data
	}
	return
}

func main() {
	fmt.Println(Cikti("Hakan", "Aksoy")) // Hakan Aksoy
}
