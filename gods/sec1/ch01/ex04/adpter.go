// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

//Structural design patterns describe the relationships between the entities.
//They are used to form large structures using classes and objects.
//These patterns are used to create a system with different system blocks in a flexible manner.
//[Adapter], [bridge], [composite], [decorator], [facade], [flyweight], [private class data], and [proxy]
//are the Gang of Four (GoF) structural design patterns.
//The private class data design pattern is the other design pattern covered in this section.

//The adapter pattern provides a wrapper with an interface required by the API client
//to link incompatible types and act as a translator between the two types.
//The adapter uses the interface of a class to be a class with another compatible
//interface. When requirements change,
//there are scenarios where class functionality needs to be changed because of incompatible interfaces.
//适配器主要就是需要实现接口定义的方法就是实现了该接口，
//因此可以称为适配器模式
//The dependency inversion principle can be adhered to by using the adapter pattern,
//when a class defines its own interface to the next level module interface implemented
//by an adapter class.
//Delegation is the other principle used by the adapter pattern.
//Multiple formats handling source-to-destination transformations are the scenarios
//where the adapter pattern is applied.
//The adapter pattern comprises the target, adaptee, adapter, and client:
//Adapter class method process

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

type Adapter struct {
	adaptee Adaptee
}

//Adaptee Struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert() {
	adaptee.adapterType = 10
	fmt.Println("Adaptee convert method", adaptee.adapterType)
}

type IProcess interface {
	process()
}

// main method
func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
