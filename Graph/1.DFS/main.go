package main

import "fmt"

func riverSizes(matrix [][]int)[]int{
	visited := make([][]bool,len(matrix))
	for i := range visited{
		visited[i] = make([]bool, len(matrix[0]))
	}

	var sizes []int
	for i := 0; i < len(matrix); i++{
		for j := 0; j < len(matrix[0]);j++{
			if visited[i][j] || matrix[i][j] == 0{
				continue
			}
			size := traverseRiver(i,j,matrix,visited)
			sizes = append(sizes, size)
		}
	}
	return sizes
}

func traverseRiver(i,j int,matrix [][]int,visited [][]bool)int{
	curretSize := 0
	nodesToExplore := [][]int{{i,j}}

	for len(nodesToExplore) > 0{
		currentNode := nodesToExplore[len(nodesToExplore)-1]
		nodesToExplore = nodesToExplore[:len(nodesToExplore)-1]
		i,j := currentNode[0],currentNode[1]

		if visited[i][j]{
			continue
		}
		visited[i][j] = true

		if matrix[i][j] == 0{
			continue
		}
		curretSize++

		univisitedNeighbors := getUniversityNeighbors(i,j,matrix,visited)
		nodesToExplore = append(nodesToExplore, univisitedNeighbors...)
	}
	return curretSize
}

func getUniversityNeighbors(i,j int,matrix [][]int,visited [][]bool)[][]int{
	var neighbors [][]int

	if i > 0 && !visited[i-1][j]{
		neighbors = append(neighbors, []int{i-1,j})
	}
	if i < len(matrix[0])-1 && !visited[i+1][j]{
		neighbors = append(neighbors, []int{i+1,j})
	}
	if j >0 && !visited[i][j-1] {
		neighbors = append(neighbors, []int{i,j-1})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1]{
		neighbors = append(neighbors, []int{i,j+1})
	}
	return neighbors
}

func main() {
	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	fmt.Println(riverSizes(matrix))
}