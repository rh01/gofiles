package main

import (
	"fmt"
	"sync"
)

// mutex using example
func main() {
	var count int       // shared resource, counter
	var lock sync.Mutex // mutex for controll multi-goroutine access count

	increment := func() {
		lock.Lock()
		count++
		defer lock.Unlock()
		fmt.Printf("Increment: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		count--
		defer lock.Unlock()
		fmt.Printf("Decrement: %d\n", count)
	}

	// 增量
	var arithmetic sync.WaitGroup
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// 减量
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
}
