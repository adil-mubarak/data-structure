package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type linkedlist struct {
	head *Node
	tail *Node
}

func (list *linkedlist) Insert(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}
func (list *linkedlist) Traverse() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
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
		if list.head != nil {
			list.head.prev = nil
		} else {
			list.tail = nil
		}
		return
	}
	current := list.head
	for current != nil && current.data != data {
		current = current.next
	}
	if current != nil {
		if current.next != nil {
			current.next.prev = current.prev
		} else {
			list.tail = current.prev
		}
		if current.prev != nil{
			current.prev.next = current.next
		}
	}
}
func main() {
	list := linkedlist{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Traverse()
	list.delete(2)
	list.Traverse()
}
