package dto

type UserInfoDto struct {
	StuId  string `json:"stu_id"`
	UserId uint64 `json:"user_id"`
	Gender uint   `json:"gender"`
	Major  string `json:"major"`
	Class  string `json:"class"`
}
