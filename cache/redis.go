package cache

import (
	"context"

	"github.com/CyberMidori/gin-ranking/config"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		DB:   config.RedisDb,
	})
	Rctx = context.Background()
}

func Zscore(member interface{}, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: member}
}
