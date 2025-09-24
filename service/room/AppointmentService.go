package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"study_room_management_backend/model"
	"study_room_management_backend/model/dto"
	"study_room_management_backend/result"
)

// Appointment
// @Summary 预约自习室
// @Tags 自习室模块
// @Description 预约自习室接口
// @Accept    json
// @Produce   json
// @Param     req body dto.AppointmentDto true "登录信息"
// @Success   200 {object} result.CodeResp "业务代码"
// @Security Bearer
// @Router /api/room/appointment [post]
func Appointment(c *gin.Context) {
	appointmentDto := &dto.AppointmentDto{}
	if err := c.BindJSON(appointmentDto); err != nil {
		fmt.Println(err)
		return
	}

	room := model.GetRoomById(appointmentDto.RoomId)
	if room.RoomName != "" {
		model.InsertAppointment(appointmentDto)
		result.Ok(c, 1, "预约成功")
	} else {
		result.Error(c, "该自习室不存在")
	}
}
