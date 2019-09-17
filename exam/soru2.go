package main

import "fmt"

type Queue struct {
	arr []int
}

// Append adds the given element at the end of the queue
func (veri *Queue) Append(data int) {
	veri.arr = append(veri.arr, data)
}

// AppendLeft adds the given element at the head of the queue
func (veri *Queue) AppendLeft(data int) {
	veri.arr = append([]int{data}, veri.arr...)
}

// Pop removes the latest element in the queue
func (veri *Queue) Pop() (pop_item int) {
	pop_item, veri.arr = veri.arr[len(veri.arr)-1], veri.arr[:len(veri.arr)-1]
	return
}

// Pop removes the first element in the queue
func (veri *Queue) PopLeft() (pop_item int) {
	pop_item, veri.arr = veri.arr[0], veri.arr[1:]
	return
}

func main() {
	q := &Queue{}
	q.Append(4) // q: [4]
	fmt.Println(q)
	q.AppendLeft(1) // q: [1, 4]
	fmt.Println(q)
	e := q.Pop() // q: [1], e: 4
	fmt.Println(q, e)
	e = q.PopLeft() // q: [0], e: 1
	fmt.Println(q, e)
}

// &{[4]}
// &{[1 4]}
// &{[1]} 4
// &{[]} 1
