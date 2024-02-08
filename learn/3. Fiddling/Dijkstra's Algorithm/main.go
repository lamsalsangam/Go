package main

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Graph [][]Edge

func dijkstra(graph Graph, start int) []int {
	numNodes := len(graph)
	distances := make([]int, numNodes)
	visited := make([]bool, numNodes)

	// Initialize distances with infinity and start node as 0
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	distances[start] = 0

	// Explore unvisited nodes
	for i := 0; i < numNodes-1; i++ {
		// Find the unvisited node with the minimum distance
		u := findMinDistanceNode(distances, visited)
		visited[u] = true

		// Update distances of adjacent nodes
		for _, edge := range graph[u] {
			v := edge.To
			if !visited[v] && distances[u] != math.MaxInt32 && distances[u]+edge.Weight < distances[v] {
				distances[v] = distances[u] + edge.Weight
			}
		}
	}

	return distances
}

// Finds the unvisited node with the minimum distance
func findMinDistanceNode(distances []int, visited []bool) int {
	minDistance := math.MaxInt32
	minNode := -1

	for i, distance := range distances {
		if !visited[i] && distance <= minDistance {
			minDistance = distance
			minNode = i
		}
	}

	return minNode
}

func main() {
	// ... (graph and start node setup)
	graph := Graph{
		[]Edge{{1, 10}, {2, 5}},
		[]Edge{{2, 2}, {3, 1}},
		[]Edge{{1, 3}, {3, 9}, {4, 2}},
		[]Edge{{4, 4}},
		[]Edge{},
	}

	start := 0

	distances := dijkstra(graph, start)
	fmt.Println("Shortest distances from node", start, "to all other nodes:")
	for i, d := range distances {
		fmt.Printf("Node %d: %d\n", i, d)
	}
}
