package redis

import (
	"fmt"
	"context"
	"framework_seed/settings"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: fmt.Sprintf("%s", cfg.Password), // no password set
		DB:        cfg.DB,  // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

// Set a key/value
func Del(key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		zap.L().Error("redis Del value fail:", zap.Error(err))
		return err
	}

	return nil
}

// Set a key/value
func Set(key string, data interface{}, expiration time.Duration) error {
	err := rdb.Set(ctx, key, data, expiration).Err()
	if err != nil {
		zap.L().Error("redis Set value fail:", zap.Error(err))
		return err
	}

	return nil
}

// Get get a key
func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		zap.L().Error("redis Get value fail:", zap.Error(err))
		return "", err
	}

	zap.L().Info("redis Get value is:", zap.String("Get", val))
	return val, nil
}

func Close() {
	if err := rdb.Close(); err !=nil {
		zap.L().Info("redis Close fail :", zap.Error(err))
	}
}