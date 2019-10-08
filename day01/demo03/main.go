/*
 * Copyright 2019 @rh01 https://github.com/rh01
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import "fmt"

func main() {
	// 数组与切片
	arr1 := [2]string{"a", "b"} // 我发现如果你此时数组中的元素超过你定义的数组大小时，报编译时错误，
	fmt.Println(arr1)           // 打印
	fmt.Println(len(arr1))      // 长度
	fmt.Println(arr1[1])        // 索引

	for _, value := range arr1 {
		fmt.Println(value) // 遍历
	}

	arr1[1] = "c" // 修改
	fmt.Println(arr1[1])

	value, ok := interface{}(arr1).([2]string) // 因此[1]string和[2]string是两种类型
	// value, ok := interface{}(arr1).([]string) // 注意：切片类型与数组类型 此时的类型推断不再适用
	if ok {
		fmt.Println(value)
	}

	// 数组的初始化
	var a [2]string // method1 先声明并初始化，默认情况为""
	fmt.Println(a)
	fmt.Println(len(a))
	a[1] = "a"
	fmt.Println(a) // 第一个没有赋值

	b := [3]string{"a", "b", "c"} // method2，段变量2
	fmt.Println(b)

	// 数组
	fmt.Println(cap(b)) // cap 总是等于len值

	// copy仅使用与slice类型
	//var bb [3]string // 必须指明bb的大小，保证类型一样
	//copy(bb, b)


	// new函数返回值指向[]string的零值
	c := new([2]string) // 此时的c为指针或者引用类型
	c[1] ="a"
	fmt.Println(*c)

	print(c)


}
