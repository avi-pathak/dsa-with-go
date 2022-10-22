package sorting

func CountingSort(arr []int) []int {

	var max int = -1

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

	}
	var count = make([]int, max+1)
	// count the numbers
	for i := 0; i < len(arr); i++ {
		count[arr[i]] += 1
	}
	//summation

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	var resultArr = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		resultArr[count[arr[i]]-1] = arr[i]
		count[arr[i]] -= 1
	}

	return resultArr

}
