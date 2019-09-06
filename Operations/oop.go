package main

import "fmt"

type MyStruct struct {
	name string
	age  int
	cost float64
}

func (a *MyStruct) Get(value string) bool {
	fmt.Println("test")
	return true
}

func main() {
	var mytest MyStruct
	fmt.Println(mytest.Get("ali"))
}

// test
// true
