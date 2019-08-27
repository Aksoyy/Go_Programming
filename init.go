package main

import "fmt"

var (
	message string = "first"
)

func main() {
	fmt.Println(message)
}

func init() {
	message = "priority"
}

// priority