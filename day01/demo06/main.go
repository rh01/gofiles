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
// Append returns the updated slice. It is therefore necessary to store the result of append
// code snippt: https://www.flysnow.org/2018/12/21/golang-sliceheader.html

type Slice []int

func (A Slice)Append(value int) {
	A1 := append(A, value)

	sh:=(*reflect.SliceHeader)(unsafe.Pointer(&A))
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n",sh.Data,sh.Len,sh.Cap)

	sh1:=(*reflect.SliceHeader)(unsafe.Pointer(&A1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n",sh1.Data,sh1.Len,sh1.Cap)
}

func main() {
	mSlice := make(Slice, 10, 10)
	mSlice.Append(5)
	fmt.Println(mSlice)
}