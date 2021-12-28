package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	go PrintSquares(numbers...)
	time.Sleep(time.Second * 3)
}

func PrintSquares(numbers ...int) {
	fmt.Println(Squares(numbers...))
}

// Я не совсем понял что от меня хотят в задании. Написать конкурентно или вывести именно с помощью интерфейса Stdout, поэтому вот 2 вариант
func PrintSquares2(numbers ...int) {
	for _, v := range Squares(numbers...) {
		os.Stdout.WriteString(strconv.Itoa(v))
		os.Stdout.Write([]byte{'\n'})
	}
}

func Squares(numbers ...int) []int {
	squares := make([]int, len(numbers))
	for i, number := range numbers {
		squares[i] = Square(number)
	}
	return squares
}

func Square(number int) int {
	return number * number
}
