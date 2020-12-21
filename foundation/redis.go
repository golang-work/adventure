package foundation

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-work/adventure/support"
	"go.uber.org/zap"
	"os"
)

func Redis() *redis.Client {
	conf := support.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	pong, err := client.Ping(context.TODO()).Result()
	if err != nil {
		support.Log.Error("redis connect ping failed, err:", zap.Any("err", err))
		os.Exit(0)
	} else {
		support.Log.Info("redis connect ping response:", zap.String("pong", pong))
	}

	return client
}
