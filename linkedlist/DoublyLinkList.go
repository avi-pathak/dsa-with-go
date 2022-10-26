package linkedlist

import "fmt"

// Node represents a node of linked list
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// LinkedList represents a linked list
type DoublyLinkedList struct {
	head *Node
	len  int
}

// func GetLinkListFromArray(arr []int) []Node {
// 	for i := 0; i < len(arr); i++ {
// 	}
// }

func (l *DoublyLinkedList) Append(value int) {
	n := newNode(value)
	if l.len == 0 {
		n.next = nil
		n.prev = nil

		l.head = &n
		l.len++
		return
	}

	// if l.len == 1 {
	// 	l.head.next = &n
	// 	n.prev = l.head
	// 	return
	// }

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

func (l *DoublyLinkedList) AppendAt(value int, position int) {
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

func (l *DoublyLinkedList) Print() {
	ptr := l.head

	for ptr != nil {
		fmt.Print(ptr.value, "  ")
		ptr = ptr.next

	}

}

func (l *DoublyLinkedList) Remove() {
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

func (l *DoublyLinkedList) RemoveAt(position int) {
	ptr := l.head

	for i := 0; i < (position - 1); i++ {
		ptr = ptr.next
	}

	ptr.next.prev = ptr.prev

	ptr.prev.next = ptr.next

	l.len--

}

func newNode(value int) Node {
	return Node{
		prev:  nil,
		value: value,
		next:  nil,
	}
}
