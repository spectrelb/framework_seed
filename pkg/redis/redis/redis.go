package redis

import (
	"fmt"
	"framework_seed/settings"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: fmt.Sprintf("%s", ), // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}