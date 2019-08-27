package main

import "fmt"

func main() {
	result := topla(1, 5, 8987, 3)
	fmt.Printf("Sayilarin toplami: %d", result)
}

func topla(gelenVeri ...int) int {
	temp := 0
	for _, parca := range gelenVeri {
		temp += parca
	}
	return temp
}

// Sayilarin toplami: 8996%
