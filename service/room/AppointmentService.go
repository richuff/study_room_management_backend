package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"study_room_management_backend/model"
	"study_room_management_backend/model/dto"
	"study_room_management_backend/result"
	"study_room_management_backend/utils"
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

// ShowAppointment
// @Summary 查看预约自习室
// @Tags 自习室模块
// @Description 预约自习室接口
// @Param user_id query int true "用户id"
// @Success   200 {object} result.CodeResp "业务代码"
// @Security Bearer
// @Router /api/room/showAppointment [get]
func ShowAppointment(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Request.FormValue("user_id"), 10, 64)
	if utils.ErrHandler(c, err) {
		return
	}
	data := model.FindAppointment(userId)
	result.Ok(c, 1, data)
}
