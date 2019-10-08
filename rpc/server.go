package main

import (
	"net/rpc"
	"go-learn/rpc/server"
	"net"
	"log"
	"net/http"
	"time"
)

func main() {

	//服务端调用
	arith := new(server.Arith)
	//Register在DefaultServer注册并公布rcvr的方法。
	//RegisterName函数类似Register函数，但使用提供的name代替rcvr的具体类型名作为服务名。
	err := rpc.Register(arith)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go http.Serve(listener, nil)
	time.Sleep(10 * time.Second)
}
