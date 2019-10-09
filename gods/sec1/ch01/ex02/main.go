// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//这一节学习Tuple的使用
//1 A tuple is a finite sorted list of elements.
//2 Tuples are typically immutable sequential collections
//3 The element has related fields of different datatypes.
//4 The only way to modify a tuple is to change the fields.
//
//Operators such as + and * can be applied to tuples.
//A database record is referred to as a tuple.

///main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

import "fmt"

func main() {
	square, cube := powerSeries(2)
	fmt.Println("square is,", square, ", cube is", cube)
}

// return tuple
// func powerSeries(a int) (int, int) {
// 	return a * a, a * a * a
// }

func powerSeries(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}
