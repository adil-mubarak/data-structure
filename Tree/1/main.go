// insertion
package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// insert a new value into the BST
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

// inorder traverse the tree and print the numbers
func (node *TreeNode) Inorder() {
	if node == nil {
		return
	}
	node.left.Inorder()
	fmt.Printf("%d ", node.value)
	node.right.Inorder()
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(30)
	root.Insert(2)
	root.Insert(7)
	root.Insert(13)

	fmt.Println("In-order traverse:")
	root.Inorder()

}
