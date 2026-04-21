package dto

type ForgetPwdDto struct {
	Code     string `json:"code"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
