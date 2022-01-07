package main

import "fmt"

func main() {
	var number int64 = 1537
	fmt.Println(setBit(number, 1, true))
}

func setBit(number int64, index int8, bit bool) int64 {
	var temp int64 = 1 << index
	if bit {
		return number | temp
	}
	return number & ^temp
}
