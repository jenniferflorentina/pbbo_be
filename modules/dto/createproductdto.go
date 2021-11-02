package dto

type CreateProductDTO struct {
	Nama        string  `json:"nama" validate:"empty=false"`
	Kode        string  `json:"kode" validate:"empty=false"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"empty=false"`
	Quantity    int64   `json:"quantity"`
	ImageUrl    string  `jsonL:"imageurl" validate:"empty=false"`
}
