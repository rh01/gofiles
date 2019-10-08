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

func main() {
	x := 123
	// 返回值为reflect.Value类型，即Pointer.Interface
	v := reflect.ValueOf(&x)

	// reflect.Struct
	// 若要修改，需要传指针给relect.ValueOf
	v.Elem().SetInt(999)

	fmt.Println(x)
}
