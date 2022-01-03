package main

import "fmt"

func main() {
	//удаляем 4-й элемент
	var n = 3

	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	fmt.Println(removeSlow(users, n))
	fmt.Println(users)
}

func removeSlow(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func removeSlow2(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeFast(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
