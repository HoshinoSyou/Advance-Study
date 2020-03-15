package sqlConn

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func Redis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Printf("Redis连接失败喵！错误信息：%v", err)
		return rdb
	}
	defer rdb.Close()
	return rdb
}
