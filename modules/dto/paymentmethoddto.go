package dto

import (
	"tubespbbo/base"
)

type PaymentMethodDTO struct {
	base.DTO
	Name string `json:"name"`
	Code string `json:"code"`
}
