package dto

type UpdateProductDTO struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
	ImageUrl    string  `jsonL:"imageurl"`
}
