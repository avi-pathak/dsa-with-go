package bst

type Node struct {
	left  *Node
	right *Node
	value int
}

type BST struct {
	root *Node
}

type TraverseArgumentFn func(int)

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

func (bst *BST) Traverse(fn TraverseArgumentFn, TraverseType string) {
	switch TraverseType {
	case TraversalType.Inorder:
		inorder(bst.root, fn)
		break
	case TraversalType.Preorder:
		preorder(bst.root, fn)
		break
	case TraversalType.Postorder:
		postorder(bst.root, fn)
		break
	}
}

//internal methods

func inorder(root *Node, fn TraverseArgumentFn) {
	if root == nil {
		return
	}

	inorder(root.left, fn)

	fn(root.value)
	inorder(root.right, fn)
}
func preorder(root *Node, fn TraverseArgumentFn) {
	if root == nil {
		return
	}
	fn(root.value)
	preorder(root.left, fn)
	preorder(root.right, fn)
}
func postorder(root *Node, fn TraverseArgumentFn) {
	if root == nil {
		return
	}
	postorder(root.left, fn)
	postorder(root.right, fn)
	fn(root.value)
}
func newNode(value int) Node {
	return Node{
		left:  nil,
		right: nil,
		value: value,
	}
}
