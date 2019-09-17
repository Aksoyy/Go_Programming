package main

import "fmt"

func main() {

	my_map := map[string]int{
		"deneme":  1,
		"deneme2": 2,
		"deneme3": 3,
	}

	value, control := my_map["deneme"]
	if !control {
		fmt.Println("not found")
	} else {
		fmt.Println(value) // 1
	}

	value2, control2 := my_map["deneme4"]
	if !control2 {
		fmt.Println("not found")
	} else {
		fmt.Println(value2)
	}

}

// 1
// not found
