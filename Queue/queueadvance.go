package main

import (
	"fmt"
)

//ArrayQueue := Queue implementation backed by Array/Slice
type ArrayQueue struct {
	Queue []int
	Front int
	Back  int
}

//NewQueue := Creates a new Queue with the provided capacity
func NewQueueAdvance(capacity int) *ArrayQueue {
	queue := new(ArrayQueue)
	queue.Queue = make([]int, capacity)
	return queue
}

//Add := inserts the element at the front of the queue
func (queue *ArrayQueue) Add(data int) {
	if queue.Back == len(queue.Queue) {
		newQueue := make([]int, 2*len(queue.Queue))
		copy(newQueue, queue.Queue)
		queue.Queue = newQueue
	}
	queue.Queue[queue.Back] = data
	queue.Back++
}

//Remove:= removes the element from the back of the queue
func (queue *ArrayQueue) Remove() int {
	if queue.Size() == 0 {
		panic("No such element")
	}

	data := queue.Queue[queue.Front]
	queue.Queue[queue.Front] = 0
	queue.Front++
	if queue.Size() == 0 {
		queue.Front = 0
		queue.Back = 0
	}
	return data
}

//Peek := Returns the element at front of the queue without deleting it
func (queue *ArrayQueue) Peek() int {
	if queue.Size() == 0 {
		panic("No such element")
	}

	return queue.Queue[queue.Front]
}

//Size := return the size of the Queue
func (queue *ArrayQueue) Size() int {
	return queue.Back - queue.Front
}

//Print := prints all the element in the queue
func (queue *ArrayQueue) Print() {
	for i := queue.Front; i < queue.Back; i++ {
		fmt.Printf("%v | ", queue.Queue[i])
	}
	fmt.Println()
}

func main() {
	queue := NewQueueAdvance(2)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Print()
	queue.Remove()
	queue.Remove()
	queue.Peek()
	queue.Print()
	queue.Add(4)
}