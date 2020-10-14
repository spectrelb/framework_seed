package users

import (
	"framework_seed/models/req/users"
	"framework_seed/pkg/Response"
	"framework_seed/service/users"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	Response.Success(c)
}

func Login(c *gin.Context) {
	var loginReq users.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		Response.Error(c, Response.CodeServerBusy)
		return
	}
	service.Login(c, &loginReq)
}