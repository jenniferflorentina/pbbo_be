package dto

import (
	"tubespbbo/base"
)

type PaymentDTO struct {
	base.DTO
	Status          string            `json:"status"`
	Amount          float32           `json:"amount"`
	NoRekening      string            `json:"noRekening"`
	PaymentMethodId int64             `json:"paymentMethodId"`
	PaymentMethod   *PaymentMethodDTO `json:"paymentMethod"`
}
