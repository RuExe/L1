package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		a int64 = 9223372036854775807
		b int64 = 2
	)
	fmt.Println(Sum(a, b))
	fmt.Println(Difference(a, b))
	fmt.Println(Multiply(a, b))
	fmt.Println(Divide(a, b))
}

func Sum(a, b int64) big.Int {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	res := new(big.Int)
	res.Add(bigA, bigB)
	return *res
}

func Difference(a, b int64) big.Int {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	res := new(big.Int)
	res.Sub(bigA, bigB)
	return *res
}

func Divide(a, b int64) big.Int {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	res := new(big.Int)
	res.Div(bigA, bigB)
	return *res
}

func Multiply(a, b int64) big.Int {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	res := new(big.Int)
	res.Mul(bigA, bigB)
	return *res
}
