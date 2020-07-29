package main

import (
	"sync"
)

func main() {
	// var count int
	// increment := func() {count++}
	// decrement := func() {count--}

	var onceA, onceB sync.Once
	var initB func()
	// var wg sync.WaitGroup
	// wg.Add(2)
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	// wg.Wait()
	// wg.Wait()

	onceA.Do(initA)
	// fmt.Printf(" count: %v", count)
}
