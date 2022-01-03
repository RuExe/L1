package main

import "fmt"

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := GroupBy(array, StepSelector(10))
	fmt.Println(res)
}

func StepSelector(step int) func(float64) int {
	return func(v float64) int {
		return int(v) / step * step
	}
}

func GroupBy(slice []float64, selector func(float64) int) map[int][]float64 {
	res := make(map[int][]float64)

	for _, v := range slice {
		key := selector(v)
		res[key] = append(res[key], v)
	}
	return res
}
