package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 这里循环的接受客户端发送的数据
	defer conn.Close()

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)

		//1等待客户端通过conn发送消息
		//2如果没有write发送，那么协程将会阻塞在这里
		fmt.Printf("服务器在等待客户端%s发送消息.\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出")

			return
		}
		//3显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}

}

func main() {
	fmt.Println("服务器开始监听了")
	//1. tcp 表示使用网络协议是tcp
	//2. 127.0.0.1：8888 表示监听本地接口8888
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	// 使用defer延时关闭
	defer listener.Close()
	for {
		// 等待客户端来连接
		//conn 实现了Conn接口的
		fmt.Printf("等待客户端来连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept suc con=%v.\n", conn)
		}
		// 这里准备起来一个协程，为客户端服务
		go process(conn)
	}
	fmt.Printf("listen suc=%v.\n", listener)
}
