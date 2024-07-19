package binarytree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) Insert(data int) {
	if tree.root == nil {
		tree.root = &Node{data: data}
	} else {
		tree.root.insert(data)
	}
}

func (n *Node) insert(data int) {
	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (tree *BinaryTree) Search(data int) bool {
	return tree.root.search(data)
}

func (n *Node) search(data int) bool {
	if n == nil {
		return false
	}
	if data < n.data {
		return n.left.search(data)
	} else if data > n.data {
		return n.right.search(data)
	}
	return true
}

func (tree *BinaryTree) InOrder() {
	tree.root.inOrder()
	fmt.Println()
}

func (n *Node) inOrder() {
	if n == nil {
		return
	}
	n.left.inOrder()
	fmt.Printf("%d ", n.data)
	n.right.inOrder()
}

func (tree *BinaryTree) PreOrder() {
	tree.root.preOrder()
	fmt.Println()
}

func (n *Node) preOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.data)
	n.left.preOrder()
	n.right.preOrder()
}

func (tree *BinaryTree) PostOrder() {
	tree.root.postOrder()
	fmt.Println()
}

func (n *Node) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d ", n.data)
}
