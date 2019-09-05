package main

import "fmt"

func main() {

	myMap := map[string]int{
		"hakan": 10,
		"nihat": 20,
		"cemil": 30,
	}

	for key, value := range myMap {
		fmt.Println("Anahtar:", key, " Deger:", value)
	}
}

// Anahtar: hakan  Deger: 10
// Anahtar: nihat  Deger: 20
// Anahtar: cemil  Deger: 30
