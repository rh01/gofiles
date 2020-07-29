package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10) // simulation queue using slice

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)

		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()                           // 发出某个信号，通知一个等待的goroutine发生什么事
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()                        // 阻塞当前goroutine运行或等待某个信号发生
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})   // enqueue
		go removeFromQueue(1 * time.Second) // dequeue
		c.L.Unlock()
	}
}
