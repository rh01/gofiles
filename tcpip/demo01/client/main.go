package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	//1. tcp
	//2. 连接本地的8888端口
	//返回一个Conn对象和一个err
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	fmt.Printf("Dial conn=%v.\n", conn)
	fmt.Printf("Dial local addr=%v.\n", conn.LocalAddr())

	// conn.
	//功能一：客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin)

	// 从终端读取一行用户输入，并准备发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("readString err=%v\n", err)
	}

	// 将line发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Printf("Send or Write err:", err)
		return
	}
	fmt.Printf("关闭连接，发送成功，发送了%d个字节的数据.\n", n)
}
