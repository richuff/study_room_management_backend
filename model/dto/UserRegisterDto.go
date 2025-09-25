package dto

type UserRegisterDto struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
