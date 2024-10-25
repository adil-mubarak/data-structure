package main

import "fmt"

func hasCycle(graph map[int][]int) bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for node := range graph {
		if !visited[node] && dfsCycle(graph, node, visited, recStack) {
			return true
		}
	}
	return false
}

func dfsCycle(graph map[int][]int, node int, visited, recStack map[int]bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] && dfsCycle(graph, neighbor, visited, recStack) {
			return true
		}
		if recStack[neighbor] {
			return true
		}
	}

	recStack[node] = false
	return false
}

func main() {
	graph := map[int][]int{
		0: {1},
		1: {2},
		2: {0},
		3: {4},
	}

	if hasCycle(graph) {
		fmt.Println("Graph has a cycle.")
	}else{
		fmt.Println("Graph doesn't have a cycle")
	}
}