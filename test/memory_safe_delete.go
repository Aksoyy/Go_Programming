package main

import "fmt"

func main() {

	apple, banana := "a", "b"
	fruits := []*string{&apple, &apple, &apple, &banana, &banana, &banana}
	delete_index := 2

	size := copy(fruits[delete_index:], fruits[delete_index+1:])
	for i, v := range fruits {
		fmt.Println("index:", i, " value:", *v)
	}
	fmt.Println(size)

	fruits[len(fruits)-1] = nil
	fruits = fruits[:len(fruits)-1]
	for i, v := range fruits {
		fmt.Println("index:", i, " value:", *v)
	}

}

// index: 0  value: a
// index: 1  value: a
// index: 2  value: b
// index: 3  value: b
// index: 4  value: b
// index: 5  value: b
// 3
// index: 0  value: a
// index: 1  value: a
// index: 2  value: b
// index: 3  value: b
// index: 4  value: b
