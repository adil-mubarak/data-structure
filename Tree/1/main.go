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

func (node *TreeNode) InOrder() {
	if node == nil {
		return
	}
	node.left.InOrder()
	fmt.Printf("%d ", node.value)
	node.right.InOrder()
}

func (node *TreeNode) PreOrder() {
	if node == nil{
		return
	}
	fmt.Printf("%d ", node.value)
	node.left.PreOrder()
	node.right.PreOrder()
}

func (node *TreeNode) PostOrder() {
	if node == nil{
		return
	}
	node.left.PostOrder()
	node.right.PostOrder()
	fmt.Printf("%d ", node.value)
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(20)
	root.Insert(3)
	root.Insert(7)
	root.Insert(15)

	fmt.Println("In-order")
	root.InOrder()
	fmt.Println("\nPre-order")
	root.PreOrder()
	fmt.Println("\nPost-order")
	root.PostOrder()
}
