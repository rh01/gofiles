package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return //
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	// num is stop conndition
	take := func(done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				// for v := range valueStream {
				v := <-valueStream // 这个需要非常注意
				select {
				case <-done:
					return //
				case takeStream <- v:
				}
				// }
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {
		return rand.Int()
	}

	for num := range take(done, repeatFn(done, rand), 10) {
		// n := num.(int)
		fmt.Println(num)
	}

}
