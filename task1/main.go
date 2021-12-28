package main

import (
	"L1/task1/classes"
	"fmt"
)

func main() {
	action := classes.Action{
		Human: classes.NewHuman("Ivan", "Ivanov", 228),
	}
	fmt.Println(action.FullName())
	fmt.Println(action.Say("Hello"))
}
