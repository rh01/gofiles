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

type MyString = string // 类型的别名类型
type MyString2 string  // 类型的再定义



func main() {
	var a MyString
	a = "hha"
	fmt.Println(a)

	// 天然合法
	as := []MyString{"aaa", "bbb", "cccc"}
	as2string := []string(as)
	fmt.Println(as2string)

	// 合法的操作
	var b MyString2
	b = "hha"
	b2str := string(b)
	fmt.Println(b2str)

	// 对于集合类类型转换这是不合法的操作
	//c := []MyString2{"hhh", "aaa", "bbb"}
	//strings := []string(c) // 编译错误

}
