package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "172.16.3.7:6379", // address
		Password: "redispasswd",     // no password set
		DB:       0,                 // use default DB
	})

	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Ping err:", err)
		return
	}
	fmt.Printf("the result of Ping is %v.\n", pong)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
