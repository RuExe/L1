package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		hash  = make(map[int]int)
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			mutex.Lock()
			for j := 0; j < 5; j++ {
				hash[index]++
			}
			defer mutex.Unlock()
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(hash)
}
