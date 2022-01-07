package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcd"
	fmt.Println(charactersInTheStringAreUnique(str))
	str = "abCdefAaf"
	fmt.Println(charactersInTheStringAreUnique(str))
	str = "aabcd"
	fmt.Println(charactersInTheStringAreUnique(str))
	str = "aA"
	fmt.Println(charactersInTheStringAreUnique(str))
}

func charactersInTheStringAreUnique(value string) bool {
	str := strings.ToLower(value)

	chars := make(map[string]bool, len(str))
	for _, c := range str {
		key := string(c)
		if _, ok := chars[key]; ok {
			return false
		}
		chars[key] = true
	}
	return true
}
