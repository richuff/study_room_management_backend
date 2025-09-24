package dto

type AppointmentDto struct {
	RoomId uint64 `json:"room_id"`
	UserId uint64 `json:"user_id"`
}
