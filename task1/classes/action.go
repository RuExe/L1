package classes

import "fmt"

type Action struct {
	Human
}

func (a Action) Say(word string) string {
	return fmt.Sprintf("%s say: %s", a.FullName(), word)
}
