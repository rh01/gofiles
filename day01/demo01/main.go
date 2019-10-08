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

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{1: "One", 2: "two", 3: "three"}
	fmt.Printf("The element is %q.\n", container[1])

	// 如何在打印其中元素之前，判断container的类型
	// 类型推断，断言
	value, ok := interface{}(container).([]string)
	// 如果ok为true则value将会赋值为真实的值，否则将会为空
	// 另外ok也可以没有，若没有，如果判断为false的话，就会出现panic
	if !ok {
		fmt.Println("nil")
	} else {
		fmt.Println(value)
	}

	value2, ok := interface{}(container).(map[int]string)
	if !ok {
		fmt.Println("nil")
	} else {
		fmt.Println(value2)
	}

	u := int16(-255)
	i := int8(u)
	fmt.Println(i)

}
