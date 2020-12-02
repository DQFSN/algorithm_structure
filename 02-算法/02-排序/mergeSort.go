package sort

import (
	"runtime"
	"sync"
)

////划分--归并
//func mergeSort(arr []int) {
//	if len(arr) > 1 {
//		middle := len(arr) / 2
//		mergeSort(arr[:middle])
//		mergeSort(arr[middle:])
//
//		temp := make([]int, len(arr), len(arr))
//		index, l, r := 0, 0, middle
//
//		for l < middle && r < len(arr) {
//			if arr[l] <= arr[r] {
//				temp[index] = arr[l]
//				l++
//			}else {
//				temp[index] = arr[r]
//				r++
//			}
//			index++
//		}
//		if l == middle {
//			copy(temp[index:],arr[r:])
//		}else {
//			copy(temp[index:], arr[l:middle])
//		}
//
//		copy(arr, temp)
//	}
//}



func merge(arr []int) {
	middle := len(arr) / 2


	left, right := make([]int, middle, middle), make([]int, len(arr) - middle, len(arr) - middle)
	copy(left, arr[:middle])
	copy(right,arr[middle:])
	leftIter, rightIter := ints(left).Iter(), ints(right).Iter()


	leftValue, hasNextLeft := leftIter()
	rightValue, hasNextRight := rightIter()
	for index := range arr {
		if !hasNextLeft {
			arr[index] = rightValue
			rightValue, hasNextRight = rightIter()
		}else if !hasNextRight {
			arr[index] = leftValue
			leftValue, hasNextLeft = leftIter()
		}else {
			if leftValue <= rightValue {
				arr[index] = leftValue
				leftValue, hasNextLeft = leftIter()
			}else {
				arr[index] = rightValue
				rightValue, hasNextRight = rightIter()
			}
		}
	}

}

func mergeSort(arr []int) {
	middle := len(arr) / 2

	if middle > 0 {
		mergeSort(arr[:middle])
		mergeSort(arr[middle:])
		merge(arr)
	}

}

func mergeSortParallel(arr []int) {
	runtime.GOMAXPROCS(2)
	i := len(arr) / 2
	if i > 0 {
		var wg = sync.WaitGroup{}
		wg.Add(2)
		go func() {
			mergeSortParallel(arr[:i])
			wg.Done()
		}()

		go func() {
			mergeSortParallel(arr[i:])
			wg.Done()
		}()

		wg.Wait()
		merge(arr)
	}
}