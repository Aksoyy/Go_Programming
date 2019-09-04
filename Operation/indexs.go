package main

import "fmt"

func main() {
	apple := "🍎"
	banana := "🍌"
	fruits := []*string{&apple, &apple, &apple, &banana, &banana, &banana}

	cut_start_index, cut_end_index := 1, 5

	fmt.Println(fruits[cut_start_index:], " :::: ", fruits[cut_end_index:])
	copy(fruits[cut_start_index:], fruits[cut_end_index:])

	cleanup_index := len(fruits) - cut_end_index + cut_start_index

	for free_index, free_end := cleanup_index, len(fruits); free_index < free_end; free_index++ {
		fruits[free_index] = nil
	}

	fruits = fruits[:cleanup_index]

	fmt.Println(fruits)
}

// [0xc000084030 0xc000084030 0xc000084040 0xc000084040 0xc000084040]  ::::  [0xc000084040]
// [0xc000084030 0xc000084040]
