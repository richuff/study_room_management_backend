package dto

type UserLoginDto struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
