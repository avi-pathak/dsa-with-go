package sorting

// bubble sort
func BubbleSort(arr []int) []int {
	var length int = len(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				var t int = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
			}
		}
	}
	return arr

}