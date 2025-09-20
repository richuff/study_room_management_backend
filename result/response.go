package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Any interface{}

// Ok 用于返回正确的响应
func Ok(c *gin.Context, code int, data Any) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

// Error 用于返回错误的响应
func Error(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// ErrorWithCode 用于返回错误的响应并且附带响应码
func ErrorWithCode(c *gin.Context, message string, code int) {
	c.JSON(code, gin.H{
		"message": message,
	})
}
