package searching

func Linear(arr []int, item int) int {
	if len(arr) == 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}
	}

	return -1
}
