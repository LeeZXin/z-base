package redis

import (
	"github.com/LeeZXin/zsf/property"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	Client *redis.Client
)

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:         property.GetString("redis.address"),
		Password:     property.GetString("redis.password"),
		DB:           property.GetInt("redis.db"),
		PoolFIFO:     true,
		PoolSize:     property.GetInt("redis.poolSize"),
		MinIdleConns: property.GetInt("redis.minIdleConns"),
		MaxIdleConns: property.GetInt("redis.maxIdleConns"),
		PoolTimeout:  time.Duration(property.GetInt("redis.poolTimeout")),
	})
}
