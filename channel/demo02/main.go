package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 当前本地机器的逻辑CPU个数
	// 默认情况下会使用全部的CPU
	// 如果你想使用指定的CPU，可以使用runtime包下的GOMAXPROCS方法
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)

	// 指定当前运行时最大可以使用的CPU个数
	// 若 n < 1，它就不会更改当前设置。
	// 本地机器的逻辑CPU数可通过 NumCPU 查询。
	// 本函数在调度程序优化后会去掉。
	var pre int
	if pre = runtime.GOMAXPROCS(2); pre < 1 {
		fmt.Println("use pre config")
	}
	fmt.Println(pre)
}
