package model

import (
	"study_room_management_backend/mapper"
	"study_room_management_backend/model/dto"
	"time"
)

type Appointment struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID    uint64    `gorm:"column:user_id" json:"user_id"`
	RoomID    uint64    `gorm:"column:room_id" json:"room_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (table *Appointment) TableName() string {
	return "appointment"
}

func InsertAppointment(dto *dto.AppointmentDto) {
	appointment := &Appointment{}
	appointment.RoomID = dto.RoomId
	appointment.UserID = dto.UserId
	appointment.CreatedAt = time.Now()
	mapper.Open.Create(appointment)
}
