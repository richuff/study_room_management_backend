package router

import (
	"github.com/gin-gonic/gin"
	"study_room_management_backend/middleware"
	"study_room_management_backend/service/code"
)

func Router() *gin.Engine {
	r := gin.Default()
	docsRouter(r)

	test := r.Group("/test")
	testRouter(test)

	api := r.Group("/api")
	api.GET("/captcha", code.SendSmsCode)          // 获取验证码
	api.GET("/captcha/verify", code.VerifySmsCode) // 验证验证码

	userRouter(api)

	api.Use(middleware.JWTAuth())
	{
		roomRouter(api)
	}

	return r
}
