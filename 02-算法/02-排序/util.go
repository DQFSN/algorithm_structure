package sort

import (
	"math/rand"
)

type ints []int

func (i ints) Iter() func() (value int, ok bool) {
	index := 0

	return func() (value int, ok bool) {
		if index >= len(i) {
			return
		}

		value, ok = i[index], true
		index++

		return
	}
}

func getTestData() [][]int {
	const TESTTIMES = 10
	const DATASIZE = 100
	arrs := make([][]int, TESTTIMES, TESTTIMES)
	for i := 0; i < TESTTIMES; i++ {
		arr := make([]int, DATASIZE, DATASIZE)
		for j := 0; j < DATASIZE; j++ {
			arr = append(arr,rand.Intn(10000))
		}
		arrs = append(arrs,arr)
	}
	return arrs
}

func checkSame(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}