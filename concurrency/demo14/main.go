package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memC := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } //永远不会退出的goroutine

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memC()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memC()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
