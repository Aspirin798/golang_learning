package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	f("direct")

	go f("goroutine")
	time.Sleep(time.Second)
	go f("Goroutine")
	time.Sleep(time.Second)
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)

	time.Sleep(time.Second)

	time.Sleep(time.Second)

	time.Sleep(time.Second)

	fmt.Println("Done")
}
