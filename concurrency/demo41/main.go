package main

import "fmt"

func main() {
	_ = func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
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

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
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

	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		var stringStream = make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	// rand := func() interface{} {
	// 	return rand.Int()
	// }

	// for num := range take(done, repeatFn(done, rand), 10) {
	// 	// n := num.(int)
	// 	fmt.Println(num)
	// }
	var msg string
	for s := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		msg += s
	}
	fmt.Printf("message: %s...", msg)
}
