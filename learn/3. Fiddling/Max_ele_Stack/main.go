package main

import "fmt"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Max() string {
	if s.IsEmpty() {
		return "There is nothing here"
	} else {
		max := (*s)[0]
		for _, v := range *s {
			if len(v) > len(max) {
				max = v
			}
		}
		return max
	}
}

func main() {
	var stack Stack

	stack.Push("This")
	stack.Push("is")
	stack.Push("Nepal")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println("Popped ELement:", x)
			fmt.Println("Maximum Value:", stack.Max())
		}
	}
}
