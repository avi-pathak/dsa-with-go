package sorting


func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := MergeSort(arr[:len(arr)/2])
	second := MergeSort(arr[len(arr)/2:])
	return merge(first, second)
}
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
