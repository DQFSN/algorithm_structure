package main

import (
	"fmt"
)

func main() {
	/*
	A := [][]int{{1,2,3},{4,5,6}}
	B := transpose(A)

	for _, r := range B {
		for _, v := range r {
			fmt.Printf("%d ",v)
		} 
		fmt.Println()
	}
	*/


	A := []int{1,2,3,4,5,5}
	fmt.Println(majorityElement2(A))
}