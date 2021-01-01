package main

import (
	"fmt"
)


func transpose(A [][]int) [][]int {
    if len(A) < 1 {
        return [][]int{}
    }
	B := make([][]int, len(A[0]), len(A[0]))

	for c := 0; c < len(A[0]); c++ {
		for r := 0; r < len(A); r++ {
			B[c] = append(B[c],A[r][c])
		}
	}

	return B
}


