package main

import "fmt"

func rec(degisken int) int {
	if degisken == 0 || degisken == 1 {
		return 1
	}
	return degisken * rec(degisken-1)
}

func main() {
	fmt.Println(rec(5))
}

// 120
