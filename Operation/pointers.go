package main

import "fmt"

func main() {

	var mess string = "TEST"
	var mess2 *string = &mess
	*mess2 = "TEST2"

	fmt.Printf("Last value: %s", mess)

}

// Last value: TEST2%
