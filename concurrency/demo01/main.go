package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("hello world")
	// done, decreament the counter
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// add, increment the counter
	wg.Add(1)
	go sayHello(&wg)
	// wait, block until the counter == 0
	wg.Wait()
	fmt.Println("Main goroutine")
}
