package main

import (
	"fmt"

	"github.com/emirpasic/gods/maps/hashmap"
)

func main() {
	hm := hashmap.New()

	json := []byte(`{"a":"1","b":"2"}`)
	err := hm.FromJSON(json)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hm) // HashMap map[b:2 a:1]
}
