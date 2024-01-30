package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func findKthToLast(ll *LinkedList, k int) (*Node, bool) {
	if k <= 0 || ll.head == nil {
		return nil, false
	}

	p1 := ll.head
	p2 := ll.head

	// Move p2 k nodes ahead
	for i := 0; i < k; i++ {
		if p2 == nil {
			return nil, false // k is greater than the length of the list
		}
		p2 = p2.next
	}

	// Move both pointers until p2 reaches the end
	for p2 != nil {
		p1 = p1.next
		p2 = p2.next
	}

	return p1, true
}

func main() {
	ll := &LinkedList{}

	ll.insert(1)
	ll.insert(2)
	ll.insert(3)
	ll.insert(4)
	ll.insert(5)

	ll.display()

	k := 2
	kthToLast, found := findKthToLast(ll, k)

	if found {
		fmt.Printf("The %dth to last element is: %d\n", k, kthToLast.data)
	} else {
		fmt.Printf("Invalid value of k or k is greater than the length of the list\n")
	}
}
