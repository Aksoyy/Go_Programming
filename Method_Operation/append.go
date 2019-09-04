package main

import "fmt"

type Iterable []int

func (iterable *Iterable) Append(items ...int) {
	*iterable = append(*iterable, items...)
}

func main() {
	arr := Iterable{1, 2}
	arr.Append(4, 5, 6)
	fmt.Println(arr) // [1 2 4 5 6]
}
