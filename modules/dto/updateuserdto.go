package dto

type UpdateUserDTO struct {
	Username string `json:"username"`
	Address   string `json:"address"`
	Phone  string `json:"phone"`
}