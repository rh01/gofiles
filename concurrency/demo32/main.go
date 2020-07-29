package main

import (
	"fmt"
	"time"
)

// goroutine协作
func main() {
	// done chan用来子goroutine和main goroutine进行通信
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		// goroutine
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			select {
			case s := <-strings:
				fmt.Println(s)
			case <-done: // 检查我们的 done channel 是否已经发出信号
				return
			}
		}()
		return terminated
	}

	var done = make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// 在1秒后取消本操作
		time.Sleep(time.Second)
		fmt.Println("Canceling doWork goroutine ...")
		close(done)
	}()

	<-terminated // 这就是我们加入从 main goroutine 的 doWork 中产生的 goroutine 的地方。
	fmt.Println(" Done")

}
