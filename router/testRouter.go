package router

import (
	"github.com/gin-gonic/gin"
	"study_room_management_backend/service/code"
)

func testRouter(test *gin.RouterGroup) {
	test.GET("/captcha", code.Captcha)
	test.GET("/captcha/verify", code.Verify)
}
