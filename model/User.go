package model

import (
	"study_room_management_backend/mapper"
	"time"
)

type User struct {
	UserId    string    `json:"user_id" gorm:"primaryKey;autoIncrement;column:user_id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
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

func GetUserByPassword(email string, password string) (bool, string) {
	user := User{}
	mapper.Open.Where("password = ?", password).Where("email = ?", email).Find(&user)
	if user.UserId != "" && user.IsDelete == false {
		return true, "登录成功"
	} else if user.UserId != "" && user.IsDelete != false {
		return false, "该用户已注销"
	}
	return false, "账号或密码错误"
}

func UpdateUser(user *User) {
	mapper.Open.Save(user)
}

func GetUserByUserID(id uint64) User {
	user := User{}
	mapper.Open.Where("user_id = ?", id).Find(&user)
	return user
}
