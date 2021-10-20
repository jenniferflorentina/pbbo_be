package dto

import (
	"tubespbbo/base"
)

type TransactionDTO struct {
	base.DTO
	UserId           int64  `json:"idUser"`
	PembayaranId     int64  `json:"idPembayaran"`
	TanggalTransaksi string `json:"tanggalTransaksi"`
	NoResi           string `json:"noResi"`
	Status           string `json:"status"`
}
