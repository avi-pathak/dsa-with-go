package stack

import "github.com/avi-pathak/dsa-with-go/linkedlist"

type LinkedStack[T comparable] struct {
	stack linkedlist.DoublyLinkedList[T]
	len   int
}

func (s *LinkedStack[T]) Push(value T) {
	s.stack.Append(value)
	s.len++
}

func (s *LinkedStack[T]) Pop() {
	s.stack.Remove()
	s.len--
}

func (s *LinkedStack[T]) Peek() T {
	return s.stack.Last()
}

func (s *LinkedStack[T]) Print() {
	s.stack.Print()
}
