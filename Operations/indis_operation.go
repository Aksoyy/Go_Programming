package main

import (
	"fmt"
)

func main() {
	//fmt.Println("test")
	degisken_ismi := []int{10, 20, 30, 90}
	degisken_ismi[2] = 1
	fmt.Println(degisken_ismi)
	var d1 [2]int
	d1[1] = 11
	d1[0] = 10

	deg2 := append(degisken_ismi, 4, 7)
	fmt.Println(deg2)

	n1 := deg2[:1]
	n2 := deg2[2:]
	n3 := deg2[1:7]
	fmt.Println(n1, n2, n3)
}

// [10 20 1 90]
// [10 20 1 90 4 7]
// [10] [1 90 4 7] [20 1 90 4 7 0]
