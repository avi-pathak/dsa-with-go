package main

import (
	"fmt"

	bst "github.com/avi-pathak/dsa-with-go/BST"
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

	// l := linkedlist.DoublyLinkedList[int]{}

	// l.Print()
	// l.Append(12)
	// l.Append(13)
	// l.Append(14)
	// l.Append(15)
	// l.AppendAt(44, 1)
	// l.AppendAt(10, 1)
	// l.AppendAt(10, 1)
	// l.Remove()
	// l.Remove()
	// l.RemoveAt(3)
	// l.Print()
	// fmt.Println("index of 13 is:", l.IndexOf(13))
	// fmt.Println("index of 44 is:", l.IndexOf(44))
	// fmt.Println("index of 12 is:", l.IndexOf(12))

	// s := stack.LinkedStack[int]{}
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)
	// s.Pop()
	// s.Pop()
	// s.Push(3)
	// s.Push(4)
	// s.Push(5)
	// s.Pop()
	// s.Pop()
	// s.Push(6)
	// s.Push(7)
	// s.Push(8)
	// s.Pop()
	// fmt.Println("Stack is ")
	// s.Print()
	// fmt.Println()

	// fmt.Println("Peek Element is:", s.Peek())

	// q := queue.Queue[int]{}
	// q.EnQueue(1)
	// q.EnQueue(2)
	// q.EnQueue(3)
	// q.DeQueue()
	// q.DeQueue()
	// q.EnQueue(3)
	// q.EnQueue(4)
	// q.EnQueue(5)
	// q.DeQueue()
	// q.DeQueue()
	// q.EnQueue(6)
	// q.EnQueue(7)
	// q.EnQueue(8)
	// q.DeQueue()
	// fmt.Println("Queue is ")
	// q.Print()
	// fmt.Println("Peek Element is ", q.Peek())
	// fmt.Println()

	// q := queue.LinkedQueue[int]{}
	// q.EnQueue(1)
	// q.EnQueue(2)
	// q.EnQueue(3)
	// q.DeQueue()
	// q.DeQueue()
	// q.EnQueue(3)
	// q.EnQueue(4)
	// q.EnQueue(5)
	// q.DeQueue()
	// q.DeQueue()
	// q.EnQueue(6)
	// q.EnQueue(7)
	// q.EnQueue(8)
	// q.DeQueue()
	// fmt.Println("Queue is ")
	// q.Print()
	// fmt.Println()
	// fmt.Println("Peek Element is ", q.Peek())

	//
	//Set
	//

	// s := set.Set[int]{}

	// s.Add(10)

	// s.Add(10)
	// fmt.Println("\n Add 10")
	// s.Add(10)
	// fmt.Println("\n Add 15")
	// s.Add(15)
	// fmt.Println("\n Add 17")
	// s.Add(17)
	// fmt.Println("\n Add 11")
	// s.Add(11)
	// fmt.Println("\n has 17", s.Has(17))
	// fmt.Println("\n delete 11")
	// s.Delete(11)
	// fmt.Println("\n has 11", s.Has(11))
	// fmt.Println("\n Count ", s.Count())

	tree := bst.BST[int]{}

	fmt.Println("insert 30")
	tree.Insert(30)
	fmt.Println("insert 20")
	tree.Insert(20)
	fmt.Println("insert 40")
	tree.Insert(40)
	fmt.Println("insert 10")
	tree.Insert(10)
	fmt.Println("insert 25")
	tree.Insert(25)
	fmt.Println("insert 35")
	tree.Insert(35)
	fmt.Println("insert 50")
	tree.Insert(50)
	fmt.Println("Inorder traversal")
	tree.Traverse(printValue, bst.TraversalType.Inorder)
	fmt.Println("\nPreorder traversal")
	tree.Traverse(printValue, bst.TraversalType.Preorder)
	fmt.Println("\nPostorder traversal")
	tree.Traverse(printValue, bst.TraversalType.Postorder)
	fmt.Println("\nlevel order traversal Itrative ")
	tree.Traverse(printValue, bst.TraversalType.LevelOrderTraversal)
	fmt.Println("\ninorder itartive traversal Itrative ")
	tree.Traverse(printValue, bst.TraversalType.InorderItrative)
	fmt.Println("\npreorder itrative traversal Itrative ")
	tree.Traverse(printValue, bst.TraversalType.PreorderItarative)
	fmt.Println("\npostorder itrative traversal Itrative ")
	tree.Traverse(printValue, bst.TraversalType.PostorderItarative)
	fmt.Println("\nSearch 40")
	n := tree.Search(40)

	if n != nil {
		fmt.Println("found")
	} else {
		fmt.Println("not found")

	}
	fmt.Println("\nHeight: ", tree.Height())
	fmt.Println("\nis valid BST: ", tree.IsValidBST())
	fmt.Println("delete Node 25: ", tree.Delete(25))
	fmt.Println("delete Node 35: ", tree.Delete(35))
	fmt.Println("delete Node 40: ", tree.Delete(40))
	fmt.Println("delete Node 20: ", tree.Delete(20))
	fmt.Println("delete Node 20: ", tree.Delete(20))
	fmt.Println("delete Node 30: ", tree.Delete(30))

	fmt.Println("inorder: ")
	tree.Traverse(printValue, bst.TraversalType.InorderItrative)

}

func printValue(value int) {
	fmt.Printf("%d  ", value)
}

/*
       30
	/      \
   20      40
  /  \     /  \
 10   25 35   50

 10 13 15 18
 15 13 10 18
 10 13 18 15
 15 13 18 10

 10 15 20 35 50 40 30
*/
