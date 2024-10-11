//search

package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Insert values
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

// Search the values
func (node *TreeNode) Search(value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	} else if value < node.value {
		return node.left.Search(value)
	} else {
		return node.right.Search(value)
	}
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(20)
	root.Insert(3)
	root.Insert(7)
	root.Insert(15)

	fmt.Println("Searching for 7:", root.Search(7))
	fmt.Println("Searching for 12:", root.Search(12))
}
