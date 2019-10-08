package main

import (
	"fmt"
)

// 声明或者定义一个接口
type Usb interface {
	// 声明两个没有实现的方法
	start()
	stop()
}

type Phone struct {
}

func (p *Phone) start() {
	fmt.Println("手机开始工作")
}

func (p *Phone) stop() {
	fmt.Println("手机停止工作")
}

// 计算机
type Computer struct {
}

// 编写一个方法working方法，接受一个Usb接口类型的变量
// 只要是实现了Usb接口（所谓实现usb接口就是指
// 实现了Usb接口声明的所有方法）
func (c *Computer) Working(usb Usb) {
	// 通过usb接口变量来调用start 和stop方法
	usb.start()
	usb.stop()
}

func main() {
	phone := &Phone{}
	c := &Computer{}
	c.Working(phone)
}
