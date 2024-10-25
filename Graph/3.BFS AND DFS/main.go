package main

import "fmt"

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[string][]string)}
}

func (g *Graph) AddEdge(source, destination string) {
	g.adjacencyList[source] = append(g.adjacencyList[source], destination)
	g.adjacencyList[destination] = append(g.adjacencyList[destination], source)
}

func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node + " ")

		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func (g *Graph) DFS(start string) {
	visited := make(map[string]bool)
	g.dfsHelper(start, visited)
}
func (g *Graph) dfsHelper(node string, visited map[string]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	fmt.Print(node + " ")
	for _, neighbor := range g.adjacencyList[node] {
		g.dfsHelper(neighbor, visited)
	}
}

func main() {
	graph := NewGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("D", "F")

	fmt.Println("BFS Traversal starting from node A:")
	graph.BFS("A")
	fmt.Println()

	fmt.Println("DFS Traversal starting from node A:")
	graph.DFS("A")
}
