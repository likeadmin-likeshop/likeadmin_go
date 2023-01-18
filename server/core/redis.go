package core

import (
	"context"
	"github.com/go-redis/redis/v9"
	"likeadmin/config"
	"log"
	"time"
)

var Redis = initRedis()

//initRedis 初始化redis客户端
func initRedis() *redis.Client {
	opt, err := redis.ParseURL(config.Config.RedisUrl)
	if err != nil {
		log.Fatal("initRedis redis.ParseURL err: ", err)
	}
	opt.PoolSize = config.Config.RedisPoolSize
	client := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("initRedis client.Ping err: ", err)
	}
	return client
}
