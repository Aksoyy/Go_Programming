package main

import "fmt"

func main() {
	//	 				 0	  1	   2	3	 4	  5	   6	7
	my_slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	delete_index := 2 // delete index

	copied := copy(my_slice[delete_index:], my_slice[delete_index+1:])
	fmt.Println(copied, my_slice) // 5 [a b d e f g h h]
	new_length := delete_index + copied

	my_slice = my_slice[:new_length]
	fmt.Println(my_slice) // [a b d e f g h]
}
