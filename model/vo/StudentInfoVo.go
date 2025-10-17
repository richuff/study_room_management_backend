package vo

type StudentInfoVo struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Gender uint   `json:"gender"`
	Major  string `json:"major"`
	Class  string `json:"class"`
}
