package dto

import (
	"tubespbbo/base"
)

type CreateExpenseDTO struct {
	base.DTO
	Nama          string   `json:"Nama" validate:"empty=false"`
	ReleaseDate   string   `json:"ReleaseDate" validate:"empty=false"`
	Quantity      float32  `json:"Quantity"`
	Category	  int64    `json:"Category" validate:"empty=false"`
	Description   string   `json:"Description"`
}
