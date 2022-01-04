package main

import (
	"L1/internal/set"
	"fmt"
)

func main() {
	strings1 := []string{"cat", "dog", "tree", "door", "bird"}
	strings2 := []string{"cat", "dog", "fish", "test", "door"}

	strings1Set := set.New(strings1)
	strings2Set := set.New(strings2)
	intersection := strings1Set.Intersection(&strings2Set)
	fmt.Println(intersection.ToSlice())
}
