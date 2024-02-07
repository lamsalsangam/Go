package main

import (
	"fmt"
	"math"
)

// Node structure
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Function to check if a binary tree is balanced
func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	// Get the heights of the left and right subtrees
	leftHeight := getHeight(root.left)
	rightHeight := getHeight(root.right)

	// Check if the difference between heights is within 1
	if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
		return false
	}

	// Recursively check if the subtrees are balanced
	return isBalanced(root.left) && isBalanced(root.right)
}

// Helper function to get the height of a subtree
func getHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.left), getHeight(root.right))
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	// Balanced tree example
	balancedRoot := &Node{
		data: 1,
		left: &Node{
			data: 2,
			left: &Node{
				data: 4,
			},
		},
		right: &Node{
			data: 3,
			right: &Node{
				data: 5,
			},
		},
	}

	// Unbalanced tree example
	unbalancedRoot := &Node{
		data: 1,
		left: &Node{
			data: 2,
			left: &Node{
				data: 4,
			},
		},
		right: &Node{
			data: 3,
			right: &Node{
				data: 6,
			},
		},
	}

	if isBalanced(balancedRoot) {
		fmt.Println("Balanced tree example: Balanced")
	} else {
		fmt.Println("Balanced tree example: Not balanced")
	}

	if isBalanced(unbalancedRoot) {
		fmt.Println("Unbalanced tree example: Balanced")
	} else {
		fmt.Println("Unbalanced tree example: Not balanced")
	}
}
