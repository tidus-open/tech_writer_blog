package tdao

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisCli *redis.Client

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisCli.Ping().Result()
	fmt.Println(redisCli, pong, err)
}
