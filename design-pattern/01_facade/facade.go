// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package facade 是外观模式实现,大部分代码使用此接口简化对facade类的访问。
// facade模块同时暴露了a和b 两个Module 的NewXXX和interface，
// 其它代码如果需要使用细节功能时可以直接调用。
package facade

import "fmt"

func NewAPI() API {
	return &APIImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

//API is facade interface of facade package
type API interface {
	Test() string
}

// facade implement
type APIImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *APIImpl) Test() string {
	aout := a.a.TestA()
	bout := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aout, bout)
}

//---------------------------------
func NewAModuleAPI() AModuleAPI {
	return &aModuleAPIImpl{}
}

type AModuleAPI interface {
	TestA() string
}

type aModuleAPIImpl struct {
}

func (a *aModuleAPIImpl) TestA() string {
	return "A module runing"
}

//---------------------------------

func NewBModuleAPI() BModuleAPI {
	return &bModuleAPIImpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleAPIImpl struct {
}

func (b *bModuleAPIImpl) TestB() string {
	return "B runing...."
}
