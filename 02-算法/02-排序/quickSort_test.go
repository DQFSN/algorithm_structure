package sort

import (
	"testing"
)

func Test_QuickSort(t *testing.T)  {
	arrs := [][]int{
		{1,4,2123,2,454,54,546,5,43,2,45,476,5,867,56,343,2,4234,346,768,6,5665},
		{102,5,12,4,3,45,3,64,6567,67,55,4,45,6,76,554,4,67,687,87,546,3,54},
	}

	for _, arr := range arrs{
		randQuickSort(arr)
		//quickSort(arr)
	}

}