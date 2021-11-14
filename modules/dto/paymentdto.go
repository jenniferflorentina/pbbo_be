package dto

import (
	"tubespbbo/base"
)

type PaymentDTO struct {
	base.DTO
	Status          string            `json:"status"`
	Amount          float32           `json:"amount"`
	AccountNumber      string            `json:"accountNumber"`
	PaymentMethodId int64             `json:"paymentMethodId"`
	PaymentMethod   *PaymentMethodDTO `json:"paymentMethod"`
}
