package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {

	type value struct {
		mu sync.Mutex
		value int
	}

	var (
		wg sync.WaitGroup
	)

	processTwoNumber := func(a, b *value) {
		defer wg.Done()
		
		a.mu.Lock()
		defer a.mu.Unlock()
		time.Sleep(2 * time.Second)

		b.mu.Lock()
		defer b.mu.Unlock()

		fmt.Printf("sum=%v\n", a.value + b.value)
	}

	var a, b value
	wg.Add(2)
	go processTwoNumber(&a, &b)
	go processTwoNumber(&b, &a)

	wg.Wait()
}