package router

import (
	"github.com/gin-gonic/gin"
	"study_room_management_backend/service/user"
)

func userRouter(api *gin.RouterGroup) {
	userApi := api.Group("/user")
	userApi.POST("/login", user.Login)
	userApi.POST("/register", user.Register)
	userApi.GET("/logoff", user.Logoff)
	userApi.GET("/checkInfo", user.CheckInfo)
	userApi.POST("/setInfo", user.SetInfo)
	userApi.POST("/setAvatar", user.SetAvatar)
}
