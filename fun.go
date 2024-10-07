package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type linkedlist struct{
	head *Node
}

func (list *linkedlist)Insert(data int){
	newNode := &Node{data:data}
	if list.head == nil{
		list.head = newNode
	}else{
		current := list.head
		for current.next != nil{
			current = current.next
		}
			current.next = newNode
	}
}

func (list *linkedlist)Traverse(){
	current := list.head
	for current != nil{
		fmt.Print(current.data," -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func (list *linkedlist)Remove(){
	current := list.head
	for current != nil && current.next != nil{
		if current.data == current.next.data{
			current.next = current.next.next
		}else{
			current = current.next
		}
	}

}
func main(){
list := linkedlist{}
list.Insert(1)
list.Insert(2)
list.Insert(2)
list.Insert(3)
list.Traverse()
list.Remove()
list.Traverse()

}