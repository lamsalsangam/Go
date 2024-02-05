package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	nodes []*Node
}

// Node structure
type Node struct {
	value   int
	edges   []*Node
	visited bool
}

// Function to create a new graph
func NewGraph() *Graph {
	return &Graph{}
}

// Function to add a node to the graph
func (g *Graph) AddNode(value int) {
	node := &Node{value: value}
	g.nodes = append(g.nodes, node)
}

// Function to add an edge between two nodes
func (g *Graph) AddEdge(from, to int) {
	fromNode := g.getNode(from)
	toNode := g.getNode(to)
	fromNode.edges = append(fromNode.edges, toNode)
}

// Helper function to get a node by its value
func (g *Graph) getNode(value int) *Node {
	for _, node := range g.nodes {
		if node.value == value {
			return node
		}
	}
	return nil
}

// Breadth-first search (BFS) function using a slice as a queue
func (g *Graph) BFS(startNode *Node) {
	queue := []*Node{}
	queue = append(queue, startNode)
	startNode.visited = true

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		fmt.Println(currentNode.value)

		for _, neighbor := range currentNode.edges {
			if !neighbor.visited {
				queue = append(queue, neighbor)
				neighbor.visited = true
			}
		}
	}
}

func main() {
	graph := NewGraph()
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)

	graph.BFS(graph.nodes[0]) // Start BFS from node 1
}
