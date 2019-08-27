package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	my_struct := person{"Bob", 20}
	fmt.Println(my_struct.age, my_struct.name)
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(s.age)
}

// 20 Bob
// {Alice 30}
// {Fred 0}
// &{Ann 40}
// Sean
// 50
// 51
