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
	sli1 := make([]string, 2) // 切片类型的初始化方法
	// 使用该初始化方法后，我们默认情况下已经为开辟了大小为2的数组空间

	//默认情况下，上面的初始化只是初始化了len值，即当前切片的大小，
	// 并没有指出切片的具体容量,那么此时容量值与切片大小值相等，但是我们都知道，
	// 切片可以看做是数组类型的封装，它的实现是由底层数组来实现的，那么未指明切片容量
	// 就说明切片数组是动态可变的
	println(cap(sli1))
	println(len(sli1))

	// 切片的方法，赋值
	sli1[0] = "a"
	// sli1 = append(sli1, 	// sli1 = append(sli1, "a")
	// sli1 = append(sli1, []string{"a"}...) // 这里我们使用了重新赋值，说明引用
	println(sli1)      // 此时打印的是引用，即地址，地址与你当前计算机有管
	println(len(sli1)) // 此时会发现
	println(cap(sli1))
	for _, value := range sli1 {
		fmt.Println(value)
	}
	/**
	  规律： 刚开始，每次增长超过size后，cap就会以2倍的大小增长，
	在达到1024时就会按照1.x倍增长
	*/

	// 指明了cap
	sli2 := make([]int, 3, 5)
	// 扩容：上述就是扩容的实现。主要需要关注的有两点，一个是扩容时候的策略，还有一个就是扩容是生成全新的内存地址还是在原来的地址后追加。
	// Append returns the updated slice. It is therefore necessary to store the result of append
	sli10 := append(sli2, []int{2, 2, 3}...) // 返回一个新的slice
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))

	fmt.Println(len(sli10))
	fmt.Println(cap(sli10))

	//for _, value := range sli2 {
	//	fmt.Println(value)
	//}

	//sli3 := sli2[6:]
	//fmt.Println(len(sli3)) // 此时的len为0，因为已经超了6
	//fmt.Println(cap(sli3)) // 此时的cap为10-6=4

	//sli4 := sli2[2:3]
	//fmt.Println(len(sli4)) // 此时的len为切分大小、长度
	//fmt.Println(cap(sli4)) // 此时的cap仍然是母序列的cap减去index即10-2
	// fmt.Println(sli4[3])   //此时panic异常，因为当前的窗口只能看到当前的切片范围内的

	//fmt.Println(sli4) // 只能看到当前窗口的内容，切片的容量可以看做透过这个窗口最多可以看到底层数组中元素的个数

}
