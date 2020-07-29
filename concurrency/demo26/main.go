package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int
	calcPool := sync.Pool{
		New: func() interface{} {
			numCalcsCreated+=1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// 使用4kb实力话pool
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i> 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte) // 断言
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created." , numCalcsCreated)
}


// 为什么要使用 Pool ，而不只是在运行时实例化对象呢？ 
// Go 语言是有 GC 的，因此实例化的对象将被自动清理。