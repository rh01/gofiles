// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplefactory

import "testing"

// TestType1 test get hiApi with factory
func TestType1(t *testing.T) {
	api := NewAPI(1)
	a := api.Say("xiaoming")
	if a != "Hi xiaoming" {
		t.Fatal("Type1 test failed")
	}
}

func TestType2(t *testing.T)  {
	api := NewAPI(2)
	a:=api.Say("xiaoming")
	if a != "Hello xiaoming" {
		t.Fatal("TestType2 failed")
	}
}


// go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 
// NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。

// 在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。
