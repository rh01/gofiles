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
	"time"
)

func main() {
	// map 不是并发安全的
	//sync.Map{} 并发安全
	m := map[int]string{
		1: "haha",
	}

	go read(m)
	time.Sleep(time.Second)

	go write(m)
	time.Sleep(30*time.Second)

	fmt.Println(m)
}

func write(m map[int]string) {
	for   {
		m[1] = "write"
		time.Sleep(1)
	}
}

func read(m map[int]string) {
	for{
		_ = m[1]
		time.Sleep(1)
	}
}
