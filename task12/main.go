package main

import (
	"L1/internal/set"
	"fmt"
)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree", "dog", "cat", "tree"}
	stringSet := set.New(strings)
	fmt.Println(stringSet.ToSlice())
}
