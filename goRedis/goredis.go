package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	const (
		RedisAddr = "9.135.178.106:6380"
		Pass      = "Teg1174421473"
	)
	//连接redis
	redisCli := redis.NewClient(&redis.Options{Addr: RedisAddr, Password: Pass})
	_, err := redisCli.Ping().Result()
	if err != nil {
		log.Fatalf("Redis Connection Failed %v", RedisAddr)
		return
	}
	//获取redis中所有key
	keys,err:=redisCli.Keys("*").Result()
	fmt.Println(keys)
	//增删改查
	op :=&redis.StatusCmd{}
	op = redisCli.Set("name","mchuang",-1)
	fmt.Println(op.Result())
	//查询
	fmt.Println(redisCli.Get("name").Result())
	//删除
	redisCli.Del("myName")
	//改key
	redisCli.Rename("name","id")
	//查询
	fmt.Println(redisCli.Get("id").Result())

}
