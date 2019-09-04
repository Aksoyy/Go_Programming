package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	fruits := [3]int{1, 2}
	fmt.Println(fruits)

	strinnng := [3]string{"a", "b", "c"}

	for _, fruit := range strinnng {
		fmt.Print(fruit)
	}
	fmt.Println()

	mySlice := []int{4, 5, 6, 6}
	fmt.Println(mySlice)
	addValue := 10
	mySlice = append(mySlice, addValue)
	fmt.Println(mySlice)

	fruitss := []string{"a", "c", "e", "g", "k"}
	apples := fruitss[:2]
	apples[0] = "p"
	apples[1] = "p"
	fmt.Println(apples)
	fmt.Println(fruitss)

	edibles := []string{"a", "b", "s", "d", "e", "r", "t", "y"}
	cut_start_index, cut_end_index := 4, 6

	fmt.Println(edibles[:cut_start_index], "--> belirtilen indisi kadar")
	fmt.Println(edibles[cut_end_index:], "--> Kesilmenin bitecegi indis")
	edibles = append(edibles[:cut_start_index], edibles[cut_end_index:]...)
	fmt.Println(edibles)

	deleteİndex := 2
	edibles = append(edibles[:deleteİndex], edibles[deleteİndex+1:]...)
	fmt.Println(edibles)

}
