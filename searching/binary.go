package searching

func Binary(arr []int, item int) int {
	if len(arr) == 0 {
		return -1
	}

	return bin(arr, 0, len(arr)-1, item)
}

func bin(arr []int, start int, end int, item int) int {
	if start > end || start < 0 || end > len(arr) {
		return -1
	} else {
		var mid int = (start + end) / 2
		if arr[mid] == item {
			return mid
		} else if arr[mid] < item {
			return bin(arr, start, mid, item)
		} else {
			return bin(arr, mid+1, end, item)
		}
	}

}
