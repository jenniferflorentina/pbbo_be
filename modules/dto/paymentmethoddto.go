package dto

import (
	"tubespbbo/base"
)

type PaymentMethodDTO struct {
	base.DTO
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}
