package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salution := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salution = "welcome"
	}()
	wg.Wait()
	fmt.Println(salution)
}
