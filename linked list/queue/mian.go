package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	frontElement := q.elements[0]
	q.elements = q.elements[1:]
	return frontElement, nil
}

func (q *Queue) peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.elements[0], nil
}
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
func (q *Queue) IsFull() bool {
	return len(q.elements) != 0
}
func main() {
	queue := Queue{}
	fmt.Println(queue.IsFull())
	fmt.Println(queue.IsEmpty())
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("queue elements are: ")

	for _, elem := range queue.elements {
		fmt.Printf("%d \n", elem)
	}
	fmt.Println("")

	fmt.Println(queue.peek())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.peek())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.IsFull())
	fmt.Println(queue.IsEmpty())

}
