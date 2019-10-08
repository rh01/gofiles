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
	ch1 := make(chan int, 3)
	// ch1 := make(chan int, 10) // 10表示channel的容量
	// 容量表示当前通道最多可以缓冲多少元素值
	// 因此带有容量的被称缓冲通道，否则为非缓冲通道

	// 发送
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	// 接收
	elems := <-ch1
	fmt.Printf("The first element received from channel ch1: %v.\n", elems)

}
