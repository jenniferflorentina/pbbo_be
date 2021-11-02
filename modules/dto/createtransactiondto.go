package dto

type CreateTransactionDTO struct {
	UserId           int64  `json:"idUser" validate:"empty=false"`
	TanggalTransaksi string `json:"tanggalTransaksi" validate:"empty=false"`
	NoResi           string `json:"noResi" validate:"empty=false"`
	Status           string `json:"status" validate:"empty=false"`
}
