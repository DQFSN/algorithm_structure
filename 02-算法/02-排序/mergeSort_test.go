package sort

import "testing"

func Test_MergeSort(t *testing.T) {
	arrs := [][]int {
		{1,32,32,4,3,45,454,564,56,5454,3,4,4,5,7,6,76,321,6,8},
		{214,3,4,35,4,656,7,57,67,876,64,54,553,4,35,5,7,68,7,5},
	}

	for _, arr := range arrs {
		temp := make([]int, len(arr), len(arr))
		copy(temp, arr)

		//mergeSort(arr)
		mergeSortParallel(arr)
		quickSort(temp)

		if !checkSame(arr, temp) {
			t.Fatalf("测试不通过, a=%v, b=%v",arr,temp)
		}
	}
}
