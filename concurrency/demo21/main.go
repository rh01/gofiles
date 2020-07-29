package main

import (
	"fmt"
	"sync"
)

// sync package once usage
func main() {
	var count int
	increment := func() {
		count++
	}
	// 安全,只调用一次
	var once sync.Once  // 定义类型为sync,Once的实力

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("count: %v\n", count)
}

// search sync.Once freq in standard library
// grep -ir sync.Once $(go env GOROOT)/src | wc -l