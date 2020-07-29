package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int 
	increment := func() {count++}
	decrement := func() {count--}

	var once sync.Once
	// var wg sync.WaitGroup
	// wg.Add(2)
	once.Do(increment)
	once.Do(decrement)
	// wg.Wait()

	fmt.Printf(" count: %v", count)
}
