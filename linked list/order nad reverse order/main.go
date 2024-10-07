package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type linkedlist struct {
	head *Node
}

func (list *linkedlist) Insert(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *linkedlist) Traverse() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("null")
}

func printReverse(node *Node) {
	if node == nil {
		return
	}
	printReverse(node.next)
	fmt.Printf("%d -> ", node.data)
}

func (list *linkedlist) Reverse() {
	printReverse(list.head)
	fmt.Println("null")
}

type dNode struct {
	data int
	next *dNode
	prev *dNode
}

type dlinkedlist struct {
	head *dNode
	tail *dNode
}

func (dlist *dlinkedlist) dInsert(data int) {
	newNode := &dNode{data: data}
	if dlist.head == nil {
		dlist.head = newNode
		dlist.tail = newNode
	} else {
		newNode.prev = dlist.tail
		dlist.tail.next = newNode
		dlist.tail = newNode
	}
}

func (dlist *dlinkedlist) dTraverse() {
	current := dlist.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("null")
}

func (dlist *dlinkedlist) dReverse() {
	current := dlist.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("null")
}
func main() {
	list := linkedlist{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Traverse()
	list.Reverse()
	fmt.Println("")
	fmt.Println("doubly linked list")
	fmt.Println("")
	dlist := dlinkedlist{}
	dlist.dInsert(10)
	dlist.dInsert(20)
	dlist.dInsert(30)
	dlist.dTraverse()
	dlist.dReverse()
	
}
