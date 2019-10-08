package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		// 放数据
		intChan <- i
	}
	// 关闭通道
	close(intChan)
}

func isPrime(intChan, resChan chan int, exitChan chan bool) {
	// for loop
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		// 如果是素数放入primeChan
		if flag {
			resChan <- num
		}
	}
	time.Sleep(time.Millisecond * 10)
	fmt.Println("有一个primeNum 协程因为取不出数据，退出")

	// 计算完成后写入
	exitChan <- true
}

func main() {
	// 初始化3个通道
	intChan := make(chan int, 1000)
	resChan := make(chan int, 80000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()
	go putNum(intChan)
	// for loop
	// 开启4个协程，从 intChan 去除数据，并判断是否为素数
	for i := 0; i < 4; i++ {
		go isPrime(intChan, resChan, exitChan)
	}

	// 看看还能不能读出数据
	go func() {
		for i := 0; i < 4; i++ {
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
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
	end := time.Now().Unix()
	fmt.Println( end-start)
}
