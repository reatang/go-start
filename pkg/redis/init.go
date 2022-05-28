package redis

import (
	"context"
	goRedis "github.com/go-redis/redis/v8"
	"log"
	"time"
)

var rdb *goRedis.Client

func InitRedis(config *RedisConfig) {
	rdb = goRedis.NewClient(&goRedis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *goRedis.Client {
	return rdb
}