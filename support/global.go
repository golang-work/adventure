package support

import (
	"github.com/go-redis/redis/v8"
	"github.com/golang-work/adventure/provider"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config provider.Config
	Viper  *viper.Viper
	Log    *zap.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)
