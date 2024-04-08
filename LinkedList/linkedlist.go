package main

import "fmt"

type Node struct {
	data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

// method inserts a new node with the given data at the end of the linked list.
func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, Next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	if current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// method inserts a new node with the given data at the beginning (front) of the linked list.
func (list *LinkedList) insertAtFront(data int) {

	if list.head == nil {
		newNode := &Node{data: data, Next: nil}
		list.head = newNode
		return
	}

	newNode := &Node{data: data, Next: list.head}
	list.head = newNode
}

// method inserts a new node with the given data immediately after
// the first occurrence of a specific afterValue in the linked list.
func (list *LinkedList) insertAfterValue(afterValue, data int) {
	newNode := &Node{data: data, Next: nil}

	current := list.head

	if current.Next != nil {
		if current.data == afterValue {
			newNode.Next = current.Next
			current.Next = newNode
			return
		}
		current = current.Next
	}
	fmt.Printf("Cannot insert node, after value %d doesn't exist", afterValue)
	fmt.Println()
}

// method inserts a new node with the given data immediately before
// the first occurrence of a specific beforeValue in the linked list.
func (list *LinkedList) insertBeforeValue(beforeValue, data int) {
	if list.head == nil {
		return
	}
	if list.head.data == beforeValue {
		newNode := &Node{data: data, Next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	if current.Next != nil {
		if current.Next.data == beforeValue {
			newNode := &Node{data: data, Next: current.Next}
			current.Next = newNode
			return
		}
		current = current.Next
	}
	fmt.Printf("Cannot insert node, before value %d doesn't exist", beforeValue)
	fmt.Println()
}

//  method prints the data values of all nodes in the linked list to the console
func (list *LinkedList) Print() {
	current := list.head
	for current.Next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.Next
	}
	fmt.Printf("%d -> ", current.data)
	fmt.Println()
}

func main() {
	// Create an empty linked list
	myList := LinkedList{}

	// Insert some data at the beginning
	myList.insertAtFront(10)
	myList.insertAtFront(20)
	myList.insertAtFront(30)
	myList.insertAtFront(40)

	// Print the contents of the linked list
	fmt.Println("Linked List Contents:")
	myList.Print()
}
