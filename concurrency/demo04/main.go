package main

import (
	"fmt"
	"sync/atomic"
)

// count 作为一种信号，来协调打印的序号
var count uint32 = 0

func main() {

	//  trigger函数会不断的获取一个名为count的变量，并判断该变量
	// 是否与 i 相同，如果相同则调用fn函数，并将count加1
	// 最后退出循环
	trigger := func(fn func(), i uint32) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				// 原子操作
				atomic.AddUint32(&count, 1)
				break
			}
			// time.Sleep(time.Nanosecond)
		}
	}

	var _ = [...]int{1, 2, 3, 4, 5}

	for i := uint32(0); i < 10; i++ {
		// go 语句 -> go 函数
		go func(i uint32) {
			// 匿名函数，赋值给变量fn
			fn := func() {
				fmt.Println(i)
			}
			trigger(fn, i)
		}(i)
	}
	// 主线程阻塞在这里，等待10的出现
	// 这个程序保证了顺序
	trigger(func() { fmt.Println("退出") }, 10)
}
