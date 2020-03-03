// Copyright 2019 The darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func main() {
	// mutex
	// Mutex是一个互斥锁，可以创建为其他结构体的字段；
	// 零值为解锁状态。
	// Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。
	// Lock: Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。
	// Unlock方法解锁m，如果m未加锁会导致运行时错误。锁和线程无关，可以由不同的线程加锁和解锁。

	// RWMutex是读写互斥锁。
	// 	该锁可以被同时多个读取者持有或唯一个写入者持有。
	// 	RWMutex可以创建为其他结构体的字段；零值为解锁状态。
	// 	RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
	// writer.

}
