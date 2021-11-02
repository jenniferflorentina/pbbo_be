package dto

type CreateUserDTO struct {
	Username   string `json:"username" validate:"empty=false"`
	Password   string `json:"password" validate:"empty=false"`
	Nama       string `json:"nama"`
	Dob        string `json:"dob"`
	Alamat     string `json:"alamat"`
	Telepon    string `json:"telepon"`
	Email      string `json:"email" validate:"empty=false"`
}
