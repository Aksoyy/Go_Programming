package main

import "fmt"

func main() {
	apple := "🍎"
	banana := "🍌"
	fruits := []*string{&apple, &apple, &apple, &banana, &banana, &banana}

	cut_start_index, cut_end_index := 4, 5

	fmt.Println(fruits[cut_start_index:], " :::: ", fruits[cut_end_index:])
	fmt.Println(fruits[cut_start_index:])
	fmt.Println(fruits[cut_end_index:])
	copy(fruits[cut_start_index:], fruits[cut_end_index:])
	fmt.Println(len(fruits), "    ", fruits)

	cleanup_index := len(fruits) - cut_end_index + cut_start_index // 6 - 4 - 3

	for free_index, free_end := cleanup_index, len(fruits); free_index < free_end; free_index++ {
		fruits[free_index] = nil
	}

	fruits = fruits[:cleanup_index]

	fmt.Println(fruits)
}

// [0x40c140 0x40c140]  ::::  [0x40c140]
// [0x40c140 0x40c140]
// [0x40c140]
// 6      [0x40c138 0x40c138 0x40c138 0x40c140 0x40c140 0x40c140]
// [0x40c138 0x40c138 0x40c138 0x40c140 0x40c140]
