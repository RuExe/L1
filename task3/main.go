package main

import (
	"fmt"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	fmt.Println(SquareSum(data...))
}

func SquareSum(numbers ...int) int {
	intCh := make(chan int)
	defer close(intCh)

	for _, number := range numbers {
		go func(n int) {
			intCh <- n * n
		}(number)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-intCh
	}
	return sum
}
