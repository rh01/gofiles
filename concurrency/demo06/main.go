package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		memoryLock sync.Mutex
		i          int
	)

	go func() {
		memoryLock.Lock()
		i++
		memoryLock.Unlock()
	}()

	memoryLock.Lock()
	if i == 0 {
		fmt.Printf("%v is 0\n", i)
	} else {
		fmt.Printf("%v is 1\n", i)
	}
	memoryLock.Unlock()
}
