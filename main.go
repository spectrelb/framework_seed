package main

import (
	"fmt"
	"framework_seed/logger"
	"framework_seed/pkg/mysql"
	"framework_seed/pkg/redis"
	"framework_seed/routes"
	"framework_seed/settings"
	"go.uber.org/zap"
	"time"
)

func main() {

	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		zap.L().Error("load config failed, err:", zap.Error(err))
		return
	}

	defer redis.Close()
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("load redis config failed, err:", zap.Error(err))
		return
	}

	r := routes.InitRoutes()
	err := r.Run(fmt.Sprintf(":%s", settings.Conf.Port))

	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

	fmt.Println(time.Date(2020, time.December, 19, 0, 0, 0, 0, time.UTC).Add(300).Format("2006-01-02"))
}
