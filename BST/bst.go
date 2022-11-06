package bst

import (
	set "github.com/avi-pathak/dsa-with-go/Set"
	"github.com/avi-pathak/dsa-with-go/queue"
	"github.com/avi-pathak/dsa-with-go/stack"
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	left  *Node[T]
	right *Node[T]
	value T
}

type BST[t constraints.Ordered] struct {
	root *Node[t]
}

type TraverseArgumentFn[T constraints.Ordered] func(T)

func (bst *BST[T]) Insert(value T) {
	root := bst.root
	tempNode := bst.root
	node := newNode(value)
	if root == nil {
		bst.root = &node
		return
	}

	for root != nil {
		tempNode = root
		if value >= root.value {
			root = root.right
		} else if value <= root.value {
			root = root.left
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
func (bst *BST[T]) Delete(value T) bool {

	root := bst.root

	var nodeToDelete *Node[T] = bst.Search(value)
	if nodeToDelete == nil {
		return false
	}
	var tempNode *Node[T] = getParentNode(root, nodeToDelete)
	//case 1 where node is a leaf node
	//node to delete is a leaf node just need to de
	if nodeToDelete.left == nil && nodeToDelete.right == nil {
		deleteLeafNode(root, nodeToDelete)
	} else {

		//case 2  where one of sub tree is nil
		// grandParent := getParentNode(root, tempNode)
		if nodeToDelete.left == nil {
			//to check whther the parent if from left or right
			if tempNode.left != nil && tempNode.left.value == nodeToDelete.value {
				tempNode.left = nodeToDelete.right
			} else {
				tempNode.right = nodeToDelete.right
			}
		} else if nodeToDelete.right == nil {
			if tempNode.left != nil && tempNode.left.value == nodeToDelete.value {
				tempNode.left = nodeToDelete.left
			} else {
				tempNode.right = nodeToDelete.left
			}
		} else {
			//case 3 where node has both the childern
			inorderPredecesor := getMinNode(nodeToDelete)
			deleteLeafNode(root, inorderPredecesor)
			nodeToDelete.value = inorderPredecesor.value
		}
	}

	return true
}

func getMinNode[T constraints.Ordered](node *Node[T]) *Node[T] {
	for node.left != nil {
		node = node.left
	}
	return node

}

func getParentNode[T constraints.Ordered](root *Node[T], node *Node[T]) *Node[T] {
	var tempNode *Node[T] = nil
	for root != nil {
		if root.value == node.value {
			return tempNode
		}
		tempNode = root
		if root.value > node.value {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil
}

func getMaxNode[T constraints.Ordered](node *Node[T]) *Node[T] {
	for node.right != nil {
		node = node.right
	}
	return node

}
func deleteLeafNode[T constraints.Ordered](root *Node[T], node *Node[T]) {
	parentNode := getParentNode(root, node)
	if parentNode.left.value == node.value {
		parentNode.left = nil
	} else {
		parentNode.right = nil
	}

}
func (bst *BST[T]) IsValidBST() bool {
	return isBST(bst.root)
}
func (bst *BST[T]) Traverse(fn TraverseArgumentFn[T], TraverseType string) {
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
	case TraversalType.InorderItrative:
		inorderItarative(bst.root, fn)
		break
	case TraversalType.LevelOrderTraversal:
		levelOrderTraversal(bst.root, fn)
		break
	case TraversalType.PreorderItarative:
		preorderItarative(bst.root, fn)
		break

	case TraversalType.PostorderItarative:
		postorderItrative(bst.root, fn)
		break
	}
}

func (bst *BST[T]) Search(value T) *Node[T] {
	root := bst.root
	for root != nil {
		if root.value == value {
			return root
		}
		if root.value > value {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil
}

func (bst *BST[T]) Height() int {
	root := bst.root
	return height[T](root)
}

// internal methods

func isBST[T constraints.Ordered](root *Node[T]) bool {
	if root == nil || root.left == nil && root.right == nil {
		return true
	}
	var l bool = true
	var r bool = true
	if root.left != nil && root.left.value > root.value {
		l = false
	}
	if root.right != nil && root.right.value < root.value {
		r = false
	}
	return r && l && isBST(root.left) && isBST(root.right)
}
func height[T constraints.Ordered](root *Node[T]) int {
	if root == nil {
		return 0
	} else {
		var d = max(height(root.left), height(root.right))
		return d + 1
	}

}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func levelOrderTraversal[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}

	q := queue.Queue[*Node[T]]{}

	q.EnQueue(root)
	for !q.IsEmpty() {
		node := q.DeQueue()
		fn(node.value)
		//push the left level
		if node.left != nil {
			q.EnQueue(node.left)
		}
		//push right level
		if node.right != nil {
			q.EnQueue(node.right)
		}
	}

}
func inorderItarative[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}
	st := stack.Stack[*Node[T]]{}
	st.Push(root)
	visited := set.Set[T]{}
	for !st.IsEmpty() {
		var el, _ = st.Pop()
		if el.left != nil && !visited.Has(el.left.value) {
			st.Push(el)
			st.Push(el.left)
		} else {
			fn(el.value)
			visited.Add(el.value)
			if el.right != nil {
				st.Push(el.right)
			}
		}

	}

}

func preorderItarative[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}
	st := stack.Stack[*Node[T]]{}
	st.Push(root)
	visited := set.Set[T]{}
	for !st.IsEmpty() {
		var el, _ = st.Pop()
		fn(el.value)
		visited.Add(el.value)
		if el.right != nil && !visited.Has(el.right.value) {
			st.Push(el.right)
		}
		if el.left != nil && !visited.Has(el.left.value) {
			st.Push(el.left)
		}

	}

}

func postorderItrative[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}
	st := stack.Stack[*Node[T]]{}
	st.Push(root)
	visited := set.Set[T]{}
	for !st.IsEmpty() {
		visited.Add(st.Peek().value)
		if st.Peek().left != nil {
			if !visited.Has(st.Peek().left.value) {
				st.Push(st.Peek().left)
				continue
			}
		}

		// Otherwise if the right child of the top node is
		// not NULL and not visited push it into the stack
		if st.Peek().right != nil {
			if !visited.Has(st.Peek().right.value) {
				st.Push(st.Peek().right)
				continue
			}
		}
		// var el, _ = st.Pop()

		// if el.left != nil && !visited.Has(el.left.value) {
		// 	st.Push(el)
		// 	if el.right != nil {
		// 		st.Push(el.right)
		// 	}
		// 	st.Push(el.left)
		// 	continue

		// }

		fn(st.Peek().value)
		// visited.Add(el.value)
		st.Pop()
		// if el.left == nil {
		// 	fn(el.value)
		// 	visited.Add(el.value)
		// }
		// if(el.left!= null && visited.Has()(el.left.value)
		// if el.right != nil && !visited.Has(el.right.value) {
		// 	if el.left != nil && !visited.Has(el.left.value) {
		// 		st.Push(el.left)
		// 	}
		// 	st.Push(el.right)
		// }

	}
}

/*
10,13,
15,18,13
*/
func inorder[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}

	inorder(root.left, fn)
	fn(root.value)
	inorder(root.right, fn)
}

func preorder[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}
	fn(root.value)
	preorder(root.left, fn)
	preorder(root.right, fn)
}
func postorder[T constraints.Ordered](root *Node[T], fn TraverseArgumentFn[T]) {
	if root == nil {
		return
	}
	postorder(root.left, fn)
	postorder(root.right, fn)
	fn(root.value)
}
func newNode[T constraints.Ordered](value T) Node[T] {
	return Node[T]{
		left:  nil,
		right: nil,
		value: value,
	}
}
