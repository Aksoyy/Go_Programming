package main

import "fmt"

type myStruct struct {
	name      string
	last_name string
}

func main() {

	test_struct := myStruct{}
	test_struct.name = "Hakan"
	test_struct.last_name = "Aksoy"

	fmt.Println("Result:", test_struct)
}

// Result: {Hakan Aksoy}
