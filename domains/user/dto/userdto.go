package dto

import (
	"tubespbbo/base"
)

type UserDTO struct {
	base.DTO
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Telepon  string `json:"telepon"`
	Email    string `json:"email"`
}
