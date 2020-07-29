package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	go func() {
		i++
	}()
	time.Sleep(time.Millisecond)
	if i == 1 {
		fmt.Printf("%v is a number\n", i)
	}
}
