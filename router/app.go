package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	docsRouter(r)

	api := r.Group("/api")
	userRouter(api)
	roomRouter(api)

	return r
}
