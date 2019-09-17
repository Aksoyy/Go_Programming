package main

import "fmt"

func main() {
	//	 				 0	  1	   2	3	 4	  5	   6	7
	edibles := []string{"a", "b", "c", "d", "e", "d", "f", "g"}
	index_of_cherry := 2

	n_copied := copy(edibles[index_of_cherry:], edibles[index_of_cherry+1:])
	new_length := index_of_cherry + n_copied

	edibles = edibles[:new_length]
	fmt.Println(edibles) // [a b d e d f g]
}
