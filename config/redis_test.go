package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

var ctx = context.Background()
var RDB *redis.Client

func TestRedisClient(t *testing.T) {
	address := "mylab:6379"
	password := ""
	db := 0

	RDB = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("abc")

}
