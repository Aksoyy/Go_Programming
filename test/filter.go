package main

import "fmt"

func main() {

	fruits := []string{"a", "b", "c", "d", "e", "b", "a", "b"}
	banana_criteria := "b"

	new_slice := []string{}
	for _, my_char := range fruits {
		if my_char == banana_criteria {
			new_slice = append(new_slice, my_char)
		}
	}

	fmt.Println("Size:", len(new_slice), "Value:", new_slice)
}

// Size: 3 Value: [b b b]
