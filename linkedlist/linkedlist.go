package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) AddToFront(data int) {
	newNode := &Node{
		data: data,
		next: list.head,
	}
	list.head = newNode
	list.size++
}

func (list *LinkedList) AddToEnd(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	list.size++
}

func (list *LinkedList) RemoveToEnd(data int) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		list.size--
		return
	}

	current := list.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
		list.size--
	}
}

func (list *LinkedList) deleteFront() {
	if list.head != nil {
		list.head = list.head.next
		return
	}
}

func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
