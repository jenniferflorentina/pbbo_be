package dto

import (
	"tubespbbo/base"
)

type TransactionDetailDTO struct {
	base.DTO
	TransactionId int64       `json:"transactionId"`
	ProductId   int64       `json:"productId"`
	Product     *ProductDTO `json:"product"`
	Quantity      int64       `json:"quantity"`
}
