package dto

type UserRegisterDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
