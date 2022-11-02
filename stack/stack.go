package stack

import "fmt"

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() {
	if len(s.stack) <= 1 {
		s.stack = []T{}
		return
	}

	s.stack = s.stack[0 : len(s.stack)-1]
}

func (s *Stack[T]) Peek() T {
	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) Print() {
	fmt.Print(s.stack)
}
