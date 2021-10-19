package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redistClient *redis.Client

func initRedis() (err error) {
	redistClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redistClient.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err: %v", err)
		return
	}
	defer redistClient.Close()
	fmt.Println("redis connect success")

	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 99, Member: "PHP"},
		&redis.Z{Score: 96, Member: "JAVA"},
		&redis.Z{Score: 95, Member: "GO"},
		&redis.Z{Score: 83, Member: "Rust"},
	}
	redistClient.ZAdd(key, items...)

}
