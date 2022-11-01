package bst

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(value int) {
	root := bst.root
	tempNode := bst.root
	node := newNode(value)
	if root == nil {
		bst.root = &node
		return
	}

	for root != nil {
		tempNode = root
		if root.value >= value {
			root = root.left
		} else if root.value <= value {
			root = root.right
		}
	}

	if tempNode.value >= value {
		tempNode.left = &node
		return
	} else if tempNode.value <= value {
		tempNode.right = &node
		return
	}
	return

}

func newNode(value int) Node {
	return Node{
		left:  nil,
		right: nil,
		value: value,
	}
}

func (bst *BST) Inorder() []int {
	root := bst.root
	inorder(root)
	return []int{}
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.value)
	inorder(root.right)
}
