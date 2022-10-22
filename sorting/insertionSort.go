package insertionSort

// insertion sort
func insertionSort(arr []float32) []float32 {

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			var k int = i - 1
			var t int = i
			var item float32 = arr[i]
			for k >= 0 && arr[k] > item {
				arr[t] = arr[k]
				k--
				t--
			}
			arr[k+1] = item

		}
	}
	return arr

}
