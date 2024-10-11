// find the closet value
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) Insert(value int) {
	if value < node.value {
		if node.left == nil {
			node.left = &TreeNode{value: value}
		} else {
			node.left.Insert(value)
		}
	} else {
		if node.right == nil {
			node.right = &TreeNode{value: value}
		} else {
			node.right.Insert(value)
		}
	}
}

// find the closet value
func (node *TreeNode) FindClosest(target int) int {
	closest := node.value
	currentNode := node
	for currentNode != nil {
		if math.Abs(float64(target-currentNode.value)) < math.Abs(float64(target-closest)) {
			closest = currentNode.value
		}

		if target < currentNode.value {
			currentNode = currentNode.left
		} else if target > currentNode.value {
			currentNode = currentNode.right
		} else {
			break
		}
	}
	return closest
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(20)
	root.Insert(3)
	root.Insert(7)
	root.Insert(15)

	target := 12
	fmt.Printf("Closest value to %d: %d\n", target, root.FindClosest(target))
}
