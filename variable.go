package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// Const exam
	fmt.Println(s)
	const n = 500000000

	fmt.Println(math.Sin(n))

}

/*
initial
1 2
true
0
apple
constant
-0.28470407323754404
*/
