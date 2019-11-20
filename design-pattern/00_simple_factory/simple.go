// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplefactory

import "fmt"

// API is a interface
type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
