package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	// 创建连接
	conn, err := redis.Dial(
		"tcp",
		"172.16.3.7:6379",
		redis.DialPassword("redispasswd"),
	)
	defer conn.Close()

	if err != nil {
		fmt.Println("redis conn err:", err)
		return
	}
	fmt.Println("redis connection success")

	// 测试ping-pong
	reply, err := conn.Do("ping")
	if err != nil {
		fmt.Println("PING err, ", err)
		return
	}
	fmt.Println("get reply should be PONG, actual reply is:", reply)

	// set 操作，存数据
	reply, err = conn.Do("set", "key", "value")
	if err != nil {
		fmt.Println("set key err:", err)
		return
	}
	//fmt.Printf("get key=key, value=%s", reply)

	// 取数据
	reply, err = redis.String(conn.Do("get", "key"))
	if err != nil {
		fmt.Println("get key err", err)
		return
	}
	fmt.Println("get value is", reply)

}
