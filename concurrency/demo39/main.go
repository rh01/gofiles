package main

import "fmt"

func main() {
	repeat := func(done <-chan interface{},
		values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	// num is stop conndition
	take := func(
		done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				v := <-valueStream
				select {
				case <-done:
					return //
				case takeStream <- v:
				}
			}
		}()
		return takeStream
	}

	var done = make(chan interface{})
	defer close(done)

	takeStream := take(done, repeat(done, 1), 10)
	// defer close(takeStream)
	// n := num
	for {
		num, ok := <-takeStream
		if !ok {
			break
		}
		fmt.Printf("%v ", num)
	}
}
