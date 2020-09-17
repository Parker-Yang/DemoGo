package quick_sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 3, 7, 6, 4, 1, 0, 2, 9, 10, 8}
	sort := QuickSort(arr, 0, len(arr)-1)
	t.Log(sort)
}
func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	rand.Shuffle(10000, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	for i := 0; i < b.N; i++ {
		_ = QuickSort(arr, 0, len(arr)-1)
	}
}
