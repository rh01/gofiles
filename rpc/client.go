package main

import (
	"net/rpc"
	"log"
	"go-learn/rpc/server"
	"fmt"
)

func main() {

	client, err := rpc.DialHTTP("tcp", ":3030")
	if err != nil {
		log.Fatal("Dial http error,", err)
	}

	args := &server.Args{A: 10, B: 2}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
