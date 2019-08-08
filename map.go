package main

import (
	"fmt"
)

func main() {

	mapping := make(map[string]string)
	mapping["HK"] = "CN"
	mapping["MO"] = "CN"
	mapping["XR"] = "SR"
	fmt.Println("map:", mapping)

	mapping2 := map[string]string{
		"HK": "CN", // Hong Kong(HK)
		"MO": "CN", // Macao(MO)
		"XR": "SR" /* Kosova(XR) */}
	fmt.Println("map2:", mapping2)

	v1 := mapping["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(mapping))

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
