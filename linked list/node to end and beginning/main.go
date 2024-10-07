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
	if dlist.tail == nil {
		dlist.head = newNode
		dlist.tail = newNode
	} else {
		newNode.prev = dlist.tail
		dlist.tail.next = newNode
		dlist.tail = newNode
	}
}

func (dlist *dlinkedlist) beginning(data int) {
	newNode := &dNode{data: data}
	if dlist.head == nil {
		dlist.head = newNode
		dlist.tail =newNode
	}else{
		newNode.next = dlist.head
		dlist.head.prev = newNode
		dlist.head = newNode
	}
}

func (dlist *dlinkedlist) dTraverse() {
	current := dlist.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
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
func (list *linkedlist) addtobeginning(data int) {
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}
func (list *linkedlist) Traverse() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
func main() {
	list := linkedlist{}
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Traverse()
	list.addtobeginning(1)
	list.Traverse()
	list.Insert(5)
	list.Traverse()
	list.addtobeginning(0)
	list.Traverse()
	// doubly linked list
	fmt.Println()
	fmt.Println("doubly linked list")
	fmt.Println()
	dlist := dlinkedlist{}
	dlist.dInsert(1)
	dlist.dInsert(2)
	dlist.dInsert(3)
	dlist.dTraverse()
	dlist.beginning(0)
	dlist.dTraverse()

}
