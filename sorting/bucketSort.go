package sorting

import "fmt"

func BucketSort(arr []float32) []float32 {
	var n int = len(arr)
	var bucket = make(map[int][]float32, len(arr))

	for i := 0; i < n; i++ {
		bucket[int(arr[i]*10)] = []float32{}
	}
	for i := 0; i < n; i++ {
		var key int = int(arr[i] * 10)
		var _, ok = bucket[key]

		if ok {
			bucket[key] = InsertionSort(append(bucket[key], arr[i]))
		}
	}

	fmt.Println("bucket", bucket)
	fmt.Println()
	keys := make([]int, 0, len(bucket))
	for k := range bucket {
		keys = append(keys, k)
	}
	keys = BubbleSort(keys)
	var res = make([]float32, 0)
	for i := 0; i < len(keys); i++ {
		res = append(res, bucket[int(keys[i])]...)
	}
	return res
}
