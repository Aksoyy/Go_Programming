package main

import "fmt"

func main() {

	result := karsilastir(3, 5, 7, 4, 1)
	fmt.Println("result:", result) // result: 7

	sum := topla(3, 5, 7, 9, 1)
	fmt.Println("sum:", sum) // sum: 25

}

func topla(data ...int) (ali int) {
	ali = 0
	for _, oku := range data {
		ali += oku
	}
	return
}

func karsilastir(data ...int) int {

	control := data[0]
	for indis, deger := range data {
		//fmt.Println(indis)
		if indis != 0 {
			//fmt.Println(indis)
			if control < deger {
				control = deger
			}
		}
	}
	return control
}
