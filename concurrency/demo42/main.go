package main

import (
	"fmt"
	"math/rand"
	"time"
)

// pipline版本，下一步进行优化
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

	_ = func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return //
					case valueStream <- v:
					}
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

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	rand := func() interface{} {
		return rand.Intn(50000000)
	}

	primeFinder := func(done <-chan interface{}, valueStream <-chan int) <-chan interface{} {
		// var primeStream chan int
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for v := range valueStream {
				flag := true
				for i := 2; i <= v/2; i++ {
					if v%i == 0 {
						flag = false
						break
					}
				}
				if flag {
					select {
					case <-done:
						return
					case primeStream <- v:
					}
				}
			}
		}()
		return primeStream
	}

	// primeFinder := func(done <-chan interface{},
	// 	intStream <-chan int,
	// ) <-chan interface{} {
	// 	primeStream := make(chan interface{})
	// 	go func() {
	// 		defer close(primeStream)
	// 		for integer := range intStream {
	// 			integer -= 1
	// 			prime := true
	// 			for divisor := integer - 1; divisor > 1; divisor-- {
	// 				if integer%divisor == 0 {
	// 					prime = false
	// 					break
	// 				}
	// 			}

	// 			if prime {
	// 				select {
	// 				case <-done:
	// 					return
	// 				case primeStream <- integer:
	// 				}
	// 			}
	// 		}
	// 	}()
	// 	return primeStream
	// }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	randIntStream := toInt(done, repeatFn(done, rand))

	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
