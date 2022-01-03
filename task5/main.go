package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	duration := 10 * time.Second

	intCh := make(chan int)
	go func() {
		for {
			intCh <- rand.Int()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for v := range intCh {
			fmt.Println(v)
		}
	}()

	time.Sleep(duration)
}
