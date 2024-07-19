package main

import (
	"fmt"
	"structs/binarytree"
	"structs/linkedlist"
	"structs/queue"
)

func main() {
	//queue
	intQueue := queue.Queue[int]{}
	intQueue.Enqueue(1)
	intQueue.Enqueue(2)
	intQueue.Enqueue(3)

	deqValue, deqOk := intQueue.Dequeue()
	fmt.Println("Dequeue:", deqValue, deqOk)

	frontValue, frontOk := intQueue.Front()
	fmt.Println("Front:", frontValue, frontOk)
	fmt.Println("Size:", intQueue.Size())
	//binary tree

	tree := binarytree.BinaryTree{}
	tree.Insert(100)
	tree.Insert(200)
	tree.Insert(50)
	tree.Insert(50)

	//linked list
	list := linkedlist.LinkedList{}

	list.PrintList()
	list.AddToFront(1)
	list.AddToFront(1)
	list.PrintList()
	list.AddToEnd(10)
	list.PrintList()
	list.AddToFront(100)
	list.PrintList()
}
