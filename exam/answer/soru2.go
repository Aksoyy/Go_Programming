// Queue tipi için aşağıdaki kullanımı sağlayan metodları yazın.

package main

import "fmt"

type Queue struct {
	arr []int
}

// Append adds the given element at the end of the queue
func (q *Queue) Append(e int) {
	q.arr = append(q.arr, e)
}

// AppendLeft adds the given element at the head of the queue
func (q *Queue) AppendLeft(e int) {
	q.arr = append([]int{e}, q.arr...)
}

// Pop removes the latest element in the queue
func (q *Queue) Pop() int {
	lastIndex := len(q.arr) - 1
	last := q.arr[lastIndex]
	q.arr[len(q.arr)-1] = 0

	q.arr = q.arr[:lastIndex]
	return last
}

// Pop removes the first element in the queue
func (q *Queue) PopLeft() int {
	first := q.arr[0]
	q.arr[0] = 0
	q.arr = q.arr[1:]
	return first
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
