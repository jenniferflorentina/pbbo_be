package dto

import (
	"tubespbbo/base"
)

type ExpenseDTO struct {
	base.DTO
	Nama          string   `json:"Nama"`
	ReleaseDate   string   `json:"ReleaseDate"`
	Quantity      float32  `json:"Quantity"`
	Category	  int64    `json:"Category"`
	Description   string   `json:"Description"`
}
