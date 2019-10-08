package main

import "fmt"

// 向管道中intChan写入50个整数
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		//放入数据
		intChan <- i
		fmt.Printf("WriteData: %v.\n", i)
	}
	// 关闭管道
	close(intChan)
}

// 向管道intChan中读入writeData写入的数据
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData: %v\n", v)
	}
	// readData读取完数据后，级任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道，数据管道和退出管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan) // 写放入到协程中
	go readData(intChan, exitChan)

	for {
		if _, ok := <-exitChan; !ok {
			break
		}
	}
}
