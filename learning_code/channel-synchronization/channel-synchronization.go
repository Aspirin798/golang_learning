package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- false
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
}
