package dto

type CreateProductDTO struct {
	Name        string  `json:"name" validate:"empty=false"`
	Code        string  `json:"code" validate:"empty=false"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
	ImageUrl    string  `jsonL:"imageurl" validate:"empty=false"`
}
