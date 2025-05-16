package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("hello world")
	for i := range 1_000_000 {
		time.Sleep(time.Second * 10)
		fmt.Println("index ", i)
	}
}
