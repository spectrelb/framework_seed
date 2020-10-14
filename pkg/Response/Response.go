package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Error(c *gin.Context, code MyCode) {
	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: code.Msg(),
	})
}

func ErrorWithMsg(c *gin.Context, code MyCode, msg string) {
	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: msg,
	})
}

func Success(c *gin.Context, data... interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}