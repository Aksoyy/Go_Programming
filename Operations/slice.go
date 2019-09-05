package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 6, 8, 9, 2}

	for i := 3; i < 11; i++ {
		mySlice = append(mySlice, i)
	}

	fmt.Println(mySlice) // [1 3 6 8 9 2 3 4 5 6 7 8 9 10]
}
