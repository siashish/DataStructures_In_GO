package main

import (
	"errors"
	"fmt"
)

// Queue := := Queue implementation backed by Array/Slice
type Queue struct {
	data []int
}

// NewQueue :=  Creates a new Queue with the provided capacity
func NewQueue(capacity int) *Queue {
	queue := new(Queue)
	queue.data = make([]int, capacity)
	return queue
}

// Enqueue := inserts the element at the front of the queue
func (queue *Queue) Enqueue(data int) {
	queue.data = append(queue.data, data)
}

// Dequeue := removes the element from the back of the queue
func (queue *Queue) Dequeue() (int, error) {
	if queue.Size() == 0 {
		return 0, errors.New("queue is empty")
	}
	data := queue.data[0]
	queue.data = queue.data[1:]
	return data, nil
}

// Size :=  return the size of the Queue
func (queue *Queue) Size() int {
	return len(queue.data)
}

// Peek := Returns the element at front of the queue without deleting it
func (queue *Queue) Peek() (int, error) {
	if queue.Size() == 0 {
		return 0, errors.New("queue is empty")
	}
	return queue.data[0], nil
}

// Print := prints all the element in the queue
func (queue *Queue) Print() {
	for i := 0; i < len(queue.data); i++ {
		fmt.Printf("%v | ", queue.data[i])
	}
	fmt.Println()
}

func main() {
	queue := NewQueue(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()
	fmt.Println(queue.Size())
	data, err := queue.Dequeue()
	fmt.Println(data, err)
	data, err = queue.Dequeue()
	fmt.Println(data, err)
	data, err = queue.Dequeue()
	fmt.Println(data, err)
	queue.Print()
	data, err = queue.Dequeue()
	fmt.Println(data, err)
	data, err = queue.Dequeue()
	fmt.Println(data, err)
	data, err = queue.Peek()
	fmt.Println(data, err)
	queue.Enqueue(4)
	queue.Print()
}
