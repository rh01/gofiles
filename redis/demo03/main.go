package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {
	// 操作哈希
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

	// 开始操作哈希
	reply, err := conn.Do("HSet", "user01", "name", "tom")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("get reply is", reply)

	// 这里取数据
	// 使用redis指令即可
	res, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("get res is", res)
}
