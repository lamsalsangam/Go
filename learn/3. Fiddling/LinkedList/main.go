package main

import "fmt"

// A linked list consists of nodes, where each node stores a piece of data and a reference to the next node in the sequence.
type Node struct {
	data int
	next *Node
}

// The LinkedList struct represents the linked list and has a reference to the head node.
type LinkedList struct {
	head *Node
}

// Insert adds a new node with the given data to the beginning of the linked list
func (ll *LinkedList) insert(data int) {
	new := &Node{data: data, next: ll.head}
	ll.head = new
}

func (ll *LinkedList) delete(key int) {
	current := ll.head

	// If the key is present at the head
	if current != nil && current.data == key {
		ll.head = current.next
		return
	}

	var prev *Node
	for current != nil && current.data != key {
		prev = current
		current = current.next
	}

	// If the key is not present
	if current == nil {
		return
	}

	// Unlink the node
	prev.next = current.next
}

// Search checks if the given key exists in the linked list
func (ll *LinkedList) search(key int) bool {
	current := ll.head
	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}
	return false
}

// Display prints the elements of the linked list
func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	linkedList := &LinkedList{}
	linkedList.insert(3)
	linkedList.insert(7)
	linkedList.insert(1)
	linkedList.insert(9)
	linkedList.insert(5)

	fmt.Println("Original Linked List:")
	linkedList.display()

	keyToSearch := 7
	fmt.Printf("Is %d present in the linked list? %t\n", keyToSearch, linkedList.search(keyToSearch))

	keyToDelete := 1
	linkedList.delete(keyToDelete)
	fmt.Printf("Linked List after deleting %d:\n", keyToDelete)
	linkedList.display()
}
