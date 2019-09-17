package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)  // Type  main.order
	fmt.Println("Value ", v) // Value  {456 56}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)
}
