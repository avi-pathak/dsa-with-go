package queue

import (
	"errors"
	"fmt"
)

type Queue[T comparable] struct {
	queue []T
}

func (q *Queue[T]) EnQueue(value T) {
	q.queue = append(q.queue, value)
}

func (q *Queue[T]) Print() {
	fmt.Println(q.queue)
}

func (q *Queue[T]) DeQueue() T {
	if len(q.queue) == 0 {
		fmt.Println("queue is empty")
		errors.New("empty name")
	}
	var res T = q.queue[0]
	q.queue = q.queue[1:]

	return res

}

func (q *Queue[T]) Peek() T {
	if len(q.queue) == 0 {
		fmt.Println("queue is empty")
		errors.New("empty name")
	}

	return q.queue[len(q.queue)-1]

}

func (q *Queue[T]) IsEmpty() bool {
	if len(q.queue) == 0 {
		return true
	}

	return false

}
