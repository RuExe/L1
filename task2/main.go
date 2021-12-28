package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	PrintSquares(numbers...)
}

func PrintSquares(numbers ...int) {
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, number := range numbers {
		go func(n int) {
			defer wg.Done()
			fmt.Println(Square(n))
		}(number)
	}
	wg.Wait()
}

func Square(number int) int {
	return number * number
}
