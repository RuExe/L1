package main

import (
	"fmt"
	"sync"
)

//type Increment struct {
//	count int64
//}

//func (c *Increment) inc() {
//	atomic.AddInt64(&c.count, 1)
//}

type Increment struct {
	count int
	lock  sync.Mutex
}

func (i *Increment) Increase() {
	i.lock.Lock()
	i.count++
	defer i.lock.Unlock()
}

func main() {
	increment := Increment{count: 0}
	var wg sync.WaitGroup
	n := 100_000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				increment.Increase()
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(increment.count)
}
