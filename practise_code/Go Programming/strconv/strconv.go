package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* 	v := "20"
	   	fmt.Printf("%T,%v\n", v, v)
	   	s, err := strconv.Atoi(v)
	   	if err == nil {

	   		fmt.Printf("%T,%v\n", s, s)
	   		return
	   	}
	   	fmt.Println("err:", err) */

	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

}
