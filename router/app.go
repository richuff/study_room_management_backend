package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"study_room_management_backend/docs"
	"study_room_management_backend/service/user"
)

func Router() *gin.Engine {
	r := gin.Default()
	docsRouter(r)

	api := r.Group("/api")
	userRouter(api)

	return r
}

func docsRouter(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func userRouter(api *gin.RouterGroup) {
	userApi := api.Group("/user")
	userApi.POST("/login", user.Login)
	userApi.POST("/register", user.Register)
}
