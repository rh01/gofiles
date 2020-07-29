package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subsribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup // 第一个wg，用来控制函数内部的goroutine
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()

		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // 第二个wg，用来控制程序在写入stdout执勤不会推出
	clickRegistered.Add(3)
	subsribe(button.Clicked, func() {
		fmt.Println("Maxmizing windows")
		clickRegistered.Done()
	})

	subsribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subsribe(button.Clicked, func() {
		fmt.Println("Mouse moved")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()  // 通知所有注册的处理程序
	clickRegistered.Wait()
}
