package dto

import (
	"tubespbbo/base"
)

type ExpenseDTO struct {
	base.DTO
	Nama        string  `json:"nama"`
	ReleaseDate string  `json:"releaseDate"`
	Quantity    float32 `json:"quantity"`
	Category    int64   `json:"category"`
	Description string  `json:"description"`
}
