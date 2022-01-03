package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	workersCount := 5

	intCh := DataGenerator()
	CreatePrintWorkers(intCh, workersCount)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	done := make(chan bool)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		fmt.Println("Какие-то действия которые необходимы для корректного завершения программы")
		close(intCh)
		done <- true
	}()

	fmt.Println("Ожидание сигнала")
	<-done
	fmt.Println("Выход из программы")
}

func DataGenerator() chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- rand.Int():
				time.Sleep(time.Second)
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func CreatePrintWorkers(ch <-chan int, count int) {
	for i := 0; i < count; i++ {
		go func() {
			for v := range ch {
				fmt.Println(v)
			}
		}()
	}
}
