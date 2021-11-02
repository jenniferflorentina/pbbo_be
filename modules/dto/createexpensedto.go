package dto

import (
	"tubespbbo/base"
)

type CreateExpenseDTO struct {
	base.DTO
	Nama        string  `json:"nama" validate:"empty=false"`
	ReleaseDate string  `json:"releaseDate" validate:"empty=false"`
	Quantity    float32 `json:"quantity"`
	Category    int64   `json:"category" validate:"empty=false"`
	Description string  `json:"description"`
}
