package main

import "fmt"

func main() {
	apple := "🍎"
	banana := "🍌"
	fruits := []*string{&apple, &apple, &apple, &banana, &banana, &banana}
	delete_index := 2

	copy(fruits[delete_index:], fruits[delete_index+1:])
	fmt.Println(len(fruits)) // 6
	fruits[len(fruits)-1] = nil
	fmt.Println(len(fruits)) // 6
	fruits = fruits[:len(fruits)-1]
	fmt.Println(len(fruits)) // 5

	fmt.Println(fruits)                // [0xc0000101e0 0xc0000101e0 0xc0000101f0 0xc0000101f0 0xc0000101f0]
	fmt.Println(fruits[len(fruits)-1]) // 0xc0000101f0
}
