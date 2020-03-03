package main

import (
	"fmt"
	"sync"
)

func main() {
	// Once 对象表示只执行一次的对象
	var once sync.Once

	onceBody := func() {
		fmt.Println("Only Once")
	}

	done := make(chan bool)
	for i:= 0; i< 10; i++ {
		go func() {
			// 当且仅当Do第一次被调用的时候，onceBody才执行
			// 如果once.Do(f)被多次调用，只有第一次调用会执行f，
			// 即使f每次调用Do 提供的f值不同。
			// 需要给每个要执行仅一次的函数都建立一个Once类型的实例。
			once.Do(onceBody)
			done <- true
		}()
	}
	// summary：Do用于必须刚好运行一次的初始化。
	// 因为f是没有参数的，因此可能需要使用闭包来提供给Do方法调用：
	for i:= 0; i < 10; i++ {
		<-done
	}
}
