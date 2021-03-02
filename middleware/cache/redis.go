package cache

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"matching-engine/config"
	"matching-engine/log"
)

var RedisClient *redis.Client

func Init() {
	conf := config.Conf.Redis

	//address := conf.Addr
	//password := conf.Password
	//db := conf.DB
	//TODO 这里使用常量才能运行通过？用变量都不通过
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "mylab:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Logger.Info("Connected to redis", zap.String("addr", conf.Addr))
}
