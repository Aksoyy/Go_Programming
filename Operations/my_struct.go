package main

import "fmt"

func main() {

	type myStruct struct {
		name     string
		lastname string
		age      int
		cost     float64
	}

	newStruct := myStruct{
		name:     "hakan",
		lastname: "aksoy",
		age:      10,
		cost:     1000.0,
	}

	fmt.Println(newStruct.age, newStruct.name) // 10 hakan
	fmt.Println(newStruct)                     // {hakan aksoy 10 1000}
}
