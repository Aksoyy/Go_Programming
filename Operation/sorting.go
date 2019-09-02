package main

import "fmt"
import "sort"

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// Sorted control
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

// Strings: [a b c]
// Ints:    [2 4 7]
// Sorted:  true
