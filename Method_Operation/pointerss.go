package main

import (
	"fmt"
)

func main() {

	var carbon string = "Carbon"
	var ptrCarbon *string

	ptrCarbon = &carbon

	*ptrCarbon = "Altered " + *ptrCarbon

	fmt.Println(ptrCarbon)  // 0xc000010050 (Her kosumda degısebılır)
	fmt.Println(*ptrCarbon) // Altered Carbon

}
