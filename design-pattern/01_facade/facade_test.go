// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package facade 是外观模式实现,大部分代码使用此接口简化对facade类的访问。
// facade模块同时暴露了a和b 两个Module 的NewXXX和interface，
// 其它代码如果需要使用细节功能时可以直接调用。
package facade

import "testing"

const expect = "A module runing\nB runing...."

func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}
