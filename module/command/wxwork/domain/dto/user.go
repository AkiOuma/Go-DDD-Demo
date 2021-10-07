package dto

type UserDTO struct {
	UserID string `json:"Userid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"telephone"`
	Avatar string `json:"avatar"`
}
