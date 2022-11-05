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

// internal methods
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
