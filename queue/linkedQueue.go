package queue

import "github.com/avi-pathak/dsa-with-go/linkedlist"

type LinkedQueue[T comparable] struct {
	queue linkedlist.DoublyLinkedList[T]
}

func (q *LinkedQueue[T]) EnQueue(value T) {
	q.queue.Append(value)
}

func (q *LinkedQueue[T]) DeQueue() {
	q.queue.RemoveAt(0)
}

func (q *LinkedQueue[T]) Print() {
	q.queue.Print()
}
func (q *LinkedQueue[T]) Peek() T {
	return q.queue.Last()
}
