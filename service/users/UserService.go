package service

import (
	users2 "framework_seed/dao/users"
	"framework_seed/models/req/users"
	"framework_seed/pkg/Response"
	"framework_seed/pkg/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)



func Login(c *gin.Context, req *users.LoginReq) {
	//redis.Del("userinfo")
	_, err := redis.Get("userinfo")
	if err != nil {
		exist, err := users2.ExistUserByUserName(req.Username)
		if err != nil || !exist {
			zap.L().Error("login fail", zap.Error(err))
			Response.Error(c, Response.CodeLoginFail)
			return
		}
	}

	redis.Set("userinfo", req.Username, time.Second * 24)
	zap.L().Info("login success", zap.String("login success", req.Username))
	Response.Success(c)
}