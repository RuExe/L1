package classes

import "fmt"

type Human struct {
	firstName, secondName string
	age                   int
}

func NewHuman(firstName, secondName string, age int) Human {
	return Human{
		firstName:  firstName,
		secondName: secondName,
		age:        age,
	}
}

func (h Human) FullName() string {
	return fmt.Sprintf("%s %s", h.firstName, h.secondName)
}
