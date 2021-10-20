package dto

import (
	"tubespbbo/base"
)

type TransactionDetailDTO struct {
	base.DTO
	TransaksiId int64 `json:"idTransaksi"`
	ProdukId    int64 `json:"idProduk"`
	Jumlah      int64 `json:"jumlah"`
}
