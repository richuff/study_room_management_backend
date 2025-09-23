package room

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"study_room_management_backend/model"
	"study_room_management_backend/result"
	"study_room_management_backend/utils"
)

// Show
// @Summary 查看自习室状态
// @Tags 自习室模块
// @Description 自习室查看接口
// @Param room_id  query int true "自习室id"
// @Success   200 {object} result.CodeResp "业务代码"
// @Security Bearer
// @Router /api/room/show [get]
func Show(c *gin.Context) {
	roomId, err := strconv.ParseUint(c.Request.FormValue("room_id"), 10, 64)
	if utils.ErrHandler(c, err) {
		return
	}
	room := model.GetRoomById(roomId)
	if room.RoomName != "" {
		result.Ok(c, 1, room)
	} else {
		result.Error(c, "该自习室不存在")
	}
}
