// basic operation in bst
package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// insert the value
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

// check values contains
func (node *TreeNode) Contains(value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	} else if value < node.value {
		return node.left.Contains(value)
	} else {
		return node.right.Contains(value)
	}
}

// finding the minimum value
func (node *TreeNode) FindMin() int {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.value
}

// remove or delete a value from the tree
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

// InOrder traversal(left,root,right)
func (node *TreeNode) InOrder() {
	if node == nil {
		return
	}
	node.left.InOrder()
	fmt.Printf("%d ", node.value)
	node.right.InOrder()

}

// PreOrder traversal(root,left,right)
func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.value)
	node.left.PreOrder()
	node.right.PreOrder()
}

// PostOrder traversal(left,right,root)
func (node *TreeNode) PostOrder() {
	if node == nil {
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

	fmt.Println("In-order traversal:")
	root.InOrder()

	fmt.Println("\nPre-order traversal")
	root.PreOrder()

	fmt.Println("\nPost-order traversal")
	root.PostOrder()

	fmt.Println("\nContains 7:", root.Contains(7))
	fmt.Println("\nContains 12:", root.Contains(12))

	fmt.Println("\nDeleting 5..")
	root.Delete(5)
	fmt.Println("In-order traversal after deletion:")
	root.InOrder()
}
