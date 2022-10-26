package sorting

func RadixSort(arr []int) []int {
	var max int = getMax(arr)

	for i := 1; max/i > 0; i *= 10 {
		arr = countingSort(arr, i)
	}

	return arr
}

func getMax(arr []int) int {
	if len(arr) == 0 {

		return -1
	}
	var max int = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

	}

	return max
}
func countingSort(arr []int, decPlace int) []int {

	var count = make([]int, 10)
	// count the numbers
	for i := 0; i < len(arr); i++ {
		count[(arr[i]/decPlace)%(10)] += 1
	}
	//summation

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	var resultArr = make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		resultArr[count[(arr[i]/decPlace)%(10)]-1] = arr[i]
		count[(arr[i]/decPlace)%(10)] -= 1
	}

	return resultArr

}
