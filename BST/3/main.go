// validate that given is a bst
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

// checking that tree is bst
func IsValidBST(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.value <= *min {
		return false
	}
	if max != nil && node.value >= *max {
		return false
	}
	return IsValidBST(node.left, min, &node.value) && IsValidBST(node.right, &node.value, max)
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(20)
	root.Insert(3)
	root.Insert(7)
	root.Insert(15)

	fmt.Println("Is valid BST", IsValidBST(root, nil, nil))

}
