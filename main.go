package main

import (
	"fmt"
	"framework_seed/pkg/mysql"
	"framework_seed/pkg/redis/redis"
	"framework_seed/routes"
	"framework_seed/settings"
)

func main() {

	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}


	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}


	r := routes.InitRoutes()
	r.Run(fmt.Sprintf(":%s", settings.Conf.Port))
}
