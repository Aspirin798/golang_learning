package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "MESSAGE"
		/* 		time.Sleep(1 * time.Second) */
		signals <- "SIGNALS"
	}()
	time.Sleep(3 * time.Second)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("recevied message", msg)
	case sig := <-signals:
		fmt.Println("recevied signal", sig)
	default:
		fmt.Println("no activity")
	}
}
