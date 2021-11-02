package dto

type CreatePaymentMethodDTO struct {
	Nama string `json:"nama" validate:"empty=false"`
	Kode string `json:"kode" validate:"empty=false"`
}
