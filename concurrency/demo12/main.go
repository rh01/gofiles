package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salt string) {
			defer wg.Done()
			fmt.Println(salt)
		}(salutation)
	}
	wg.Wait()
}
