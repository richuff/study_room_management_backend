package model

import (
	"study_room_management_backend/mapper"
	"time"
)

type User struct {
	UserId    string    `json:"user_id"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsDelete  bool      `json:"is_delete"`
}

func (table *User) TableName() string {
	return "user"
}

func CreateUser(user *User) bool {
	if GetUserByEmail(user.Email) {
		return false
	}
	mapper.Open.Create(&user)
	return true
}

func GetUserByEmail(email string) bool {
	user := User{}
	mapper.Open.Where("email = ?", email).Find(&user)
	if user.UserId != "" {
		return true
	}
	return false
}

func GetUserByPassword(email string, password string) bool {
	user := User{}
	mapper.Open.Where("password = ?", password).Where("email = ?", email).Find(&user)
	if user.UserId != "" {
		return true
	}
	return false
}

func UpdateUser(user *User) {
	mapper.Open.Save(user)
}
