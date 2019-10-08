package main

import "fmt"

// channel 遍历
// 通道若没有关闭，则会出现deadlock panic
// 通道若关闭，正常遍历数据，便完成后，就会退出遍历
func main() {
	intChans := make(chan int, 3)
	intChans <- 10
	intChans <- 20
	intChans <- 30
	// required!
	close(intChans)

	// if channel not close, then
	// fatal error: all goroutines are asleep - deadlock!
	for value := range intChans {
		fmt.Println(value)
	}
}
