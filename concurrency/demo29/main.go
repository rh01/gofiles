package main

import "fmt"

// Command safty concurrency model
//////////////////////////////// 
func main() {
	chanOwn := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Recieved: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwn()
	consumer(results)
}
