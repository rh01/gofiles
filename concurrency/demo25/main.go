package main

import (
	"fmt"
	"sync"
)

// sync.pool
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	myPool.Get() // if no, create it using New ，调用new
	instance := myPool.Get() /// 将先前检索到的实例放在池中， 这就增加了实例的可用数量。 调用new
	myPool.Put(instance)
	myPool.Get() // new  不会调用
}


// 为什么要使用 Pool ，而不只是在运行时实例化对象呢？ 
// Go 语言是有 GC 的，因此实例化的对象将被自动清理。