package main

import "fmt"

func main() {

	apple, banana := "a", "b"
	fruits := []*string{&apple, &apple, &apple, &banana, &banana, &banana}

	cut_start_index, cut_end_index := 1, 4
	copy(fruits[cut_start_index:], fruits[cut_end_index:])

	for index, value := range fruits {
		fmt.Printf("%d. deger --> %s\n", index, *value)
	}
	fmt.Println("Size:", len(fruits))
	//					6				3				1
	new_slice_size := len(fruits) - cut_end_index + cut_start_index

	//		4			6
	for free_index, free_end := new_slice_size, len(fruits); free_index < free_end; free_index++ {
		fruits[free_index] = nil
	}
	fruits = fruits[:new_slice_size]

	for index, value := range fruits {
		fmt.Printf("%d. deger --> %s\n", index, *value)
	}
	fmt.Println("Size:", len(fruits))

}

// 0. deger --> a
// 1. deger --> b
// 2. deger --> b
// 3. deger --> b
// 4. deger --> b
// 5. deger --> b
// Size: 6
// 0. deger --> a
// 1. deger --> b
// 2. deger --> b
// Size: 3
