package main

import "github.com/emirpasic/gods/sets/hashset"

func main() {
	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)       // 1, 5 (random order)
	set.Contains(1)        // true
	set.Contains(1, 5)     // true
	set.Contains(1, 6)     // false
	_ = set.Values()       // []int{5,1} (random order)
	set.Clear()            // empty
	set.Empty()            // true
	set.Size()             // 0
}
