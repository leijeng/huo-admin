package utils

import (
	"context"
	"github.com/leijeng/huo-admin/common/config"
	"github.com/leijeng/huo-core/core"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"sync"
)

var redisClient *redis.Client

var redisMutex sync.Mutex

func Redis() (*redis.Client, error) {
	redisCfg := config.Ext.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.Db,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		core.Log.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		core.Log.Info("redis connect ping response:", zap.String("pong", pong))
	}

	return client, err
}

func InitRedis() *redis.Client {
	redisMutex.Lock()
	defer redisMutex.Unlock()

	if redisClient != nil {
		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			redisClient, _ = Redis()
		}
	} else {
		redisClient, _ = Redis()
	}
	return redisClient
}
