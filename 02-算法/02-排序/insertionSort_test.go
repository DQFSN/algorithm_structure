package sort

import (
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	arrs := getTestData()

	for _, arr := range arrs {
		temp := make([]int, len(arr), len(arr))
		copy(temp,arr)
		insertionSort(arr)
		quickSort(temp)

		if !checkSame(arr, temp) {
			t.Fatalf("测试不通过, a=%v, b=%v",arr,temp)
		}
	}
}
