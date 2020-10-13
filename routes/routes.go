package routes

import (
	"fmt"
	"framework_seed/controller/users"
	"framework_seed/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() (r *gin.Engine) {
	r = gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	user := r.Group("/user")
	{
		user.POST("/register", users.Register)
		user.POST("/login", users.Login)
	}

	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		if _, err := c.FormFile("file"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("uploaded fail!"))
			return
		}
		file, _ := c.FormFile("file")

		fmt.Println(file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	return
}
