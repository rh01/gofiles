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

func (u User) SayHello() {
	fmt.Println("Hello World")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	///// 判断类型，是否使我们想要的类型
	if k := v.Kind(); k != reflect.Struct {
		fmt.Println("xxx")
		return
	}


	v := reflect.ValueOf(o)
	fmt.Println("Field")



	// 获取字段信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	// 获取方法的信息
	fmt.Println("Methods")
	for j := 0; j < t.NumMethod(); j++ {
		m := t.Method(j)
		// val := v.Method(j).Interface()
		fmt.Printf("%6s: %v", m.Name, m.Type)
	}

	// 如果此时的类型为指针类型，将不会取出、正常运行
}

func main() {
	u := User{
		Id:   1,
		Name: "Ok",
		Age:  12,
	}

	// Info(&u) // panic
	Info(u) // panic
}
