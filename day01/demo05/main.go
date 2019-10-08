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

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//s := make([]byte, 200)
	//ptr := unsafe.Pointer(&s[0])
	//fmt.Printf("%v.\n", ptr)

	// 1. 从内存地址中构造一个slice
	var ptr unsafe.Pointer
	length := 3

	// 这个也就是参考了runtime/slice.go中的slice的定义实现的
	var s1 = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{ptr, length, length}

	s := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Println(cap(s))

	// 2. 在GO的反射中存在一个阈值对应的数据结构SliceHeader，可以
	// 使用SliceHeader来构造slice
	// SliceHeader是Slice运行时的具体表现，它的结构定义如下：
	// type SliceHeader struct {
	//	Data uintptr
	//	Len  int
	//	Cap  int
	//}
	// 正好对应Slice的三要素，Data指向具体的底层数据源数组，Len代表长度，Cap代表容量。
	var o []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&o))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(ptr)
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n",sliceHeader.Data,sliceHeader.Len,sliceHeader.Cap)


}
