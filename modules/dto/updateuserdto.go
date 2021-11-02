package dto

type UpdateUserDTO struct {
	Username string `json:"username"`
	Alamat   string `json:"alamat"`
	Telepon  string `json:"telepon"`
}