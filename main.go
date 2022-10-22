package main

import (
	"fmt"

	"github.com/avi-pathak/dsa-with-go/sorting"
)

func main() {
	unsortedArr := []float32{9.0, 8.0, 7.0, 65.0, 54.0, 2.0, 1.0, 3.0}

	fmt.Println("InsertionSort")
	var sortedArrFromInsertionSort = sorting.InsertionSort(unsortedArr)
	fmt.Println(sortedArrFromInsertionSort)
	fmt.Println("BubbelSort")

	arr := []int{5, 4, 1, 3, 2, 45, 2, 54, 56}
	fmt.Println(sorting.BubbleSort(arr))
	fmt.Println("SelectionSort")
	arr2 := []int{5, 4, 1}
	fmt.Println(sorting.SelectionSort(arr2))

	fmt.Println("MergeSort")
	fmt.Println(sorting.MergeSort(arr))

	fmt.Println("Counting Sort ")
	fmt.Println(sorting.CountingSort(arr))

	unsorted := []float32{0.15, 0.78, 0.63, .585, 0.578, 0.564, 0.5}
	sorted := sorting.BucketSort(unsorted)
	fmt.Println("unsorted Array", unsorted)
	fmt.Println()
	fmt.Println("sorted Array", sorted)
}
