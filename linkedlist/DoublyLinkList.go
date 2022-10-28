package linkedlist

import "fmt"

// Node represents a node of linked list
type Node[T comparable] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

// LinkedList represents a linked list
type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	len  int
}

func (l *DoublyLinkedList[T]) Append(value T) {
	n := newNode(value)
	if l.len == 0 {
		n.next = nil
		n.prev = nil

		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			n.prev = ptr
			l.len++
			return
		}
		ptr = ptr.next
	}

}

func (l *DoublyLinkedList[T]) AppendAt(value T, position int) {
	n := newNode(value)
	if l.len == 0 {
		n.next = nil
		n.prev = nil

		l.head = &n
		l.len++
		return
	}

	if position == 1 {
		l.head.prev = &n
		n.next = l.head
		l.head = &n
		l.len++
		return
	}

	ptr := l.head

	for i := 0; i < position-1; i++ {
		ptr = ptr.next
	}
	n.next = ptr.next
	n.prev = ptr
	ptr.next = &n

	l.len++

}

func (l *DoublyLinkedList[T]) Print() {
	ptr := l.head

	for ptr != nil {
		fmt.Print(ptr.value, "  ")
		ptr = ptr.next

	}

}

func (l *DoublyLinkedList[T]) Remove() {
	ptr := l.head

	if l.len == 1 {
		l.len--
		l.head = nil
		return

	}

	for ptr.next != nil {
		ptr = ptr.next
	}
	// fmt.Print("asdf", ptr.value, ptr.prev.value)
	ptr.prev.next = nil
	ptr.prev = nil
	l.len--

}

func (l *DoublyLinkedList[T]) RemoveAt(position int) {
	ptr := l.head

	for i := 0; i < (position - 1); i++ {
		ptr = ptr.next
	}

	ptr.next.prev = ptr.prev

	ptr.prev.next = ptr.next

	l.len--

}

func newNode[T comparable](value T) Node[T] {
	return Node[T]{
		prev:  nil,
		value: value,
		next:  nil,
	}
}

func (l *DoublyLinkedList[T]) IndexOf(value T) int {
	ptr := l.head

	var index int = 0

	for ptr != nil {
		if ptr.value == value {
			return index
		}
		ptr = ptr.next
		index += 1
	}

	return -1
}
