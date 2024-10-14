package main

import "fmt"

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

func (node *TreeNode) Height() int {
	if node == nil {
		return 0
	}

	lefHeight := node.left.Height()
	rightHeigt := node.right.Height()

	if lefHeight > rightHeigt {
		return lefHeight + 1
	} else {
		return rightHeigt + 1
	}
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(20)
	root.Insert(3)
	root.Insert(7)
	root.Insert(15)

	fmt.Println("Checking the height:", root.Height())
}
