package main

import "fmt"

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(str string) {
	*q = append(*q, str)
}

func (q *Queue) Pop() (string, bool) {
	if q.IsEmpty() {
		return "", false
	} else {
		index := len(*q) - 1
		element := (*q)[0]
		*q = (*q)[(index + 1):]
		return element, true
	}
}

func main() {
	var queue Queue

	queue.Push("This")
	queue.Push("is")
	queue.Push("Nepal")

	for len(queue) > 0 {
		x, y := queue.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
