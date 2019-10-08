package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup, i int) {
	fmt.Println("hello world", i)
	// done, decreament the counter
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// add, increment the counter
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sayHello(&wg, i)
	}
	// wait, block until the counter == 0
	wg.Wait()
	// it wont't print Main routine and exit, until after
	// all the goroutine complete and done his job
	fmt.Println("Main routine")
}
