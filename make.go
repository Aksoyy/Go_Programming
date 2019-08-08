package main

import "fmt"

func main() {

	s := make([]string, 7)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// test := make([][]int, 3, 3)
	var test [3][3]int

	for a := 0; a < 3; a++ {

		for b := 0; b < 3; b++ {
			test[a][b] = (a + 1) * (b + 1)
			// fmt.Println(test[a][b])
		}
		// test[a][a] = a * a
	}
	test[1][2] = 5
	fmt.Println(test)

}

// emp: [      ]
// set: [a b c    ]
// get: c
// len: 7
// apd: [a b c     d e f]
// cpy: [a b c     d e f]
// sl1: [c  ]
// sl2: [a b c  ]
// sl3: [c     d e f]
// dcl: [g h i]
// 2d:  [[0] [1 2] [2 3 4]]
// [[1 2 3] [2 4 5] [3 6 9]]
