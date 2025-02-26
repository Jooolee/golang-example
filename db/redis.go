package db

import (
	"github.com/go-redis/redis/v8"
)

// RedisClient 全局的 Redis 客户端实例
var RedisClient *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码，若没有则为空
		DB:       0,                // 使用的数据库编号
	})
}
