package main

import (
	"reflect"
	"fmt"
)

func main() {

	var i int = 10
	// interface to Value
	iVal := reflect.ValueOf(i)
	fmt.Println(iVal)

	iType := reflect.TypeOf(&i)
	fmt.Println(iType)

	// Value to interface
	iInterface := iVal.Interface()
	fmt.Println(iInterface)

	// interface to int
	iOrig := iInterface.(int)
	fmt.Println(iOrig)

	// kind
	iValKind := iVal.Kind()
	fmt.Println(iValKind)

	// change value
	//iVal.Elem().SetInt(100)
	//fmt.Println(i)
}
