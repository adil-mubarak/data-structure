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

func (list *linkedlist) delete(value int) {
	if list.head == nil {
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.data != value{
		current = current.next
	}
	if current != nil{
		current.next = current.next.next
	}
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
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.Insert(50)
	list.Traverse()
	list.delete(30)
	list.Traverse()
}
