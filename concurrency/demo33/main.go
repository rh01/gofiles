package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandomStream closure  exited")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}

			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("2 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(1 * time.Second)
}
