package dto

type UpdateTransactionDTO struct {
	UserId           int64  `json:"idUser"`
	TanggalTransaksi string `json:"tanggalTransaksi"`
	NoResi           string `json:"noResi"`
	Status           string `json:"status"`
}
