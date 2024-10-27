// height calculation
package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// insert values
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

// height of tree
func (node *TreeNode) height() int {
	if node == nil {
		return 0
	}
	leftHeight := node.left.height()
	rightHeight := node.right.height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(4)
	root.Insert(2)
	root.Insert(1)
	root.Insert(7)
	root.Insert(15)
	root.Insert(3)
	root.Insert(20)
	root.Insert(23)

	fmt.Println("Height of the tree:", root.height())
}
