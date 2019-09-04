package main

import (
	"fmt"
)

func ReLU(nums ...float64) []float64 {
	new_nums := make([]float64, len(nums))
	for idx, value := range nums {
		if value > 0 {
			new_nums[idx] = value
		} else {
			new_nums[idx] = 0.
		}
	}
	return new_nums
}

func main() {
	nums := []float64{1., 0.2, 0., 0., -0.1, 0.1}

	nums = ReLU(nums...)

	fmt.Println(nums) // [1 0.2 0 0 0 0.1]
}
