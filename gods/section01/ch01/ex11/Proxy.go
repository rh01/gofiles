// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/**
The proxy pattern forwards to a real object and acts as an interface to others.
The proxy pattern controls access to an object and provides additional functionality.
The additional functionality can be related to authentication,
authorization, and providing rights of access to the resource-sensitive object.
The real object need not be modified while providing additional logic.
Remote, smart, virtual, and protection proxies are the scenarios where this pattern is applied.
It is also used to provide an alternative to extend functionality with inheritance and object composition.
A proxy object is also referred to as a surrogate, handle, or wrapper.
*/

//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//IRealObject interface
type IRealObject interface {
	performAction()
}

//RealObject struct
type RealObject struct{}

// RealObject class implements IRealObject interface. The class has method performAction.
//RealObject class method performAction
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

//VirtualProxy struct
type VirtualProxy struct {
	realObject *RealObject
}

//VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}

// main method
func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}
