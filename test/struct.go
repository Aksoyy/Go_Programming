package main

import "fmt"

func main() {

	type test_data struct {
		name, lastname string
		age            int
	}

	my_struct := test_data{
		name:     "ali",
		lastname: "aksoy",
		age:      22,
	}

	fmt.Println("Struct:", my_struct) // Struct: {ali aksoy 22}
}
