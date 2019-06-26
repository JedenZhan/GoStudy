package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	var r string
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.err", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("Set", "name", "TomJerry")
	if err != nil {
		fmt.Println("啊啊啊啊错误了", err)
	}

	r, err = redis.String(conn.Do("Get", "name"))
	fmt.Println("connect and do Ok", r)
}