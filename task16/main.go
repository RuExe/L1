package main

import (
	"fmt"
)

func main() {
	fmt.Println(quickSort([]int{5, 6, 7, 2, 1, 0}))
}

func quickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice[:]
	}

	pivot := slice[0]
	less := []int{}
	greater := []int{}
	for v := range slice {
		if v > pivot {
			greater = append(greater, v)
		} else {
			less = append(less, v)
		}
	}

	res := make([]int, 0, len(slice))
	res = append(res, less...)
	res = append(res, pivot)
	res = append(res, greater...)
	return res
}
