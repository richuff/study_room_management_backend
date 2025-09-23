package router

import (
	"github.com/gin-gonic/gin"
	"study_room_management_backend/service/room"
)

func roomRouter(api *gin.RouterGroup) {
	roomApi := api.Group("/room")
	roomApi.GET("/show", room.Show)
	roomApi.POST("/appointment", room.Appointment)
}
