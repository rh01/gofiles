package main

import (
	"fmt"
	"net/http"
)

func main() {

	// wrap response into result struct
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkState := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
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
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkState(done, urls...) {
		if result.Error != nil {
			fmt.Printf("err: %v\n", result.Error)
			continue
		}
		fmt.Printf("response: %v\n", result.Response.Status)
	}
}
