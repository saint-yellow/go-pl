package cache

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

const (
	RankKey = "rank"
)

func TaskViewKey(id uint) string {
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}

func InitializeRedisClient(host, port, password, dbName string) {
	db, err := strconv.Atoi(dbName)
	if err != nil {
		panic(fmt.Errorf("invalid redis DB name: %s", err.Error()))
	}
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB: db,
	})
	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Println("Redis is successfully connected")
	RedisClient = client
}