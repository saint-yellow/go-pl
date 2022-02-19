package bootstrap

import (
	"fmt"

	"github.com/saint-yellow/go-pl/toys/gohub/pkg/config"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/redis"
)

// setupRedis 初始化 Redis
func setupRedis() {
	// 建立 Redis 连接
    redis.Connect(
        fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
        config.GetString("redis.username"),
        config.GetString("redis.password"),
        config.GetInt("redis.database"),
    )
}
