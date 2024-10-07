package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedlist struct {
	head *Node
}

func arraytolinkedlist(arr []int) *linkedlist {
	list := &linkedlist{}
	for _, value := range arr {
		list.Insert(value)
	}
	return list
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
		fmt.Printf(" %d ->", current.data)
		current = current.next
	}
	fmt.Println(" nil")
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := arraytolinkedlist(arr)
	list.Traverse()
}
