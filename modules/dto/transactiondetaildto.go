package dto

import (
	"tubespbbo/base"
)

type TransactionDetailDTO struct {
	base.DTO
	TransaksiId int64       `json:"idTransaksi"`
	ProductId   int64       `json:"idProduk"`
	Product     *ProductDTO `json:"product"`
	Jumlah      int64       `json:"jumlah"`
}
