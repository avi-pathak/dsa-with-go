package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() (Tr T, err error) {
	if len(s.stack) >= 0 {
		var value T = s.stack[len(s.stack)-1]
		s.stack = s.stack[0 : len(s.stack)-1]
		return value, errors.New("stack is Empty")
	}
	return
}

func (s *Stack[T]) Peek() T {

	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}

func (s *Stack[T]) Print() {
	fmt.Print(s.stack)
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	} else {
		return false
	}

}
