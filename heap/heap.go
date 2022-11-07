package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	heap []T
}

func (h *Heap[T]) Add(value T) {
	h.heap = append(h.heap, value)
	h.heap = heapify[T](h.heap)
}

func heapify[T constraints.Ordered](arr []T) []T {
	var count int = len(arr) - 1
	for count >= 0 {
		var leftIndex = (2 * count) + 1
		var rightIndex = (2 * count) + 2

		if leftIndex > len(arr) && rightIndex > len(arr) {
			count--
			continue
		}

		if leftIndex < len(arr) && arr[leftIndex] > arr[count] {
			arr[count], arr[leftIndex] = arr[leftIndex], arr[count]
		}

		if rightIndex < len(arr) && arr[rightIndex] > arr[count] {
			arr[count], arr[rightIndex] = arr[rightIndex], arr[count]
		}

		count--
	}

	return arr
}

func (h *Heap[T]) GetHeapAsArray() []T {
	return h.heap

}
func (h *Heap[T]) GetMaxNode() T {
	return h.heap[0]
}
