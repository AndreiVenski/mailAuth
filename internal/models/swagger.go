package models

type UserSwagger struct {
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
