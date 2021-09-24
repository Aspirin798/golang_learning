package main

import (
	"fmt"
	"time"
)

func worke(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	fmt.Println("done")
	time.Sleep(3 * time.Second)
	done <- false

}

func main() {

	done := make(chan bool, 1)
	go worke(done)

	<-done
}
