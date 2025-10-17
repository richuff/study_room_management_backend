package model

import (
	"fmt"
	"study_room_management_backend/mapper"
	"study_room_management_backend/model/dto"
	"study_room_management_backend/utils"
)

type Student struct {
	StuId  string `gorm:"primaryKey;column:stu_id" json:"stu_id"`
	UserId uint64 `gorm:"column:user_id" json:"user_id"`
	Gender uint   `gorm:"column:gender" json:"gender"`
	Major  string `gorm:"column:major" json:"major"`
	Class  string `gorm:"column:class" json:"class"`
	User   User   `gorm:"foreignKey:UserId"`
}

func (Student) TableName() string {
	return "student"
}

func GetStudentByUserID(id uint64) Student {
	student := Student{}
	mapper.Open.Where("user_id = ?", id).Find(&student)
	return student
}

func InsertStudentInfo(dto *dto.UserInfoDto) {
	student := &Student{}
	err := utils.SimpleCopyProperties(student, dto)
	if err != nil {
		fmt.Println(err)
		return
	}
	mapper.Open.Create(student)
}
