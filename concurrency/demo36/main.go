package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkState := func(done <-chan interface{}, urls ...string) <-chan Result {
		var results = make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()

		return results
	}

	var done = make(chan interface{})
	defer close(done)

	urls := []string{"http://www.google.com", "a", "b", "c"}
	var errCount int
	for result := range checkState(done, urls...) {
		if result.Error != nil {
			fmt.Printf("err: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
