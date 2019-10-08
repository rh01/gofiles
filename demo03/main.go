package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：实现一个计算1-200的阶乘运算，使用channel的方式来做这个事情
// 程序有以下问题：
// 1. 锁的代价太高
// 2. time.Sleep()不好控制，比如计算1-100但是你设置休眠1秒，
// 会导致程序提前退出

var (
	dataMap = make(map[int]int, 200)
	// 声明一个全局的互斥锁
	// lock 是一个全局的锁对象
	lock sync.Mutex
)

// 计算阶乘的函数
func jiecheng(n int) {

	b := 1
	for i := 1; i <= n; i++ {
		b *= i
	}

	// 这里可以将计算后的数据放入到map中
	// 问题一：如何解决使用全局锁
	// fatal error: concurrent map writes
	lock.Lock()
	dataMap[n] = b // 这个是否会出现并发写的问题？
	lock.Unlock()

}

func main() {
	for i := 1; i <= 20; i++ {
		go jiecheng(i)
	}

	// 问题2, 如果不加这个sleep会怎么样，直接退出
	time.Sleep(5*time.Nanosecond)

	// 问题3
	lock.Lock()
	for key, value := range dataMap {
		fmt.Printf("key: %3d, value: %10d.\n", key, value)
	}
	lock.Unlock()
}
