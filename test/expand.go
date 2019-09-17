package main

import "fmt"

func main() {

	bucket := []string{"a", "b", "c"}
	apples := []string{"d", "e"}
	expand_index := 2

	bucket = append(
		bucket[:expand_index],
		append(apples, bucket[expand_index:]...)...,
	)

	fmt.Println(bucket) // [a b d e c]
}
