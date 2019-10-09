// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// list 是 golang 标准包的 container/list 包所提供的
	// 该 list 是一个双端队列，可以存放任意类型的数据，因为
	// Elem 是 interface{}类型的，因此我们在遍历的时候，
	// 需要进行类型断言。
	var li list.List

	li.PushBack("112")
	li.PushBack(23)
	li.PushBack([]int{1, 3})

	// Element 有两个方法分别是next和prev，这里可以将Element看作
	// 双端队列的指针，多指针链表也可以，一般情况使用Element的指针
	// 方法来遍历
	for i := li.Front(); i != nil; i = i.Next() {
		// 这是一种方法
		// fmt.Println(i.Value)

		// 使用 switch 进行类型断言
		switch i.Value.(type) {
		case int:
			value := i.Value.(int)
			fmt.Println("int value is", value)
		case string:
			value := i.Value.(string)
			fmt.Println("string value is", value)
		case []int:
			value := i.Value.([]int)
			fmt.Println("int slice value is", value)
		default:
			fmt.Println("no such type fitted")
		}
	}
}
