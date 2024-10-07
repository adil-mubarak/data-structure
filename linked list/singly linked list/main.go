package main

import "fmt"

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
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func (list *linkedlist) delete(data int) {
	if list.head == nil {
		return
	}
	if list.head.data == data {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

func main() {
	list := linkedlist{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Traverse()
	list.delete(2)
	list.Traverse()
}
