package main

import "fmt"

func main() {

	first_slice := []string{"deneme", "deneme2", "deneme3"}
	for index, value := range first_slice {
		fmt.Printf("%d. deger: %s\n", index, value)
	}

	second_slice := first_slice[:2]
	fmt.Println(second_slice)

	result := 0
	for _, value := range first_slice {
		for _, value2 := range second_slice {
			if value == value2 {
				result++
			}
		}
	}
	fmt.Println("benzerlik sayisi:", result)
}

// 0. deger: deneme
// 1. deger: deneme2
// 2. deger: deneme3
// [deneme deneme2]
// benzerlik sayisi: 2
