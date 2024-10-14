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

// find minimum value
func (node *TreeNode) FindMin() int {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.value
}

// remove value  or delete value
func (node *TreeNode) Delete(value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = node.left.Delete(value)
	} else if value > node.value {
		node.right = node.right.Delete(value)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			node.value = node.right.FindMin()
			node.right = node.right.Delete(node.value)
		}
	}
	return node
}

func (node *TreeNode) InOrder() {
	if node == nil {
		return
	}
	node.left.InOrder()
	fmt.Printf("%d ",node.value)
	node.right.InOrder()
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(3)
	root.Insert(20)
	root.Insert(5)
	root.Insert(7)
	root.Insert(15)

	root.InOrder()
	fmt.Println("find min value:",root.FindMin())
	root.Delete(5)
	root.InOrder()
}