package model

import (
	"study_room_management_backend/mapper"
	"time"
)

type User struct {
	UserId    uint64    `json:"user_id" gorm:"primaryKey;autoIncrement;column:user_id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	IsDelete  bool      `gorm:"column:is_delete;type:boolean" json:"is_delete"`
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
	if user.Name != "" {
		return true
	}
	return false
}

func GetUserRe(email string) User {
	user := User{}
	mapper.Open.Where("email = ?", email).Find(&user)
	return user
}

func GetMessageByPassword(email string, password string) (bool, string) {
	user := User{}
	mapper.Open.Where("password = ?", password).Where("email = ?", email).Find(&user)
	if user.Name != "" && user.IsDelete == false {
		return true, "登录成功"
	} else if user.Name != "" && user.IsDelete != false {
		return false, "该用户已注销"
	}
	return false, "账号或密码错误"
}

func GetUserByPassword(email string, password string) User {
	user := User{}
	mapper.Open.Where("password = ?", password).Where("email = ?", email).Find(&user)
	return user
}

func UpdateUser(user *User) {
	mapper.Open.Save(user)
}

func GetUserByUserID(id uint64) User {
	user := User{}
	mapper.Open.Where("user_id = ?", id).Find(&user)
	return user
}

func GetUserByUserName(name string) *User {
	user := User{}
	mapper.Open.Where("name = ?", name).Find(&user)
	return &user
}

func UpdateAvatar(userId uint64, url string) bool {
	user := GetUserByUserID(userId)
	if user.Name == "" && user.IsDelete == true {
		return false
	}
	user.Avatar = url
	mapper.Open.Save(user).Where("user_id = ?", userId)
	return true
}
