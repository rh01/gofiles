package main

import (
	"bytes"
	"fmt"
	"sync"
)

// unsafe concurrency example
func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buf bytes.Buffer
		for _, v := range data {
			fmt.Fprintf(&buf, "%c", v)
		}
		fmt.Println(buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}
