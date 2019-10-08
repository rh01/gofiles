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
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func main() {
	// 嵌入字段，类型判断
	m := Manager{
		User:  User{1, "OK", 12},
		title: "123",
	}
	// 利用反射包取出匿名字段
	v := reflect.ValueOf(m)
	// t := reflect.TypeOf(m)

	// 取出User字段
	fmt.Printf("%#v.\n", v.Field(0))
	// 那么如何取出User中的Id字段呢？
	fmt.Printf("%#v.\n", v.FieldByIndex([]int{0, 0})) // 第一个0为User字段，第二个0位User字段中的第一个字段

	// 怎么通过反射来修改某个例子

}
