package main

import (
	"errors"
	"fmt"
)

// Stack definition of stack object
type Stack struct {
	top      int
	capacity int
	data     []interface{}
}

// init intialises a new stack object and returns a pointer for the same
func (stack *Stack) init(capacity int) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.data = make([]interface{}, capacity)

	return stack
}

// NewStack creates a new stack object
func NewStack(capacity int) *Stack {
	return new(Stack).init(capacity)
}

// IsFull checks whether the stack has reached it's capacity or not
func (stack *Stack) IsFull() bool {
	return stack.capacity == len(stack.data)-1
}

// IsEmpty checks whether the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

// Size returns the current size of the Stack
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

// Peek returns the data without removing from the top of stack
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	peekelement := stack.data[stack.top]
	return peekelement, nil
}

// Clear removes all the elements in the stack
func (stack *Stack) Clear() {
	stack.data = nil
	stack.top = -1
}

// Push inserts the data in the top of the stack
func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack overflow")
	}
	stack.top++
	stack.data[stack.top] = data

	return nil
}

// Pop removes the data from the top of the stack
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Stack Underflow")
	}
	popelement := stack.data[stack.top]
	stack.top--

	return popelement, nil
}

func (stack *Stack) Print() {
	fmt.Print("[")
	for i := 0; i <= stack.top; i++ {
		fmt.Print(stack.data[i])
		if i != stack.top {
			fmt.Print(",")
		}
	}
	fmt.Print("]\n")
}

func main() {
	stack := NewStack(10)
	stack.Print()
	stack.Push(7)
	stack.Print()
	stack.Push(9)
	stack.Print()
	stack.Push(8)
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	fmt.Println("Size", stack.Size())
	peek, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Peek element", peek)
	fmt.Println("capacity", stack.capacity)
}
