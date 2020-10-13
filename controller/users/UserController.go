package users

import (
	"fmt"
	"framework_seed/models/req/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context)  {
	var loginReq users.LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		fmt.Println()
	}
	fmt.Println(loginReq)
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func Login(c *gin.Context)  {
	var loginReq users.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": loginReq,
	})
}