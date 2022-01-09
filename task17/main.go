package main

import (
	"fmt"
)

func main() {
	searched := 85
	a := []int{1, 2, 6, 8, 9, 10, 14, 32, 51, 62, 85, 103}
	fmt.Println(binarySearch(a, searched))
}

func binarySearch(slice []int, searched int) int {
	left, right := 0, len(slice)-1
	for left <= right {
		middle := (left + right) / 2
		if slice[middle] < searched {
			left = middle + 1
		} else if slice[middle] > searched {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
