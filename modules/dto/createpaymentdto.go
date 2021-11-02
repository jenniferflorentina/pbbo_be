package dto

type CreatePaymentDTO struct {
	NoRekening      string  `json:"noRekening" validate:"empty=false"`
	PaymentMethodId int64   `json:"paymentMethodId"`
	Amount          float32 `json:"amount"`
	TransaksiId     int64   `json:"transaksiId"`
}
