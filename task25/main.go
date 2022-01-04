package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start wait")
	sleep(time.Second * 3)
	fmt.Println("End")
}

func sleep(d time.Duration) {
	<-time.After(d)
}
