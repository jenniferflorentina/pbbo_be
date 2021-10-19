package dto

import (
	"tubespbbo/base"
)

type ProductDTO struct {
	base.DTO
	Nama	        string            `json:"nama"`
	Kode	        string	          `json:"kode"`
	Description     string            `json:"descreption"`
	Price			float32           `json:"price"`
	Quantity   		int64 			  `json:"quantity"`
	ImageUrl		string			  `jsonL:"imageurl"`
}
