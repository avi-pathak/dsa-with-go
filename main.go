package main

import (
	"github.com/avi-pathak/dsa-with-go/linkedlist"
)

func main() {
	// unsortedArr := []float32{9.0, 8.0, 7.0, 65.0, 54.0, 2.0, 1.0, 3.0}

	// fmt.Println("InsertionSort")
	// var sortedArrFromInsertionSort = sorting.InsertionSort(unsortedArr)
	// fmt.Println(sortedArrFromInsertionSort)
	// fmt.Println("BubbelSort")

	// arr := []int{5, 4, 1, 3, 2, 45, 2, 54, 56}
	// fmt.Println(sorting.BubbleSort(arr))
	// fmt.Println("SelectionSort")
	// arr2 := []int{5, 4, 1}
	// fmt.Println(sorting.SelectionSort(arr2))

	// fmt.Println("MergeSort")
	// fmt.Println(sorting.MergeSort(arr))

	// fmt.Println("Counting Sort ")
	// fmt.Println(sorting.CountingSort(arr))

	// fmt.Println("BUcket Sort ")
	// unsorted := []float32{0.15, 0.78, 0.63, .585, 0.578, 0.564, 0.5}
	// sorted := sorting.BucketSort(unsorted)
	// fmt.Println("unsorted Array", unsorted)
	// fmt.Println()
	// fmt.Println("sorted Array", sorted)

	// fmt.Println("Radix Sort ")
	// unsortedForRadix := []int{455, 55, 85, 755, 855, 322, 711, 933}
	// sortedFromRadix := sorting.RadixSort(unsortedForRadix)
	// fmt.Println("unsorted Array", unsortedForRadix)
	// fmt.Println()
	// fmt.Println("sorted Array", sortedFromRadix)

	// fmt.Println("Array is: ", unsortedForRadix)
	// fmt.Println("Sorted Array is: ", sortedFromRadix)
	// fmt.Println("Linear search Attempting find value 322")
	// fmt.Println("value found using linear search at index", searching.Linear(unsortedForRadix, 322))

	// fmt.Println("Binary search Attempting find value 322")
	// fmt.Println("")
	// fmt.Println("value found using Binary search at index", searching.Binary(sortedFromRadix, 322))

	l := linkedlist.DoublyLinkedList{}

	l.Print()
	l.Append(12)
	l.Append(13)
	l.Append(14)
	l.Append(15)
	l.AppendAt(44, 1)
	l.AppendAt(10, 1)
	l.AppendAt(10, 1)
	l.Remove()
	l.Remove()
	l.RemoveAt(3)
	l.Print()
}
