package main

import (
	"fmt"
	"time"
)

// 每隔一秒钟打印一个Hello World
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

// 捕获协程的panic，因为在实际开发中，有可能一个协程
// 因为panic而导致整个程序崩溃
func test() {
	// 这里可以使用defer+recover
	defer func() {
		// 捕获test抛出的panic
		if err := recover(); err!= nil{
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	// panic: assignment to entry in nil map
	myMap[0] = "golang"

}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("OK")
	}
}
