package main

import "fmt"

func main() {

	fruits := []string{"a", "b", "c"}
	banana := "ADD"
	insert_index := 1

	fruits = append(
		fruits[:insert_index],
		append([]string{banana}, fruits[insert_index:]...)...,
	)

	fmt.Println("Size:", len(fruits), "Value:", fruits)
}

// Size: 4 Value: [a ADD b c]
