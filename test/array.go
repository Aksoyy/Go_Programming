package main

import "fmt"

func main() {

	first_array := [2]int{5, 10}
	for index, value := range first_array {
		fmt.Printf("index:%d --> value:%d\n", index, value)
	}

}
