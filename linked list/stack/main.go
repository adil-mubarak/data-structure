package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) pop()(int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	topIndex := len(s.elements) - 1
	topElement := s.elements[topIndex]
	s.elements = s.elements[:topIndex]
	return topElement, nil
}

func (s *Stack) peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
func (s *Stack) IsFull()bool{
	return len(s.elements) != 0
}
func main() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack.IsFull())
	
	
	fmt.Println("stack elements: ")

	for _, elem := range stack.elements{
		fmt.Printf("%d ",elem)
	}
	fmt.Println()

	fmt.Println(stack.peek())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.IsFull())
}
