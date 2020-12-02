package sort

import (
	"math/rand"
)

func partition(arr []int) (primeIndex int)  {

	primeIndex = 0

	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] < arr[len(arr) - 1] {
			arr[i], arr[primeIndex] = arr[primeIndex],arr[i]
			primeIndex++
		}
	}
	arr[primeIndex], arr[len(arr) - 1] = arr[len(arr) - 1], arr[primeIndex]

	return
}

func quickSort(arr []int){
	if len(arr) > 1 {
		primeIndex := partition(arr)
		quickSort(arr[:primeIndex])
		quickSort(arr[primeIndex+1:])
	}
}


func randQuickSort(arr []int) {
	if len(arr) > 1{
		primeIndex := rand.Intn(len(arr))
		arr[primeIndex], arr[len(arr) - 1] = arr[len(arr) - 1], arr[primeIndex]
		primeIndex = partition(arr)
		randQuickSort(arr[:primeIndex])
		randQuickSort(arr[primeIndex+1:])
	}
}

