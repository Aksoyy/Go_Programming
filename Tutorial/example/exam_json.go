package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
)

func main() {
	m := hashmap.New()
	m.Put("a", "1")
	m.Put("b", "2")
	m.Put("c", "3")

	json, err := m.ToJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json)) // {"a":"1","b":"2","c":"3"}