package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	mas := [...]int{1, 2, 3, -5, 10, 32, -51}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, v := range mas {
			ch1 <- v
		}
	}()

	go func() {
		for v := range ch1 {
			ch2 <- v * v
		}
	}()

	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
