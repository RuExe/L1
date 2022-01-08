package main

import (
	"fmt"
	"reflect"
)

func main() {
	var_str := "string"
	var_int := 10
	var_bool := true
	var_ch := make(chan int)

	fmt.Println(ReflectDeterminateType(var_str))
	fmt.Println(ReflectDeterminateType(var_int))
	fmt.Println(ReflectDeterminateType(var_bool))
	fmt.Println(ReflectDeterminateType(var_ch))
}

func ReflectDeterminateType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}
