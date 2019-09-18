package main

import "fmt"

func main() {

	var my_slice []int
	var my_slice2 = make([]int, 4, 4)
	fmt.Println(my_slice, my_slice2) // [] [0 0 0 0]

	fmt.Println(len(my_slice)) // 0
	my_slice = append(my_slice, 11, 22, 33, 44, 55)
	fmt.Println("size:", len(my_slice), "value:", my_slice) // 	size: 5 value: [11 22 33 44 55]

	my_slice = append(my_slice[:3], my_slice[2:]...)
	fmt.Println(my_slice) // [11 22 33 33 44 55]

	fmt.Println("size:", len(my_slice), " capacity:", cap(my_slice)) // size: 6  capacity: 6
}
