package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main() {
	// 创建一个容量为3，可以存放任意类型的管道
	allChan:= make(chan interface{}, 3)

	allChan <- 10
	allChan <- "cat"
	cat := Cat{"小花猫", 20}
	allChan <- cat

	// 这时候该怎么取
	<- allChan
	<- allChan

	newCat := <- allChan
	// 编译与运行时的区别
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	// 如何取出Name字段 -> 类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
}
