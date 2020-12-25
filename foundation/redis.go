package foundation

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang-work/adventure/support"
	"go.uber.org/zap"
	"os"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     support.Config["database"].GetString("redis.addr"),
		Password: support.Config["database"].GetString("redis.password"),
		DB:       support.Config["database"].GetInt("redis.db"),
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
