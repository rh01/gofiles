package main

import (
	"fmt"
)

// 放入1-2000到numChan
func putNum(numChan chan int) {
	for i := 1; i < 2000; i++ {
		// 放数字
		numChan <- i
	}
	// 关闭通道
	close(numChan)
}

// 计算1+2+...+n的值
func computeSum(numChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		num, ok := <-numChan;
		if !ok {
			break
		}
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		// 放入数据
		resChan <- map[int]int{num:sum}
	}

	// 写入exitChan true
	exitChan <- true
	fmt.Println("有一个 computeSum 协程退出")
}

func main() {
	numChan := make(chan int, 1000)
	resChan := make(chan map[int]int, 2000)
	exitChan := make(chan bool, 8)

	// 第一个协程放入数据
	go putNum(numChan)

	// 开启8个协程取出numChan的数，并计算1+..+n的值，并放到resChan
	for i := 1; i <= 8; i++ {
		go computeSum(numChan, resChan, exitChan)
	}

	go func() {
		for i := 1; i <= 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()
	// 此时便利我们的primNum，把结果取出来
	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		// 将结果输出
		for key, value := range res {
			fmt.Printf("1+..+%d=%d\n", key, value)
		}

	}
	fmt.Println("main线程退出")
}
