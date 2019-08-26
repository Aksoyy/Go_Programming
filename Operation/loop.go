package main

import (
	"fmt"
)

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1 // i+=1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// unconditional loop
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	k := 5
	for ; k < 7; k++ {
		fmt.Println("deneme")
	}
}
