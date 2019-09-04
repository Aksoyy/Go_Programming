package main

import (
	"fmt"
)

func main() {
	records := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	batchSize := 3

	batchedRecords := func(data []int, batchSize int) [][]int {
		var batches [][]int

		for batchSize < len(data) {
			data, batches = data[batchSize:], append(batches, data[0:batchSize:batchSize])
		}
		batches = append(batches, data)
		return batches
	}(records, batchSize)

	result := func(veri int) string {
		var sonuc string
		for ; veri > 0; veri-- {
			sonuc += "hakan"
		}
		return sonuc
	}(3)

	fmt.Println(result)         // hakanhakanhakan
	fmt.Println(batchedRecords) // [[0 1 2] [3 4 5] [6 7 8] [9]]
}
