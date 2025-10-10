package model

import (
	"study_room_management_backend/mapper"
	"time"
)

type Room struct {
	RoomID        uint64    `gorm:"primaryKey;autoIncrement;column:room_id" json:"room_id"`
	RoomName      string    `gorm:"unique;not null;column:room_name;size:60" json:"room_name"`
	Floor         uint8     `gorm:"not null;column:floor" json:"floor"`
	SeatTotal     *uint16   `gorm:"column:seat_total" json:"seat_total,omitempty"`         // 允许 NULL
	SeatAvailable *uint16   `gorm:"column:seat_available" json:"seat_available,omitempty"` // 允许 NULL
	RoomType      *uint8    `gorm:"column:room_type" json:"room_type,omitempty"`           // 指针可存 NULL
	Desc          string    `gorm:"column:desc;size:500" json:"desc,omitempty"`
	OpenTime      *string   `gorm:"column:open_time;type:time" json:"open_time,omitempty"` // hh:mm:ss
	CloseTime     *string   `gorm:"column:close_time;type:time" json:"close_time,omitempty"`
	Status        *uint8    `gorm:"column:status" json:"status,omitempty"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:update_at;autoUpdateTime" json:"update_at"`
}

const (
	RoomTypeNormal  uint8 = 1 // 普通区
	RoomTypeQuiet   uint8 = 2 // 静音区
	RoomTypeDiscuss uint8 = 3 // 讨论区
)

const (
	RoomStatusOffline  uint8 = 0 // 被占用
	RoomStatusOnline   uint8 = 1 // 正常
	RoomStatusMaintain uint8 = 2 // 维护
)

// TableName 设置表名
func (Room) TableName() string {
	return "room"
}

func GetRoomById(roomId uint64) *Room {
	room := &Room{}
	mapper.Open.Where("room_id=?", roomId).Find(&room)
	return room
}

// GetRoomIdle 获取空闲的的自习室
func GetRoomIdle() []Room {
	rooms := make([]Room, 10)
	mapper.Open.Model(Room{}).
		Where("status = ?", RoomStatusOnline).
		Find(&rooms)

	return rooms
}
