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

func main() {
	u := User{
		Id:   0,
		Name: "OK",
		Age:  2,
	}

	Set(&u)
	fmt.Println(u)
}

func Set(user *User) {
	v := reflect.ValueOf(user)

	// 判断当前的user是否被设置
	// 修改属性的内容时，要求结构体中属性名首字母大写才可以设置
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	}else{
		v = v.Elem()
	}

	f := v.FieldByName("Name1")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}