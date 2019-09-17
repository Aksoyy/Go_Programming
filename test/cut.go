package main

import "fmt"

func main() {

	first, last := 3, 7
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	aralik := value[first-1 : last-1]
	fmt.Println(aralik) // [3 4 5 6]

	result := append(value[:first-1], value[last:]...)
	fmt.Println(result) // [1 2 8 9 10]
}
