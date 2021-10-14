package dto

import (
	"tubespbbo/base"
)

type UserDTO struct {
	base.DTO
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}
